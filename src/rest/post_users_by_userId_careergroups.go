package rest

import (
	"context"
	"errors"

	pkgerrors "github.com/pkg/errors"
	"github.com/sky0621/cv-admin/src/ent"
	"github.com/sky0621/cv-admin/src/ent/helper"
	eskill "github.com/sky0621/cv-admin/src/ent/skill"
)

// PostUsersByUserIdCareergroups キャリアグループ新規登録
// キャリアグループ新規登録
// (POST /users/{byUserId}/careergroups)
func (s *strictServerImpl) PostUsersByUserIdCareergroups(ctx context.Context, request PostUsersByUserIdCareergroupsRequestObject) (PostUsersByUserIdCareergroupsResponseObject, error) {
	if request.Body == nil {
		return nil, errors.New("request body is nil")
	}
	userCareerGroup := *request.Body

	if err := userCareerGroup.Validate(); err != nil {
		return PostUsersByUserIdCareergroups400JSONResponse{n400("validation failed")}, err
	}

	var entUserCareerGroup *ent.UserCareerGroup
	if err := helper.WithTransaction(ctx, s.dbClient, func(tx *ent.Tx) error {
		var err error
		/*
		 * キャリアグループ本体の登録
		 */
		entUserCareerGroup, err = ToEntUserCareerGroupCreate(userCareerGroup, request.ByUserId, tx.UserCareerGroup.Create()).Save(ctx)
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
			entCareer, err := ToEntUserCareerCreate(career, entUserCareerGroup.ID, tx.UserCareer.Create()).Save(ctx)
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
				entDescriptions, err := tx.UserCareerDescription.CreateBulk(descriptionCreates...).Save(ctx)
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
					entCareerTask, err := ToEntCareerTaskCreate(task, entCareer.ID, tx.CareerTask.Create()).Save(ctx)
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
						entDescriptions, err := tx.CareerTaskDescription.CreateBulk(taskDescriptionCreates...).Save(ctx)
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
					entCareerSkillGroup, err := ToEntCareerSkillGroupCreate(skillGroup, entCareer.ID, tx.CareerSkillGroup.Create()).Save(ctx)
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
							entSkill, err := tx.Skill.Query().Where(eskill.Key(*skill.Skill.Key)).Only(ctx)
							if err != nil {
								return pkgerrors.WithMessagef(err, "Key: %v", *skill.Skill.Key)
							}
							if entSkill == nil {
								return errors.New("no skill")
							}
							careerSkillCreates = append(careerSkillCreates, ToEntCareerSkillCreate(skill, entCareerSkillGroup.ID, entSkill.ID, tx.CareerSkill.Create()))

							keys = append(keys, *skill.Skill.Key)
						}
						careerSkills, err := tx.CareerSkill.CreateBulk(careerSkillCreates...).Save(ctx)
						if err != nil {
							return err
						}

						for i, careerSkill := range careerSkills {
							skill, err := tx.Skill.Query().Where(eskill.Key(keys[i])).Only(ctx)
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
		return nil, err
	}

	return PostUsersByUserIdCareergroups201JSONResponse{N201CreatedUserCareerGroupJSONResponse(ToSwaggerUserCareerGroup(entUserCareerGroup))}, nil
}
