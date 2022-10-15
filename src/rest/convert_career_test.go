package rest

import (
	"testing"

	"github.com/sky0621/golang-utils/convert"

	"github.com/stretchr/testify/assert"
)

func TestToSwaggerUserCareerPeriodFrom(t *testing.T) {
	//
	actual := ToSwaggerUserCareerPeriodFrom("202010")
	if *actual.Year != 2020 {
		t.Fail()
	}
	if *actual.Month != 10 {
		t.Fail()
	}
}

func TestCareerPeriodFromToEntUserCareerFrom(t *testing.T) {
	tests := map[string]struct {
		u      *CareerPeriodFrom
		expect string
	}{
		"":             {u: &CareerPeriodFrom{Year: convert.ToPtr(2022), Month: convert.ToPtr(12)}, expect: "202212"},
		"0埋め":          {u: &CareerPeriodFrom{Year: convert.ToPtr(2022), Month: convert.ToPtr(1)}, expect: "202201"},
		"nil":          {u: nil, expect: ""},
		"Year is nil":  {u: &CareerPeriodFrom{Month: convert.ToPtr(12)}, expect: ""},
		"Month is nil": {u: &CareerPeriodFrom{Year: convert.ToPtr(2022)}, expect: ""},
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
