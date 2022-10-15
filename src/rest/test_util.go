package rest

import "strings"

func toPLenStr(s string, len int) *string {
	sb := strings.Builder{}
	for i := 0; i < len; i++ {
		sb.WriteString(s)
	}
	sbs := sb.String()
	return &sbs
}
