package processor

import "strings"

func FormatString(s string, scope scope) string {
	for variable, value := range scope {
		s = strings.ReplaceAll(s, "{"+variable+"}", value)
	}
	return s
}
