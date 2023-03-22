package pstring

import "strings"

// JoinString ...
func JoinString(input string) string {
	v := strings.Fields(input)
	return strings.Join(v, " ")
}

// IsSpace ...
func IsSpace(v string) bool {
	return len(JoinString(v)) < 1
}
