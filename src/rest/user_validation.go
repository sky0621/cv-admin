package rest

import (
	v "github.com/go-ozzo/ozzo-validation/v4"
)

func (b BirthDay) Validate() error {
	return v.ValidateStruct(&b, yearRule(&b.Year, v.Required), monthRule(&b.Month, v.Required), dayRule(&b.Day, v.Required))
}

func (u UserAttribute) Validate() error {
	return v.ValidateStruct(&u,
		stringRule(&u.Name, 1, 100, v.Required),
		stringRule(&u.Nickname, 1, 100, v.NilOrNotEmpty),
		urlRule(&u.AvatarUrl, v.NilOrNotEmpty),
		requiredRule(&u.Birthday),
		stringRule(&u.Job, 1, 400, v.NilOrNotEmpty),
		stringRule(&u.BelongTo, 1, 400, v.NilOrNotEmpty),
		stringRule(&u.Pr, 1, 4000, v.NilOrNotEmpty),
	)
}

func (u UserActivity) Validate() error {
	return v.ValidateStruct(&u,
		stringRule(&u.Name, 1, 100, v.Required),
		urlRule(&u.Url, v.NilOrNotEmpty),
		stringRule(&u.Icon, 1, 40, v.NilOrNotEmpty),
	)
}

func (u UserQualification) Validate() error {
	return v.ValidateStruct(&u,
		stringRule(&u.Name, 1, 100, v.Required),

		// FIXME: 残りを実装！
	)
}
