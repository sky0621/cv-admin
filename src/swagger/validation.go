package swagger

import (
	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func (b BirthDay) Validate() error {
	return v.ValidateStruct(&b,
		v.Field(
			&b.Year,
			v.Required,
			v.Min(1900),
			v.Max(3000),
		),
		v.Field(
			&b.Month,
			v.Required,
			v.Min(1),
			v.Max(12),
		),
		v.Field(
			&b.Day,
			v.Required,
			v.Min(1),
			v.Max(31),
		),
	)
}

func (u UserAttribute) Validate() error {
	return v.ValidateStruct(&u,
		v.Field(
			&u.Key,
			v.Required,
			v.RuneLength(1, 20),
			is.LowerCase,
			is.Alphanumeric,
		),
		v.Field(
			&u.Name,
			v.Required,
			v.RuneLength(1, 100),
		),
		v.Field(
			&u.Nickname,
			v.NilOrNotEmpty,
			v.RuneLength(1, 100),
		),
		v.Field(
			&u.AvatarUrl,
			v.NilOrNotEmpty,
			v.RuneLength(1, 4096),
			is.URL,
		),
		v.Field(
			&u.Birthday,
			v.Required,
		),
		v.Field(
			&u.Job,
			v.NilOrNotEmpty,
			v.RuneLength(1, 400),
		),
		v.Field(
			&u.BelongTo,
			v.NilOrNotEmpty,
			v.RuneLength(1, 400),
		),
		v.Field(
			&u.Pr,
			v.NilOrNotEmpty,
			v.RuneLength(1, 4000),
		),
	)
}
