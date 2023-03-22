package pstring

import "regexp"

// ReplaceMultipleSpace ...
func ReplaceMultipleSpace(value string) (result string) {
	re := regexp.MustCompile(`\s+`)
	out := re.ReplaceAllString(value, " ")
	return out
}

// replaceStringWithRegex ...
func replaceStringWithRegex(src string, regex string, replaceText string) string {
	reg := regexp.MustCompile(regex)
	return reg.ReplaceAllString(src, replaceText)
}
