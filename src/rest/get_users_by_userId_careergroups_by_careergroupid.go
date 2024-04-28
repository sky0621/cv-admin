package rest

import (
	"context"

	"github.com/sky0621/cv-admin/src/ent"
	"github.com/sky0621/cv-admin/src/ent/user"
	"github.com/sky0621/cv-admin/src/ent/usercareergroup"
)

// GetUsersByUserIdCareergroupsByCareerGroupId
// １キャリア取得
// (GET /users/{byUserId}/careergroups/{byCareerGroupId})
func (s *strictServerImpl) GetUsersByUserIdCareergroupsByCareerGroupId(ctx context.Context, request GetUsersByUserIdCareergroupsByCareerGroupIdRequestObject) (GetUsersByUserIdCareergroupsByCareerGroupIdResponseObject, error) {
	careerGroup, err := s.dbClient.UserCareerGroup.Query().
		Where(usercareergroup.HasUserWith(user.ID(request.ByUserId)), usercareergroup.ID(request.ByCareerGroupId)).WithCareers(func(q *ent.UserCareerQuery) {
		q.WithCareerDescriptions()
		q.WithCareerSkillGroups(func(q *ent.CareerSkillGroupQuery) {
			q.WithCareerSkills(func(q *ent.CareerSkillQuery) {
				q.WithSkill(func(q *ent.SkillQuery) {
					q.WithSkillTag()
				})
			})
		})
		q.WithCareerTasks(func(q *ent.CareerTaskQuery) {
			q.WithCareerTaskDescriptions()
		})
	}).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return GetUsersByUserIdCareergroupsByCareerGroupId200JSONResponse(ToSwaggerUserCareerGroup(careerGroup)), nil
}
