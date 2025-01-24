package rest

import (
	"cmp"
	"context"
	"log"
	"slices"
	"sort"

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
		if entSkillTag == nil {
			log.Printf("entSkillTag is nil\n")
			continue
		}
		var userSkills []UserSkill
		for _, entSkill := range entSkillTag.Edges.Skills {
			if entSkill == nil {
				log.Printf("entSkillTag[ID:%d][Name:%s] entSkill == nil\n", entSkillTag.ID, entSkillTag.Name)
				continue
			}
			var versions []UserSkillVersion
			for _, entCareerSkill := range entSkill.Edges.CareerSkills {
				if entCareerSkill == nil {
					log.Printf("entSkill[ID:%d][Name:%s] entCareerSkill == nil\n", entSkill.ID, entSkill.Name)
					continue
				}
				if entCareerSkill.Edges.CareerSkillGroup == nil {
					log.Printf("entCareerSkill[ID:%d] entCareerSkill.Edges.CareerSkillGroup == nil\n", entCareerSkill.ID)
					continue
				}
				if entCareerSkill.Edges.CareerSkillGroup.Edges.Career == nil {
					log.Printf("entCareerSkill[ID:%d] entCareerSkill.Edges.CareerSkillGroup.Edges.Career == nil\n", entCareerSkill.ID)
					continue
				}
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
	// Versionsを新しいもの順にソート
	for _, userSkillTag := range userSkillTags {
		if userSkillTag.Skills != nil {
			for _, skill := range *userSkillTag.Skills {
				if skill.Versions != nil {
					sort.SliceStable(*skill.Versions, func(i, j int) bool {
						vi := (*skill.Versions)[i]
						vj := (*skill.Versions)[j]

						// FromのYearで比較
						if vi.From.Year != nil && vj.From.Year != nil {
							if *vi.From.Year != *vj.From.Year {
								return *vi.From.Year > *vj.From.Year
							}
						}
						// FromのMonthで比較
						if vi.From.Month != nil && vj.From.Month != nil {
							if *vi.From.Month != *vj.From.Month {
								return *vi.From.Month > *vj.From.Month
							}
						}
						// ToのYearで比較
						if vi.To.Year != nil && vj.To.Year != nil {
							if *vi.To.Year != *vj.To.Year {
								return *vi.To.Year > *vj.To.Year
							}
						}
						// ToのMonthで比較
						if vi.To.Month != nil && vj.To.Month != nil {
							return *vi.To.Month > *vj.To.Month
						}

						return false
					})
				}
			}
		}
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
