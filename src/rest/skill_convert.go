package rest

import (
	"github.com/sky0621/cv-admin/src/ent"
)

// TODO: goverter での置き換えを試す。

func ToEntSkillTagCreate(s SkillTag, c *ent.SkillTagCreate) *ent.SkillTagCreate {
	return c.
		SetName(*s.Name).
		SetCode(*s.Code)
}

func ToSwaggerSkillTag(entSkillTag *ent.SkillTag) SkillTag {
	return SkillTag{
		SkillTagID: &entSkillTag.ID,
		Name:       &entSkillTag.Name,
		Code:       &entSkillTag.Code,
	}
}

func ToSwaggerSkillTags(entSkillTags []*ent.SkillTag) []SkillTag {
	var skillTags []SkillTag
	for _, entSkillTag := range entSkillTags {
		skillTags = append(skillTags, ToSwaggerSkillTag(entSkillTag))
	}
	return skillTags
}

func ToEntSkillCreate(s Skill, c *ent.SkillCreate) *ent.SkillCreate {
	return c.
		SetName(*s.Name).
		SetCode(*s.Code).
		SetNillableURL(s.Url)
}

func ToSwaggerSkill(entSkill *ent.Skill) Skill {
	return Skill{
		SkillID: ToPtr(entSkill.ID),
		Name:    &entSkill.Name,
		Code:    &entSkill.Code,
		Url:     entSkill.URL,
	}
}

func ToSwaggerSkills(entSkills []*ent.Skill) []Skill {
	var skills []Skill
	for _, entSkill := range entSkills {
		skills = append(skills, ToSwaggerSkill(entSkill))
	}
	return skills
}
