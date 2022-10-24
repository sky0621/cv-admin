package schema

import (
	"errors"
	"unicode/utf8"
)

// MEMO: 空文字は no error なので、必須チェックには別途 NotEmpty() が必要。
func minRuneCount(minLen int) func(s string) error {
	return func(s string) error {
		if s == "" {
			return nil
		}
		if utf8.RuneCountInString(s) < minLen {
			return errors.New("too short string")
		}
		return nil
	}
}

func maxRuneCount(maxLen int) func(s string) error {
	return func(s string) error {
		if utf8.RuneCountInString(s) > maxLen {
			return errors.New("too long string")
		}
		return nil
	}
}

func rangeRuneCount(minLen, maxLen int) func(s string) error {
	minFunc := minRuneCount(minLen)
	maxFunc := maxRuneCount(maxLen)
	return func(s string) error {
		if err := minFunc(s); err != nil {
			return err
		}
		if err := maxFunc(s); err != nil {
			return err
		}
		return nil
	}
}

func fixedRuneCount(len int) func(s string) error {
	rangeFunc := rangeRuneCount(len, len)
	return func(s string) error {
		if err := rangeFunc(s); err != nil {
			return err
		}
		return nil
	}
}
