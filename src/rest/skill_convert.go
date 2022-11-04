package rest

import "github.com/sky0621/cv-admin/src/ent"

// TODO: goverter での置き換えを試す。

func ToEntSkillCreate(s Skill, c *ent.SkillCreate) *ent.SkillCreate {
	res := c.
		SetName(*s.Name).
		SetKey(*s.Key).
		SetNillableURL(s.Url).
		SetNillableTagKey(s.TagKey)
	return res
}

func ToSwaggerSkill(entSkill *ent.Skill) Skill {
	return Skill{
		Name:   &entSkill.Name,
		Key:    &entSkill.Key,
		Url:    entSkill.URL,
		TagKey: entSkill.TagKey,
	}
}

func ToSwaggerSkills(entSkills []*ent.Skill) []Skill {
	var skills []Skill
	for _, entSkill := range entSkills {
		skills = append(skills, ToSwaggerSkill(entSkill))
	}
	return skills
}
