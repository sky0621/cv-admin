package converter

import (
	"strconv"
	"strings"
	"time"

	"github.com/sky0621/cv-admin/src/ent"
	"github.com/sky0621/cv-admin/src/swagger"
)

func ToEntUserCreate(u swagger.UserAttribute, c *ent.UserCreate) *ent.UserCreate {
	return c.
		SetName(*u.Name).
		SetNillableNickname(u.Nickname).
		SetNillableAvatarURL(u.AvatarUrl).
		SetBirthdayYear(*u.Birthday.Year).
		SetBirthdayMonth(*u.Birthday.Month).
		SetBirthdayDay(*u.Birthday.Day).
		SetNillableJob(u.Job).
		SetNillableBelongTo(u.BelongTo).
		SetNillablePr(u.Pr)
}

func ToEntUserUpdate(ua swagger.UserAttribute, c *ent.UserUpdateOne) *ent.UserUpdateOne {
	return c.
		SetName(*ua.Name).
		SetNillableNickname(ua.Nickname).
		SetNillableAvatarURL(ua.AvatarUrl).
		SetBirthdayYear(*ua.Birthday.Year).
		SetBirthdayMonth(*ua.Birthday.Month).
		SetBirthdayDay(*ua.Birthday.Day).
		SetNillableJob(ua.Job).
		SetNillableBelongTo(ua.BelongTo).
		SetNillablePr(ua.Pr)
}

func ToSwaggerUserAttribute(u *ent.User) swagger.UserAttribute {
	var ua swagger.UserAttribute
	ua.Id = &u.ID
	ua.Name = &u.Name
	ua.Nickname = u.Nickname
	ua.AvatarUrl = u.AvatarURL
	ua.Birthday = &swagger.BirthDay{
		Year:  &u.BirthdayYear,
		Month: &u.BirthdayMonth,
		Day:   &u.BirthdayDay,
	}
	ua.Job = u.Job
	ua.BelongTo = u.BelongTo
	ua.Pr = u.Pr
	return ua
}

func ToEntUserActivityCreate(ua swagger.UserActivity, userID int, c *ent.UserActivityCreate) *ent.UserActivityCreate {
	return c.
		SetName(*ua.Name).
		SetNillableURL(ua.Url).
		SetNillableIcon(ua.Icon).
		SetUserID(userID)
}

func ToSwaggerUserActivity(entActivity *ent.UserActivity) swagger.UserActivity {
	var activity swagger.UserActivity
	activity.Name = &entActivity.Name
	activity.Url = entActivity.URL
	activity.Icon = entActivity.Icon
	return activity
}

func ToSwaggerUserActivities(entActivities []*ent.UserActivity) swagger.UserActivities {
	var activities swagger.UserActivities
	for _, entActivity := range entActivities {
		activities = append(activities, ToSwaggerUserActivity(entActivity))
	}
	return activities
}

func ToEntUserQualificationCreate(ua swagger.UserQualification, userID int, c *ent.UserQualificationCreate) *ent.UserQualificationCreate {
	return c.
		SetName(*ua.Name).
		SetNillableOrganization(ua.Organization).
		SetNillableURL(ua.Url).
		SetGotDate(ua.GotDate.String()).
		SetNillableMemo(ua.Memo).
		SetUserID(userID)
}

func ToSwaggerUserQualificationGotDate(entGotDate *string) *swagger.QualificationGotDate {
	entGotDateStr := *entGotDate
	dates := strings.Split(entGotDateStr, "-")
	if len(dates) != 3 {
		return nil
	}
	toInt := func(s string) int {
		n, _ := strconv.Atoi(s)
		return n
	}
	tz, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil
	}
	return &swagger.QualificationGotDate{Time: time.Date(toInt(dates[0]), time.Month(toInt(dates[1])), toInt(dates[2]), 0, 0, 0, 0, tz)}
}

func ToSwaggerUserQualification(entQualification *ent.UserQualification) swagger.UserQualification {
	var qualification swagger.UserQualification
	qualification.Name = &entQualification.Name
	qualification.Organization = entQualification.Organization
	qualification.Url = entQualification.Organization
	qualification.GotDate = ToSwaggerUserQualificationGotDate(entQualification.GotDate)
	qualification.Memo = entQualification.Memo
	return qualification
}

func ToSwaggerUserQualifications(entQualifications []*ent.UserQualification) swagger.UserQualifications {
	var qualifications swagger.UserQualifications
	for _, entQualification := range entQualifications {
		qualifications = append(qualifications, ToSwaggerUserQualification(entQualification))
	}
	return qualifications
}
