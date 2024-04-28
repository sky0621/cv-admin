package rest

import (
	"context"

	"github.com/sky0621/cv-admin/src/ent"

	"github.com/sky0621/cv-admin/src/ent/usercareer"
)

// GetUsersByUserIdCareergroupsByCareerGroupIdCareersByCareerId
// １キャリア取得
// (GET /users/{byUserId}/careergroups/{byCareerGroupId}/careers/{byCareerId})
func (s *strictServerImpl) GetUsersByUserIdCareergroupsByCareerGroupIdCareersByCareerId(ctx context.Context, request GetUsersByUserIdCareergroupsByCareerGroupIdCareersByCareerIdRequestObject) (GetUsersByUserIdCareergroupsByCareerGroupIdCareersByCareerIdResponseObject, error) {
	career, err := s.dbClient.UserCareer.Query().
		Where(usercareer.ID(request.ByCareerGroupId)).
		WithCareerGroup().
		WithCareerDescriptions().
		WithCareerSkillGroups(func(q *ent.CareerSkillGroupQuery) {
			q.WithCareerSkills(func(q *ent.CareerSkillQuery) {
				q.WithSkill(func(q *ent.SkillQuery) {
					q.WithSkillTag()
				})
			})
		}).
		WithCareerTasks(func(q *ent.CareerTaskQuery) {
			q.WithCareerTaskDescriptions()
		}).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return GetUsersByUserIdCareergroupsByCareerGroupIdCareersByCareerId200JSONResponse(ToSwaggerUserCareer(career)), nil
}
