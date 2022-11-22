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
