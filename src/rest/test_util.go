package rest

import "strings"

func toPStr(s string) *string {
	return &s
}
func toPLenStr(s string, len int) *string {
	sb := strings.Builder{}
	for i := 0; i < len; i++ {
		sb.WriteString(s)
	}
	sbs := sb.String()
	return &sbs
}

func toPInt(i int) *int {
	return &i
}
