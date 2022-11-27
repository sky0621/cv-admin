package cmd

import (
	"fmt"
	"testing"
	"time"
)

func TestAge(t *testing.T) {
	nowFn := func(y, m, d int) time.Time {
		return time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.Local)
	}
	tests := []struct {
		name     string
		now      time.Time
		expected int
	}{
		{name: "a", now: nowFn(2022, 5, 1), expected: 43},
		{name: "a", now: nowFn(2022, 5, 2), expected: 43},
		{name: "a", now: nowFn(2022, 5, 3), expected: 43},
		{name: "a", now: nowFn(2022, 5, 30), expected: 43},
		{name: "a", now: nowFn(2022, 5, 31), expected: 43},
		{name: "b", now: nowFn(2022, 6, 1), expected: 43},
		{name: "c", now: nowFn(2022, 6, 2), expected: 44},
		{name: "c", now: nowFn(2022, 6, 3), expected: 44},
		{name: "c", now: nowFn(2022, 7, 1), expected: 44},
		{name: "c", now: nowFn(2022, 7, 2), expected: 44},
		{name: "c", now: nowFn(2022, 7, 3), expected: 44},
	}
	for idx, tt := range tests {
		t.Run(fmt.Sprintf("[No.%d] %s", idx+1, tt.name), func(t *testing.T) {
			actual := age(1978, 6, 2, tt.now)
			if actual != tt.expected {
				t.Errorf("actual: %d, expected: %d", actual, tt.expected)
			}
		})
	}
}

func TestDiffMonths(t *testing.T) {
	tests := []struct {
		name      string
		fromYear  int
		fromMonth int
		toYear    int
		toMonth   int
		expected  int
	}{
		{name: "同一年内 a", fromYear: 2022, fromMonth: 1, toYear: 2022, toMonth: 1, expected: 1},
		{name: "同一年内 b", fromYear: 2022, fromMonth: 1, toYear: 2022, toMonth: 2, expected: 2},
		{name: "同一年内 c", fromYear: 2022, fromMonth: 6, toYear: 2022, toMonth: 11, expected: 6},
		{name: "同一年内 d", fromYear: 2022, fromMonth: 1, toYear: 2022, toMonth: 12, expected: 12},
		{name: "異年（１年差）a", fromYear: 2021, fromMonth: 1, toYear: 2022, toMonth: 1, expected: 13},
		{name: "異年（１年差）b", fromYear: 2021, fromMonth: 1, toYear: 2022, toMonth: 2, expected: 14},
		{name: "異年（１年差）c", fromYear: 2021, fromMonth: 6, toYear: 2022, toMonth: 11, expected: 18},
		{name: "異年（１年差）d", fromYear: 2021, fromMonth: 1, toYear: 2022, toMonth: 12, expected: 24},
		{name: "異年（１年差）e", fromYear: 2021, fromMonth: 11, toYear: 2022, toMonth: 6, expected: 8},
		{name: "異年（２年差）a", fromYear: 2020, fromMonth: 1, toYear: 2022, toMonth: 1, expected: 25},
		{name: "異年（２年差）b", fromYear: 2020, fromMonth: 1, toYear: 2022, toMonth: 2, expected: 26},
		{name: "異年（２年差）c", fromYear: 2020, fromMonth: 6, toYear: 2022, toMonth: 11, expected: 30},
		{name: "異年（２年差）d", fromYear: 2020, fromMonth: 1, toYear: 2022, toMonth: 12, expected: 36},
		{name: "異年（２年差）e", fromYear: 2020, fromMonth: 11, toYear: 2022, toMonth: 6, expected: 20},
	}
	for idx, tt := range tests {
		t.Run(fmt.Sprintf("[No.%d] %s", idx+1, tt.name), func(t *testing.T) {
			actual := diffMonths(tt.fromYear, tt.fromMonth, tt.toYear, tt.toMonth)
			if actual != tt.expected {
				t.Errorf("actual: %d, expected: %d", actual, tt.expected)
			}
		})
	}
}
func TestMergeMonths(t *testing.T) {
	tests := []struct {
		name     string
		dms      int
		expected string
	}{
		{name: "a", dms: 0, expected: "0ヶ月"},
		{name: "b", dms: 1, expected: "1ヶ月"},
		{name: "c", dms: 2, expected: "2ヶ月"},
		{name: "c", dms: 11, expected: "11ヶ月"},
		{name: "d", dms: 12, expected: "1年"},
		{name: "e", dms: 13, expected: "1年1ヶ月"},
		{name: "f", dms: 23, expected: "1年11ヶ月"},
		{name: "g", dms: 24, expected: "2年"},
		{name: "h", dms: 25, expected: "2年1ヶ月"},
	}
	for idx, tt := range tests {
		t.Run(fmt.Sprintf("[No.%d] %s", idx+1, tt.name), func(t *testing.T) {
			actual := mergeMonths(tt.dms)
			if actual != tt.expected {
				t.Errorf("actual: %s, expected: %s", actual, tt.expected)
			}
		})
	}
}
