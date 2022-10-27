package rest

import (
	"fmt"
	"strconv"

	"github.com/sky0621/cv-admin/src/ent"
	"github.com/sky0621/golang-utils/convert"
)

// TODO: goverter での置き換えを試す。

func toSwaggerUserCareerPeriod(entCareerPeriod string, careerPeriod CareerPeriod) CareerPeriod {
	if len(entCareerPeriod) != 6 {
		return nil
	}

	year, err := strconv.Atoi(entCareerPeriod[:4])
	if err != nil {
		panic(err) // MEMO: ありえないよう、DB格納時にチェック済み想定
	}
	month, err := strconv.Atoi(entCareerPeriod[4:])
	if err != nil {
		panic(err) // MEMO: ありえないよう、DB格納時にチェック済み想定
	}

	careerPeriod.Set(convert.ToPtr(year), convert.ToPtr(month))
	return careerPeriod
}

func ToSwaggerUserCareerPeriodFrom(entCareerPeriod string) *CareerPeriodFrom {
	p := toSwaggerUserCareerPeriod(entCareerPeriod, &CareerPeriodFrom{})
	if p == nil {
		return nil
	}
	return p.(*CareerPeriodFrom)
}

func ToSwaggerUserCareerPeriodTo(entCareerPeriod *string) *CareerPeriodTo {
	if entCareerPeriod == nil {
		return nil
	}
	p := toSwaggerUserCareerPeriod(*entCareerPeriod, &CareerPeriodTo{})
	if p == nil {
		return nil
	}
	return p.(*CareerPeriodTo)
}

func ToSwaggerUserCareerDescription(entUserCareerDescription *ent.UserCareerDescription) string {
	return entUserCareerDescription.Description
}

func ToSwaggerUserCareerDescriptions(entUserCareerDescriptions []*ent.UserCareerDescription) *[]string {
	var careerDescriptions []string
	for _, entUserCareerDescription := range entUserCareerDescriptions {
		careerDescriptions = append(careerDescriptions, ToSwaggerUserCareerDescription(entUserCareerDescription))
	}
	if careerDescriptions == nil {
		return nil
	}
	return &careerDescriptions
}

func ToSwaggerCareerSkillGroup(entCareerSkillGroup *ent.CareerSkillGroup) *CareerSkillGroup {
	if entCareerSkillGroup == nil {
		return nil
	}
	for _, entCareerSkill := range entCareerSkillGroup.Edges.CareerSkills {
		fmt.Println(entCareerSkill)
	}
	return &CareerSkillGroup{
		Label: &entCareerSkillGroup.Label,
	}
}

func ToSwaggerCareerSkillGroups(entCareerSkillGroups []*ent.CareerSkillGroup) *[]CareerSkillGroup {
	var careerSkillGroups []CareerSkillGroup
	for _, entCareerSkillGroup := range entCareerSkillGroups {
		careerSkillGroups = append(careerSkillGroups, *ToSwaggerCareerSkillGroup(entCareerSkillGroup))
	}
	return &careerSkillGroups
}

func ToSwaggerUserCareer(entCareer *ent.UserCareer) UserCareer {
	return UserCareer{
		Name:        &entCareer.Name,
		Description: ToSwaggerUserCareerDescriptions(entCareer.Edges.CareerDescriptions),
		From:        ToSwaggerUserCareerPeriodFrom(entCareer.From),
		To:          ToSwaggerUserCareerPeriodTo(entCareer.To),
		SkillGroups: ToSwaggerCareerSkillGroups(entCareer.Edges.CareerSkillGroups),
	}
}

func ToSwaggerUserCareerGroup(entCareerGroup *ent.UserCareerGroup) UserCareerGroup {
	careerGroup := UserCareerGroup{}
	careerGroup.Label = &entCareerGroup.Label
	careerGroup.Careers = &[]UserCareer{}
	for _, entCareer := range entCareerGroup.Edges.Careers {
		*careerGroup.Careers = append(*careerGroup.Careers, ToSwaggerUserCareer(entCareer))
	}
	return careerGroup
}

func ToSwaggerUserCareerGroups(entCareerGroups ent.UserCareerGroups) UserCareerGroups {
	var careerGroups UserCareerGroups
	for _, entCareerGroup := range entCareerGroups {
		careerGroups = append(careerGroups, ToSwaggerUserCareerGroup(entCareerGroup))
	}
	return careerGroups
}

func ToEntUserCareerGroupCreate(u UserCareerGroup, userID int, c *ent.UserCareerGroupCreate) *ent.UserCareerGroupCreate {
	return c.
		SetLabel(*u.Label).
		SetUserID(userID)
}

func (f *CareerPeriodFrom) ToEntUserCareerFrom() string {
	if f == nil || f.Year == nil || f.Month == nil {
		return ""
	}
	return fmt.Sprintf("%d%02d", *f.Year, *f.Month)
}

func (f *CareerPeriodTo) ToEntUserCareerTo() string {
	if f == nil || f.Year == nil || f.Month == nil {
		return ""
	}
	return fmt.Sprintf("%d%02d", *f.Year, *f.Month)
}

func ToEntUserCareerCreate(u UserCareer, userCareerGroupID int, c *ent.UserCareerCreate) *ent.UserCareerCreate {
	return c.
		SetCareerGroupID(userCareerGroupID).
		SetName(*u.Name).
		SetFrom(u.From.ToEntUserCareerFrom()).
		SetTo(u.To.ToEntUserCareerTo())
}

func ToEntUserCareerDescriptionCreate(description string, userCareerID int, c *ent.UserCareerDescriptionCreate) *ent.UserCareerDescriptionCreate {
	return c.
		SetCareerID(userCareerID).
		SetDescription(description)
}

func ToEntUserCareerSkillGroup(u CareerSkillGroup, userCareerID int, c *ent.CareerSkillGroupCreate) *ent.CareerSkillGroupCreate {
	return c.
		SetCareerID(userCareerID).
		SetLabel(*u.Label)
}

func ToEntUserCareerSkill(u CareerSkill, careerSkillGroupID int, c *ent.CareerSkillCreate) *ent.CareerSkillCreate {
	return c.
		SetCareerSkillGroupID(careerSkillGroupID).
		SetName(*u.Name).
		SetNillableURL(u.Url).
		SetNillableVersion(u.Version)
}
