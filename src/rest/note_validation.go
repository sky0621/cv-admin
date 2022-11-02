package rest

import v "github.com/go-ozzo/ozzo-validation/v4"

func (u UserNote) Validate() error {
	return v.ValidateStruct(&u,
		stringRule(*u.Label, 1, 100, v.Required),

		// FIXME: 残りを実装！
	)
}
