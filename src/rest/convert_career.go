package rest

import (
	"fmt"
	"strconv"

	"github.com/sky0621/cv-admin/src/ent"
)

// TODO: goverter での置き換えを試す。

func ToSwaggerUserCareerPeriodFrom(entFrom string) *CareerPeriodFrom {
	// TODO: validation

	year, err := strconv.Atoi(entFrom[:4])
	if err != nil {
		panic(err) // MEMO: ありえないよう、DB格納時にチェック済み想定
	}
	month, err := strconv.Atoi(entFrom[4:])
	if err != nil {
		panic(err) // MEMO: ありえないよう、DB格納時にチェック済み想定
	}
	return &CareerPeriodFrom{
		Year:  toPInt(year),
		Month: toPInt(month),
	}
}

func ToSwaggerUserCareerDescription(entUserCareerDescription *ent.UserCareerDescription) string {
	return entUserCareerDescription.Description
}

func ToSwaggerUserCareerDescriptions(entUserCareerDescriptions []*ent.UserCareerDescription) *CareerDescriptions {
	careerDescriptions := make([]string, len(entUserCareerDescriptions))
	for _, entUserCareerDescription := range entUserCareerDescriptions {
		careerDescriptions = append(careerDescriptions, ToSwaggerUserCareerDescription(entUserCareerDescription))
	}
	if careerDescriptions == nil {
		return nil
	}
	return &careerDescriptions
}

func ToSwaggerUserCareer(entCareer *ent.UserCareer) UserCareer {
	return UserCareer{
		Name:        &entCareer.Name,
		Description: ToSwaggerUserCareerDescriptions(entCareer.Edges.CareerDescriptions),
		From:        ToSwaggerUserCareerPeriodFrom(entCareer.From),
	}
}

func ToSwaggerUserCareerGroup(entCareerGroup *ent.UserCareerGroup) UserCareerGroup {
	careerGroup := UserCareerGroup{}
	careerGroup.Label = &entCareerGroup.Label
	careerGroup.Careers = &UserCareers{}
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
