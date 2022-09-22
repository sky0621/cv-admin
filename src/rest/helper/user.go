package helper

import (
	"github.com/sky0621/cv-admin/src/ent"
	"github.com/sky0621/cv-admin/src/swagger"
)

func ConvertSwaggerUserAttributeToEntUserCreate(userAttribute swagger.UserAttribute, c *ent.UserCreate) *ent.UserCreate {
	return c.
		SetName(*userAttribute.Name).
		SetNillableNickname(userAttribute.Nickname).
		SetNillableAvatarURL(userAttribute.AvatarUrl).
		SetBirthdayYear(*userAttribute.Birthday.Year).
		SetBirthdayMonth(*userAttribute.Birthday.Month).
		SetBirthdayDay(*userAttribute.Birthday.Day).
		SetNillableJob(userAttribute.Job).
		SetNillableBelongTo(userAttribute.BelongTo).
		SetNillablePr(userAttribute.Pr)
}

func ConvertSwaggerUserAttributeToEntUserUpdate(userAttribute swagger.UserAttribute, c *ent.UserUpdateOne) *ent.UserUpdateOne {
	return c.
		SetName(*userAttribute.Name).
		SetNillableNickname(userAttribute.Nickname).
		SetNillableAvatarURL(userAttribute.AvatarUrl).
		SetBirthdayYear(*userAttribute.Birthday.Year).
		SetBirthdayMonth(*userAttribute.Birthday.Month).
		SetBirthdayDay(*userAttribute.Birthday.Day).
		SetNillableJob(userAttribute.Job).
		SetNillableBelongTo(userAttribute.BelongTo).
		SetNillablePr(userAttribute.Pr)
}

func ConvertEntUserToSwaggerUserAttribute(user *ent.User) swagger.UserAttribute {
	var userAttribute swagger.UserAttribute
	userAttribute.Id = &user.ID
	userAttribute.Name = &user.Name
	userAttribute.Nickname = user.Nickname
	userAttribute.AvatarUrl = user.AvatarURL
	userAttribute.Birthday = &swagger.BirthDay{
		Year:  &user.BirthdayYear,
		Month: &user.BirthdayMonth,
		Day:   &user.BirthdayDay,
	}
	userAttribute.Job = user.Job
	userAttribute.BelongTo = user.BelongTo
	userAttribute.Pr = user.Pr
	return userAttribute
}
