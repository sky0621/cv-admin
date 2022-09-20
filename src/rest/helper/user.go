package helper

import (
	"github.com/sky0621/cv-admin/src/ent"
	"github.com/sky0621/cv-admin/src/swagger"
)

func ConvertSwaggerUserAttributeToEntUserCreate(userAttribute swagger.UserAttribute, c *ent.UserCreate) *ent.UserCreate {
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

func ConvertSwaggerUserAttributeToEntUserUpdate(userAttribute swagger.UserAttribute, c *ent.UserUpdate) *ent.UserUpdate {
	return c.
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

func ConvertEntUserToSwaggerUserAttribute(user *ent.User) swagger.UserAttribute {
	var userAttribute swagger.UserAttribute
	userAttribute.Key = &user.Key
	userAttribute.Name = &user.Name
	userAttribute.Nickname = user.Nickname
	userAttribute.AvatarUrl = user.AvatarURL
	year := int32(user.BirthdayYear)
	month := int32(user.BirthdayMonth)
	day := int32(user.BirthdayDay)
	userAttribute.Birthday = &swagger.BirthDay{
		Year:  &year,
		Month: &month,
		Day:   &day,
	}
	userAttribute.Job = user.Job
	userAttribute.BelongTo = user.BelongTo
	userAttribute.Pr = user.Pr
	return userAttribute
}
