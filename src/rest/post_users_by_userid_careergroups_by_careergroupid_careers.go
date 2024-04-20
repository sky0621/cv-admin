package rest

import (
	"context"
	"errors"

	"github.com/sky0621/cv-admin/src/ent/helper"

	pkgerrors "github.com/pkg/errors"
	"github.com/sky0621/cv-admin/src/ent/skill"

	"github.com/sky0621/cv-admin/src/ent"
)

// PostUsersByUserIdCareergroupsByCareerGroupIdCareers
// １キャリアグループ内の１キャリア登録
// (POST /users/{byUserId}/careergroups/{byCareerGroupId}/careers)
func (s *strictServerImpl) PostUsersByUserIdCareergroupsByCareerGroupIdCareers(ctx context.Context, request PostUsersByUserIdCareergroupsByCareerGroupIdCareersRequestObject) (PostUsersByUserIdCareergroupsByCareerGroupIdCareersResponseObject, error) {
	if request.Body == nil {
		return nil, errors.New("request body is nil")
	}
	userCareer := *request.Body

	careerGroupId := request.ByCareerGroupId

	if err := userCareer.Validate(); err != nil {
		return PostUsersByUserIdCareergroupsByCareerGroupIdCareers400JSONResponse{n400("validation failed")}, err
	}

	var entCareer *ent.UserCareer
	if err := helper.WithTransaction(ctx, s.dbClient, func(tx *ent.Tx) error {
		var err error
		/*
		 * キャリア本体の登録
		 */
		entCareer, err = ToEntUserCareerCreate(userCareer, careerGroupId, tx.UserCareer.Create()).Save(ctx)
		if err != nil {
			return err
		}

		/*
		 * キャリア説明群の登録
		 */
		if userCareer.Description != nil {
			var descriptionCreates []*ent.UserCareerDescriptionCreate
			for _, description := range *userCareer.Description {
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
		if userCareer.Tasks != nil {
			var careerTasks []*ent.CareerTask
			for _, task := range *userCareer.Tasks {
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
		if userCareer.SkillGroups != nil {
			var careerSkillGroups []*ent.CareerSkillGroup
			for _, skillGroup := range *userCareer.SkillGroups {
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
					var skillIDs []int
					var careerSkillCreates []*ent.CareerSkillCreate
					for _, skill := range *skillGroup.Skills {
						entSkill, err := tx.Skill.Get(ctx, *skill.Skill.SkillID)
						if err != nil {
							return pkgerrors.WithMessagef(err, "Skill.SkillID: %v", *skill.Skill.SkillID)
						}
						if entSkill == nil {
							return errors.New("no skill")
						}
						careerSkillCreates = append(careerSkillCreates, ToEntCareerSkillCreate(skill, entCareerSkillGroup.ID, entSkill.ID, tx.CareerSkill.Create()))

						skillIDs = append(skillIDs, *skill.Skill.SkillID)
					}
					careerSkills, err := tx.CareerSkill.CreateBulk(careerSkillCreates...).Save(ctx)
					if err != nil {
						return err
					}

					for i, careerSkill := range careerSkills {
						skill, err := tx.Skill.Query().Where(skill.ID(skillIDs[i])).WithSkillTag().Only(ctx)
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
		return nil
	}); err != nil {
		return nil, err
	}

	return PostUsersByUserIdCareergroupsByCareerGroupIdCareers201JSONResponse(ToSwaggerUserCareer(entCareer)), nil
}
