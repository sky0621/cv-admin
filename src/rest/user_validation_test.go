package rest

import (
	"testing"

	"github.com/sky0621/golang-utils/convert"

	"github.com/stretchr/testify/assert"
)

func testNormalBirthday() *BirthDay {
	return &BirthDay{
		Year:  convert.ToPtr(1900),
		Month: convert.ToPtr(1),
		Day:   convert.ToPtr(1),
	}
}

func TestUserAttribute_validate_NoError(t *testing.T) {
	tests := map[string]struct{ u UserAttribute }{
		"名前がある":   {u: UserAttribute{Name: convert.ToPtr("名前"), Birthday: testNormalBirthday()}},
		"名前が1字":   {u: UserAttribute{Name: convert.ToPtr("x"), Birthday: testNormalBirthday()}},
		"名前が100字": {u: UserAttribute{Name: toPLenStr("x", 100), Birthday: testNormalBirthday()}},

		// TODO: 残りのケース
	}

	for testName, tt := range tests {
		tt := tt
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			err := tt.u.Validate()
			assert.NoError(t, err)
		})
	}
}

func TestUserAttribute_validate_HasError(t *testing.T) {
	tests := map[string]struct {
		u      UserAttribute
		errStr string
	}{
		"名前が無い":       {u: UserAttribute{Birthday: testNormalBirthday()}, errStr: "name: cannot be blank."},
		"名前がnil":      {u: UserAttribute{Name: nil, Birthday: testNormalBirthday()}, errStr: "name: cannot be blank."},
		"名前が空文字":      {u: UserAttribute{Name: convert.ToPtr(""), Birthday: testNormalBirthday()}, errStr: "name: cannot be blank."},
		"名前が100字を超える": {u: UserAttribute{Name: toPLenStr("x", 101), Birthday: testNormalBirthday()}, errStr: "name: the length must be between 1 and 100."},

		// TODO: 残りのケース
	}

	for testName, tt := range tests {
		tt := tt
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			err := tt.u.Validate()
			if assert.Error(t, err) {
				if err.Error() != tt.errStr {
					t.Error(err)
				}
				return
			}
			t.Error(err)
		})
	}
}
