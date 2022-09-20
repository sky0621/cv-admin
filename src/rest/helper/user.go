package helper

import (
	"github.com/sky0621/cv-admin/src/ent"
	"github.com/sky0621/cv-admin/src/swagger"
)

func ConvertUserAttribute(userAttribute swagger.UserAttribute, c *ent.UserCreate) *ent.UserCreate {
	return c.
		SetKey(*userAttribute.Key).
		SetName(*userAttribute.Name).
		SetNillableNickname(userAttribute.Nickname).
		SetNillableAvatarURL(userAttribute.AvatarUrl).
		SetBirthdayYear(int(*userAttribute.Birthday.Year)).
		SetBirthdayMonth(int(*userAttribute.Birthday.Month)).
		SetBirthdayDay(int(*userAttribute.Birthday.Day)).
		SetNillableJob(userAttribute.Job).
		SetNillableBelongTo(userAttribute.BelongTo).
		SetNillablePr(userAttribute.Pr)
}
