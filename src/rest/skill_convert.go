package rest

import "github.com/sky0621/cv-admin/src/ent"

// TODO: goverter での置き換えを試す。

func ToEntSkillCreate(ua Skill, c *ent.SkillCreate) *ent.SkillCreate {
	return c.
		SetName(*ua.Name).
		SetNillableURL(ua.Url)
}
