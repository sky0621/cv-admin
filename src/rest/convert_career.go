package rest

import "github.com/sky0621/cv-admin/src/ent"

func ToSwaggerUserCareerGroup(entCareerGroup *ent.UserCareerGroup) UserCareerGroup {
	var careerGroup UserCareerGroup
	careerGroup.Label = &entCareerGroup.Label
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
