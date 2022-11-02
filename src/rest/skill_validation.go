package rest

import (
	"errors"
	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/sky0621/cv-admin/src/ent"
	"github.com/sky0621/cv-admin/src/ent/helper"
)

func (u Skill) Validate(skills []*ent.Skill) error {
	err := v.ValidateStruct(&u,
		stringRule(*u.Name, 1, 100, v.Required),
		urlRule(*u.Url, v.NilOrNotEmpty),
	)
	if err != nil {
		return err
	}

	for _, s := range helper.PickupSkillName(skills) {
		if s == *u.Name {
			return errors.New("already registered under the same name")
		}
	}
	return nil
}
