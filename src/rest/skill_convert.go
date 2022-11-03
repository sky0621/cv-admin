package rest

import "github.com/sky0621/cv-admin/src/ent"

// TODO: goverter での置き換えを試す。

func ToEntSkillCreate(ua Skill, c *ent.SkillCreate) *ent.SkillCreate {
	return c.
		SetName(*ua.Name).
		SetNillableURL(ua.Url)
}

func ToSwaggerSkill(entSkill *ent.Skill) Skill {
	return Skill{
		Name: &entSkill.Name,
		Url:  entSkill.URL,
	}
}

func ToSwaggerSkills(entSkills []*ent.Skill) []Skill {
	var skills []Skill
	for _, entQualification := range entSkills {
		skills = append(skills, ToSwaggerSkill(entQualification))
	}
	return skills
}
