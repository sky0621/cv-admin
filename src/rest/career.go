package rest

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	pkgerrors "github.com/pkg/errors"

	"github.com/sky0621/cv-admin/src/ent"
	"github.com/sky0621/cv-admin/src/ent/helper"
	eskill "github.com/sky0621/cv-admin/src/ent/skill"
	"github.com/sky0621/cv-admin/src/ent/user"
	"github.com/sky0621/cv-admin/src/ent/usercareergroup"
)

// GetUsersByUserIdCareergroups キャリアグループ群取得
// (GET /users/{byUserId}/careergroups)
func (s *ServerImpl) GetUsersByUserIdCareergroups(ctx echo.Context, byUserId UserId) error {
	// TODO: 直接 user_id 指定できるはず。。
	careerGroups, err := s.dbClient.UserCareerGroup.Query().
		Where(usercareergroup.HasUserWith(user.ID(byUserId))).WithCareers(func(q *ent.UserCareerQuery) {
		q.WithCareerDescriptions()
		q.WithCareerSkillGroups(func(q *ent.CareerSkillGroupQuery) {
			q.WithCareerSkills(func(q *ent.CareerSkillQuery) {
				q.WithSkill()
			})
		})
		q.WithCareerTasks(func(q *ent.CareerTaskQuery) {
			q.WithCareerTaskDescriptions()
		})
	}).All(ctx.Request().Context())
	if err != nil {
		return sendClientError(ctx, http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, ToSwaggerUserCareerGroups(careerGroups))
}

// PostUsersByUserIdCareergroups キャリアグループ新規登録
// (POST /users/{byUserId}/careergroups)
func (s *ServerImpl) PostUsersByUserIdCareergroups(ctx echo.Context, byUserId UserId) error {
	var userCareerGroup UserCareerGroup
	if err := ctx.Bind(&userCareerGroup); err != nil {
		return sendClientError(ctx, http.StatusBadRequest, err.Error())
	}

	if err := userCareerGroup.Validate(); err != nil {
		return sendClientError(ctx, http.StatusBadRequest, err.Error())
	}

	rCtx := ctx.Request().Context()

	var entUserCareerGroup *ent.UserCareerGroup
	if err := helper.WithTransaction(rCtx, s.dbClient, func(tx *ent.Tx) error {
		var err error
		/*
		 * キャリアグループ本体の登録
		 */
		entUserCareerGroup, err = ToEntUserCareerGroupCreate(userCareerGroup, byUserId, tx.UserCareerGroup.Create()).Save(rCtx)
		if err != nil {
			return err
		}

		if userCareerGroup.Careers == nil {
			return nil
		}

		var entCareers []*ent.UserCareer
		/*
		 * キャリアグループ配下のキャリア群
		 */
		for _, career := range *userCareerGroup.Careers {
			/*
			 * キャリア本体の登録
			 */
			entCareer, err := ToEntUserCareerCreate(career, entUserCareerGroup.ID, tx.UserCareer.Create()).Save(rCtx)
			if err != nil {
				return err
			}

			/*
			 * キャリア説明群の登録
			 */
			if career.Description != nil {
				var descriptionCreates []*ent.UserCareerDescriptionCreate
				for _, description := range *career.Description {
					descriptionCreates = append(descriptionCreates, ToEntUserCareerDescriptionCreate(description, entCareer.ID, tx.UserCareerDescription.Create()))
				}
				entDescriptions, err := tx.UserCareerDescription.CreateBulk(descriptionCreates...).Save(rCtx)
				if err != nil {
					return err
				}

				entCareer.Edges.CareerDescriptions = entDescriptions
			}

			/*
			 * キャリア配下のタスク群
			 */
			if career.Tasks != nil {
				var careerTasks []*ent.CareerTask
				for _, task := range *career.Tasks {
					/*
					 * タスク本体の登録
					 */
					entCareerTask, err := ToEntCareerTaskCreate(task, entCareer.ID, tx.CareerTask.Create()).Save(rCtx)
					if err != nil {
						return err
					}

					/*
					 * タスク説明群の登録
					 */
					if task.Description != nil {
						var taskDescriptionCreates []*ent.CareerTaskDescriptionCreate
						for _, description := range *task.Description {
							taskDescriptionCreates = append(taskDescriptionCreates, ToEntCareerTaskDescriptionCreate(description, entCareerTask.ID, tx.CareerTaskDescription.Create()))
						}
						entDescriptions, err := tx.CareerTaskDescription.CreateBulk(taskDescriptionCreates...).Save(rCtx)
						if err != nil {
							return err
						}
						entCareerTask.Edges.CareerTaskDescriptions = entDescriptions
					}

					careerTasks = append(careerTasks, entCareerTask)
				}

				entCareer.Edges.CareerTasks = careerTasks
			}

			/*
			 * キャリア配下のスキルグループ群
			 */
			if career.SkillGroups != nil {
				var careerSkillGroups []*ent.CareerSkillGroup
				for _, skillGroup := range *career.SkillGroups {
					/*
					 * スキルグループ本体の登録
					 */
					entCareerSkillGroup, err := ToEntCareerSkillGroupCreate(skillGroup, entCareer.ID, tx.CareerSkillGroup.Create()).Save(rCtx)
					if err != nil {
						return err
					}

					/*
					 * スキルグループ配下のスキル群の登録
					 */
					if skillGroup.Skills != nil {
						var keys []string
						var careerSkillCreates []*ent.CareerSkillCreate
						for _, skill := range *skillGroup.Skills {
							entSkill, err := tx.Skill.Query().Where(eskill.Key(*skill.Skill.Key)).Only(rCtx)
							if err != nil {
								return pkgerrors.WithMessagef(err, "Key: %v", *skill.Skill.Key)
							}
							if entSkill == nil {
								return errors.New("no skill")
							}
							careerSkillCreates = append(careerSkillCreates, ToEntCareerSkillCreate(skill, entCareerSkillGroup.ID, entSkill.ID, tx.CareerSkill.Create()))

							keys = append(keys, *skill.Skill.Key)
						}
						careerSkills, err := tx.CareerSkill.CreateBulk(careerSkillCreates...).Save(rCtx)
						if err != nil {
							return err
						}

						for i, careerSkill := range careerSkills {
							skill, err := tx.Skill.Query().Where(eskill.Key(keys[i])).Only(rCtx)
							if err != nil {
								return err
							}
							careerSkill.Edges.Skill = skill
						}

						entCareerSkillGroup.Edges.CareerSkills = careerSkills
					}
					careerSkillGroups = append(careerSkillGroups, entCareerSkillGroup)
				}
				entCareer.Edges.CareerSkillGroups = careerSkillGroups
			}

			entCareers = append(entCareers, entCareer)
		}

		entUserCareerGroup.Edges.Careers = entCareers

		return nil
	}); err != nil {
		return sendClientError(ctx, http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, ToSwaggerUserCareerGroup(entUserCareerGroup))
}

// キャリアグループ削除
// (DELETE /users/{byUserId}/careergroups/{byCareerGroupId})
func (s *ServerImpl) DeleteUsersByUserIdCareergroupsByCareerGroupId(ctx echo.Context, byUserId UserId, byCareerGroupId CareerGroupId) error {
	// FIXME:
	return ctx.String(http.StatusOK, "")
}

// キャリアグループ更新
// (PUT /users/{byUserId}/careergroups/{byCareerGroupId})
func (s *ServerImpl) PutUsersByUserIdCareergroupsByCareerGroupId(ctx echo.Context, byUserId UserId, byCareerGroupId CareerGroupId) error {
	// FIXME:
	return ctx.String(http.StatusOK, "")
}

// キャリアグループ内キャリア群最新化
// (PUT /users/{byUserId}/careergroups/{byCareerGroupId}/careers)
func (s *ServerImpl) PutUsersByUserIdCareergroupsByCareerGroupIdCareers(ctx echo.Context, byUserId UserId, byCareerGroupId CareerGroupId) error {
	// FIXME:
	return ctx.String(http.StatusOK, "")
}

// CareerPeriodIf キャリアの開始・終了年月を束ねるためのマーカーインタフェース
type CareerPeriodIf interface {
	Set(year, month *int)
}

func (f *CareerPeriodFrom) Set(year, month *int) {
	f.Year = year
	f.Month = month
}

func (f *CareerPeriodTo) Set(year, month *int) {
	f.Year = year
	f.Month = month
}
