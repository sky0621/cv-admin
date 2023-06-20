package rest

import (
	"context"

	"github.com/sky0621/cv-admin/src/ent"
	"github.com/sky0621/cv-admin/src/ent/skill"
	"github.com/sky0621/cv-admin/src/ent/user"
	"github.com/sky0621/cv-admin/src/ent/usercareergroup"
)

// GetUsersByUserIdSkills スキル群取得
// スキル群取得
// (GET /users/{byUserId}/skills)
func (s *strictServerImpl) GetUsersByUserIdSkills(ctx context.Context, request GetUsersByUserIdSkillsRequestObject) (GetUsersByUserIdSkillsResponseObject, error) {
	entSkillTags, err := s.dbClient.SkillTag.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	var userSkillTags []UserSkillTag
	for _, entSkillTag := range entSkillTags {
		entSkills, err := s.dbClient.Skill.Query().Where(skill.TagKey(entSkillTag.Key)).WithCareerSkills(func(q *ent.CareerSkillQuery) {
			q.WithCareerSkillGroup(func(q *ent.CareerSkillGroupQuery) {
				q.WithCareer(func(q *ent.UserCareerQuery) {
					q.WithCareerGroup(func(q *ent.UserCareerGroupQuery) {
						q.Where(usercareergroup.HasUserWith(user.ID(request.ByUserId)))
					})
				})
			})
		}).All(ctx)
		if err != nil {
			return nil, err
		}

		var userSkills []UserSkill
		for _, entSkill := range entSkills {
			var versions []UserSkillVersion
			for _, entCareerSkill := range entSkill.Edges.CareerSkills {
				versions = append(versions, UserSkillVersion{
					Version: entCareerSkill.Version,
					From:    ToSwaggerUserCareerPeriodFrom(entCareerSkill.Edges.CareerSkillGroup.Edges.Career.From),
					To:      ToSwaggerUserCareerPeriodTo(entCareerSkill.Edges.CareerSkillGroup.Edges.Career.To),
				})
			}
			userSkills = append(userSkills, UserSkill{
				Name:     &entSkill.Name,
				Key:      &entSkill.Key,
				Url:      entSkill.URL,
				Versions: &versions,
			})
		}

		userSkillTags = append(userSkillTags, UserSkillTag{
			TagName: &entSkillTag.Name,
			TagKey:  &entSkillTag.Key,
			Skills:  &userSkills,
		})
	}

	return GetUsersByUserIdSkills200JSONResponse{userSkillTags}, nil
}