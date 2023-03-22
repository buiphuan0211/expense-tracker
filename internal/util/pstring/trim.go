package pstring

import "strings"

// TrimSpace ...
func TrimSpace(s string) string {
	return strings.TrimSpace(s)
}

// TrimSpaceAll ...
func TrimSpaceAll(s string) string {
	return strings.TrimSpace(ReplaceMultipleSpace(s))
}
