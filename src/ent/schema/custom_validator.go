package schema

import (
	"errors"
	"unicode/utf8"
)

func maxRuneCount(maxLen int) func(s string) error {
	return func(s string) error {
		if utf8.RuneCountInString(s) > maxLen {
			return errors.New("too long string")
		}
		return nil
	}
}
