package rest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCareerPeriodFromToEntUserCareerFrom(t *testing.T) {
	tests := map[string]struct {
		u      *CareerPeriodFrom
		expect string
	}{
		"":             {u: &CareerPeriodFrom{Year: toPInt(2022), Month: toPInt(12)}, expect: "2022/12"},
		"0埋め":          {u: &CareerPeriodFrom{Year: toPInt(2022), Month: toPInt(1)}, expect: "2022/01"},
		"nil":          {u: nil, expect: ""},
		"Year is nil":  {u: &CareerPeriodFrom{Month: toPInt(12)}, expect: ""},
		"Month is nil": {u: &CareerPeriodFrom{Year: toPInt(2022)}, expect: ""},
	}

	for testName, tt := range tests {
		tt := tt
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			actual := tt.u.ToEntUserCareerFrom()
			assert.Equal(t, tt.expect, actual)
		})
	}
}
