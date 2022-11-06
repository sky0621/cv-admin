package rest

import (
	"fmt"
	"strconv"

	"github.com/sky0621/cv-admin/src/ent"
	"github.com/sky0621/golang-utils/convert"
)

// TODO: goverter での置き換えを試す。

func toSwaggerUserCareerPeriod(entCareerPeriod string, careerPeriod CareerPeriodIf) CareerPeriodIf {
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

func ToSwaggerCareerSkill(entCareer *ent.CareerSkill) CareerSkill {
	cs := CareerSkill{
		Version: entCareer.Version,
	}
	entSkill := entCareer.Edges.Skill
	if entSkill == nil {
		return cs
	}
	cs.Skill = &Skill{
		Id:     &entSkill.ID,
		Name:   &entSkill.Name,
		Key:    &entSkill.Key,
		Url:    entSkill.URL,
		TagKey: entSkill.TagKey,
	}
	return cs
}

func ToSwaggerCareerSkillGroup(entCareerSkillGroup *ent.CareerSkillGroup) *CareerSkillGroup {
	if entCareerSkillGroup == nil {
		return nil
	}
	var careerSkills []CareerSkill
	for _, entCareerSkill := range entCareerSkillGroup.Edges.CareerSkills {
		careerSkills = append(careerSkills, ToSwaggerCareerSkill(entCareerSkill))
	}
	return &CareerSkillGroup{
		Label:  &entCareerSkillGroup.Label,
		Skills: &careerSkills,
	}
}

func ToSwaggerCareerTaskDescription(entCareerTaskDescription *ent.CareerTaskDescription) string {
	return entCareerTaskDescription.Description
}

func ToSwaggerCareerTaskDescriptions(entCareerTaskDescriptions []*ent.CareerTaskDescription) *TaskDescription {
	var careerDescriptions TaskDescription
	for _, entTaskDescription := range entCareerTaskDescriptions {
		careerDescriptions = append(careerDescriptions, ToSwaggerCareerTaskDescription(entTaskDescription))
	}
	return &careerDescriptions
}

func ToSwaggerCareerTask(entCareerTask *ent.CareerTask) *CareerTask {
	if entCareerTask == nil {
		return nil
	}
	if entCareerTask.Edges.CareerTaskDescriptions == nil {
		return &CareerTask{
			Name: &entCareerTask.Name,
		}
	}

	var careerDescriptions TaskDescription
	for _, entTaskDescription := range entCareerTask.Edges.CareerTaskDescriptions {
		careerDescriptions = append(careerDescriptions, ToSwaggerCareerTaskDescription(entTaskDescription))
	}
	return &CareerTask{
		Name:        &entCareerTask.Name,
		Description: ToSwaggerCareerTaskDescriptions(entCareerTask.Edges.CareerTaskDescriptions),
	}
}

func ToSwaggerCareerTasks(entCareerTasks []*ent.CareerTask) *[]CareerTask {
	var careerTasks []CareerTask
	for _, entCareerTask := range entCareerTasks {
		careerTasks = append(careerTasks, *ToSwaggerCareerTask(entCareerTask))
	}
	return &careerTasks
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
		Tasks:       ToSwaggerCareerTasks(entCareer.Edges.CareerTasks),
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

func ToSwaggerUserCareerGroups(entCareerGroups ent.UserCareerGroups) []UserCareerGroup {
	var careerGroups []UserCareerGroup
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

func ToEntCareerTaskCreate(u CareerTask, userCareerID int, c *ent.CareerTaskCreate) *ent.CareerTaskCreate {
	return c.
		SetCareerID(userCareerID).
		SetName(*u.Name)
}

func ToEntCareerTaskDescriptionCreate(description string, careerTaskID int, c *ent.CareerTaskDescriptionCreate) *ent.CareerTaskDescriptionCreate {
	return c.
		SetCareerTaskID(careerTaskID).
		SetDescription(description)
}

func ToEntCareerSkillGroupCreate(u CareerSkillGroup, userCareerID int, c *ent.CareerSkillGroupCreate) *ent.CareerSkillGroupCreate {
	return c.
		SetCareerID(userCareerID).
		SetLabel(*u.Label)
}

func ToEntCareerSkillCreate(u CareerSkill, careerSkillGroupID, skillID int, c *ent.CareerSkillCreate) *ent.CareerSkillCreate {
	return c.
		SetCareerSkillGroupID(careerSkillGroupID).
		SetSkillID(skillID).
		SetNillableVersion(u.Version)
}
