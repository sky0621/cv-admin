package rest

import (
	"context"

	"github.com/sky0621/cv-admin/src/ent"
	"github.com/sky0621/cv-admin/src/ent/user"
	"github.com/sky0621/cv-admin/src/ent/usercareer"
	"github.com/sky0621/cv-admin/src/ent/usercareergroup"
)

// GetUsersByUserIdCareergroups キャリアグループ群取得
// キャリアグループ群取得
// (GET /users/{byUserId}/careergroups)
func (s *strictServerImpl) GetUsersByUserIdCareergroups(ctx context.Context, request GetUsersByUserIdCareergroupsRequestObject) (GetUsersByUserIdCareergroupsResponseObject, error) {
	// TODO: 直接 user_id 指定できるはず。。
	careerGroups, err := s.dbClient.UserCareerGroup.Query().
		Where(usercareergroup.HasUserWith(user.ID(request.ByUserId))).WithCareers(func(q *ent.UserCareerQuery) {
		q.WithCareerDescriptions()
		q.WithCareerSkillGroups(func(q *ent.CareerSkillGroupQuery) {
			q.WithCareerSkills(func(q *ent.CareerSkillQuery) {
				q.WithSkill()
			})
		})
		q.WithCareerTasks(func(q *ent.CareerTaskQuery) {
			q.WithCareerTaskDescriptions()
		})
	}).Order(ent.Desc(usercareergroup.FieldCreateTime), ent.Desc(usercareer.FieldCreateTime)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	return GetUsersByUserIdCareergroups200JSONResponse{ToSwaggerUserCareerGroups(careerGroups)}, nil
}
