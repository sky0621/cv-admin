package rest

import (
	"github.com/sky0621/cv-admin/src/ent"
)

// TODO: goverter での置き換えを試す。

func ToEntSkillTagCreate(s SkillTag, c *ent.SkillTagCreate) *ent.SkillTagCreate {
	return c.
		SetName(*s.Name)
}

func ToSwaggerSkillTag(entSkillTag *ent.SkillTag) SkillTag {
	return SkillTag{
		SkillTagID: &entSkillTag.ID,
		Name:       &entSkillTag.Name,
	}
}

func ToSwaggerSkillTags(entSkillTags []*ent.SkillTag) []SkillTag {
	var skillTags []SkillTag
	for _, entSkillTag := range entSkillTags {
		skillTags = append(skillTags, ToSwaggerSkillTag(entSkillTag))
	}
	return skillTags
}

func ToEntSkillCreate(s Skill, st *ent.SkillTag, c *ent.SkillCreate) *ent.SkillCreate {
	return c.
		SetName(*s.Name).
		SetNillableURL(s.Url).
		SetSkillTagID(*s.SkillTagID).
		SetSkillTag(st)
}

func ToSwaggerSkill(entSkill ent.Skill) Skill {
	s := Skill{
		SkillID: ToPtr(entSkill.ID),
		Name:    &entSkill.Name,
		Url:     entSkill.URL,
	}
	if entSkill.Edges.SkillTag != nil {
		s.SkillTagID = &entSkill.Edges.SkillTag.ID
	}
	return s
}

func ToSwaggerSkills(entSkills []*ent.Skill) []Skill {
	var skills []Skill
	for _, entSkill := range entSkills {
		if entSkill == nil {
			continue
		}
		skills = append(skills, ToSwaggerSkill(*entSkill))
	}
	return skills
}
