package rest

import (
	"context"

	"github.com/sky0621/cv-admin/src/ent"
	"github.com/sky0621/cv-admin/src/ent/careerskill"
	"github.com/sky0621/cv-admin/src/ent/careerskillgroup"
	"github.com/sky0621/cv-admin/src/ent/careertask"
	"github.com/sky0621/cv-admin/src/ent/careertaskdescription"
	"github.com/sky0621/cv-admin/src/ent/helper"
	"github.com/sky0621/cv-admin/src/ent/user"
	"github.com/sky0621/cv-admin/src/ent/useractivity"
	"github.com/sky0621/cv-admin/src/ent/usercareer"
	"github.com/sky0621/cv-admin/src/ent/usercareerdescription"
	"github.com/sky0621/cv-admin/src/ent/usercareergroup"
	"github.com/sky0621/cv-admin/src/ent/usernote"
	"github.com/sky0621/cv-admin/src/ent/usernoteitem"
	"github.com/sky0621/cv-admin/src/ent/userqualification"
)

// DeleteUsersByUserId 指定ユーザー削除
// 指定ユーザーアカウント削除
// (DELETE /users/{byUserId})
func (s *strictServerImpl) DeleteUsersByUserId(ctx context.Context, request DeleteUsersByUserIdRequestObject) (DeleteUsersByUserIdResponseObject, error) {
	if err := helper.WithTransaction(ctx, s.dbClient, func(tx *ent.Tx) error {
		byUserId := request.ByUserId
		/*
		 * user
		 *   |- user_activities
		 *   |- user_qualifications
		 *   |- user_career_groups
		 *   |   |- user_careers
		 *   |       |- user_career_descriptions
		 *   |       |- career_tasks
		 *   |           |- career_task_descriptions
		 *   |       |- career_skill_groups
		 *   |           |- career_skills
		 *   |- user_notes
		 *       |- user_note_items
		 */
		// TODO: 直接 user_id 指定できるはず。。
		careerGroups, err := tx.UserCareerGroup.Query().Where(usercareergroup.HasUserWith(user.ID(byUserId))).All(ctx)
		if err != nil {
			return err
		}
		careerGroupIDs := helper.PickupUserCareerGroupIDs(careerGroups)

		careers, err := tx.UserCareer.Query().Where(usercareer.HasCareerGroupWith(usercareergroup.IDIn(careerGroupIDs...))).All(ctx)
		if err != nil {
			return err
		}
		careerIDs := helper.PickupUserCareerIDs(careers)

		tasks, err := tx.CareerTask.Query().Where(careertask.HasCareerWith(usercareer.IDIn(careerIDs...))).All(ctx)
		if err != nil {
			return err
		}
		taskIDs := helper.PickupCareerTaskIDs(tasks)

		_, err = tx.CareerTaskDescription.Delete().Where(careertaskdescription.HasCareerTaskWith(careertask.IDIn(taskIDs...))).Exec(ctx)
		if err != nil {
			return err
		}

		_, err = tx.CareerTask.Delete().Where(careertask.IDIn(taskIDs...)).Exec(ctx)
		if err != nil {
			return err
		}

		skillGroups, err := tx.CareerSkillGroup.Query().Where(careerskillgroup.HasCareerWith(usercareer.IDIn(careerIDs...))).All(ctx)
		if err != nil {
			return err
		}
		skillGroupIDs := helper.PickupCareerSkillGroupIDs(skillGroups)

		_, err = tx.CareerSkill.Delete().Where(careerskill.HasCareerSkillGroupWith(careerskillgroup.IDIn(skillGroupIDs...))).Exec(ctx)
		if err != nil {
			return err
		}

		_, err = tx.CareerSkillGroup.Delete().Where(careerskillgroup.IDIn(skillGroupIDs...)).Exec(ctx)
		if err != nil {
			return err
		}

		_, err = tx.UserCareerDescription.Delete().Where(usercareerdescription.HasCareerWith(usercareer.IDIn(careerIDs...))).Exec(ctx)
		if err != nil {
			return err
		}

		_, err = tx.UserCareer.Delete().Where(usercareer.IDIn(careerIDs...)).Exec(ctx)
		if err != nil {
			return err
		}

		_, err = tx.UserCareerGroup.Delete().Where(usercareergroup.HasUserWith(user.ID(byUserId))).Exec(ctx)
		if err != nil {
			return err
		}

		notes, err := tx.UserNote.Query().Where(usernote.HasUserWith(user.ID(byUserId))).All(ctx)
		if err != nil {
			return err
		}
		noteIDs := helper.PickupUserNoteIDs(notes)

		_, err = tx.UserNoteItem.Delete().Where(usernoteitem.HasNoteWith(usernote.IDIn(noteIDs...))).Exec(ctx)
		if err != nil {
			return err
		}

		_, err = tx.UserNote.Delete().Where(usernote.HasUserWith(user.ID(byUserId))).Exec(ctx)
		if err != nil {
			return err
		}

		_, err = tx.UserQualification.Delete().Where(userqualification.HasUserWith(user.ID(byUserId))).Exec(ctx)
		if err != nil {
			return err
		}

		_, err = tx.UserActivity.Delete().Where(useractivity.HasUserWith(user.ID(byUserId))).Exec(ctx)
		if err != nil {
			return err
		}

		err = tx.User.DeleteOneID(byUserId).Exec(ctx)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return DeleteUsersByUserId204Response{}, nil
}
