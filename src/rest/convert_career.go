package rest

import (
	"fmt"

	"github.com/sky0621/cv-admin/src/ent"
)

func ToSwaggerUserCareerGroup(entCareerGroup *ent.UserCareerGroup) UserCareerGroup {
	var careerGroup UserCareerGroup
	careerGroup.Label = &entCareerGroup.Label
	//	careerGroup.Careers = &entCareerGroup.Edges.Careers
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
	return fmt.Sprintf("%d/%02d", *f.Year, *f.Month)
}

func ToEntUserCareerCreate(u UserCareer, userCareerGroupID int, c *ent.UserCareerCreate) *ent.UserCareerCreate {
	return c.
		SetCareerGroupID(userCareerGroupID).
		SetName(*u.Name)
}
