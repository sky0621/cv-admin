package rest

import (
	"github.com/sky0621/cv-admin/src/converter"
	"github.com/sky0621/cv-admin/src/ent"
)

// TODO: goverter での置き換えを試す。

func ToEntSkillTagCreate(s SkillTag, c *ent.SkillTagCreate) *ent.SkillTagCreate {
	return c.
		SetName(*s.Name).
		SetKey(*s.Key)
}

func ToSwaggerSkillTag(entSkillTag *ent.SkillTag) SkillTag {
	return SkillTag{
		SkillTagID: &entSkillTag.ID,
		Name:       &entSkillTag.Name,
		Key:        &entSkillTag.Key,
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
		SetKey(*s.Key).
		SetNillableURL(s.Url).
		SetNillableTagKey(s.TagKey)
}

func ToSwaggerSkill(entSkill *ent.Skill) Skill {
	return Skill{
		SkillID: converter.ToPtr(entSkill.ID),
		Name:    &entSkill.Name,
		Key:     &entSkill.Key,
		Url:     entSkill.URL,
		TagKey:  entSkill.TagKey,
	}
}

func ToSwaggerSkills(entSkills []*ent.Skill) []Skill {
	var skills []Skill
	for _, entSkill := range entSkills {
		skills = append(skills, ToSwaggerSkill(entSkill))
	}
	return skills
}
