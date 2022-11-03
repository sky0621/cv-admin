package rest

import (
	v "github.com/go-ozzo/ozzo-validation/v4"
)

func (u Skill) Validate() error {
	return v.ValidateStruct(&u,
		stringRule(&u.Name, 1, 100, v.Required),
		stringRule(&u.Key, 1, 40, v.Required),
		urlRule(&u.Url, v.NilOrNotEmpty),
		stringRule(&u.TagKey, 1, 40, v.NilOrNotEmpty),
	)
}
