package rest

import (
	"cmp"
	"context"
	"slices"

	"github.com/sky0621/cv-admin/src/ent/skilltag"

	"github.com/sky0621/cv-admin/src/ent"
	"github.com/sky0621/cv-admin/src/ent/user"
	"github.com/sky0621/cv-admin/src/ent/usercareergroup"
)

// GetUsersByUserIdSkills スキル群取得
// スキル群取得
// (GET /users/{byUserId}/skills)
func (s *strictServerImpl) GetUsersByUserIdSkills(ctx context.Context, request GetUsersByUserIdSkillsRequestObject) (GetUsersByUserIdSkillsResponseObject, error) {
	entSkillTags, err := s.dbClient.SkillTag.Query().Order(ent.Asc(skilltag.FieldOrder)).WithSkills(func(q *ent.SkillQuery) {
		q.WithCareerSkills(func(q *ent.CareerSkillQuery) {
			q.WithCareerSkillGroup(func(q *ent.CareerSkillGroupQuery) {
				q.WithCareer(func(q *ent.UserCareerQuery) {
					q.WithCareerGroup(func(q *ent.UserCareerGroupQuery) {
						q.Where(usercareergroup.HasUserWith(user.ID(request.ByUserId)))
					})
				})
			})
		})
	}).All(ctx)
	if err != nil {
		return nil, err
	}

	var userSkillTags []UserSkillTag
	for _, entSkillTag := range entSkillTags {
		var userSkills []UserSkill
		for _, entSkill := range entSkillTag.Edges.Skills {
			var versions []UserSkillVersion
			for _, entCareerSkill := range entSkill.Edges.CareerSkills {
				from := ToSwaggerUserCareerPeriodFrom(entCareerSkill.Edges.CareerSkillGroup.Edges.Career.From)
				to := ToSwaggerUserCareerPeriodTo(entCareerSkill.Edges.CareerSkillGroup.Edges.Career.To)
				versions = append(versions, UserSkillVersion{
					Version: entCareerSkill.Version,
					From:    from,
					To:      to,
					Period:  toPeriodMonth(from, to),
				})
			}
			userSkills = append(userSkills, UserSkill{
				Name:     &entSkill.Name,
				Url:      entSkill.URL,
				Versions: &versions,
				Period:   sumPeriodMonth(versions),
			})
		}
		slices.SortStableFunc(userSkills, func(a, b UserSkill) int {
			return cmp.Compare(*b.Period, *a.Period)
		})

		userSkillTags = append(userSkillTags, UserSkillTag{
			TagName: &entSkillTag.Name,
			Skills:  &userSkills,
		})
	}

	return GetUsersByUserIdSkills200JSONResponse(userSkillTags), nil
}

func toPeriodMonth(from *CareerPeriodFrom, to *CareerPeriodTo) *CareerPeriodMonth {
	if from == nil || to == nil {
		return nil
	}
	f := int(*from.Year)*12 + *from.Month
	t := int(*to.Year)*12 + *to.Month
	r := t - f + 1
	return &r
}

func sumPeriodMonth(versions []UserSkillVersion) *CareerPeriodMonth {
	sum := 0
	for _, version := range versions {
		sum += *version.Period
	}
	return &sum
}
