package contentprocessing

import (
	"regexp"
	"strings"
)

var indentPattern = regexp.MustCompile(`\s+`)
var allowedPattern = regexp.MustCompile(`[^\w\s'."*;<>]+`)
var forRemovingPunctuation = regexp.MustCompile(`[^\sa-zA-Z0-9_]+`)

func RemoveIndent(s string) string {
	s = Trim(s)
	return indentPattern.ReplaceAllString(s, " ")
}

func Trim(s string) string {
	s = strings.TrimLeft(s, " ")
	s = strings.TrimRight(s, " ")
	return s
}

func Filter(s string) string {
	s = Trim(s)
	return allowedPattern.ReplaceAllString(s, " ")
}

func ReplaceNotAllowedChars(s string) string {
	return allowedPattern.ReplaceAllString(s, " ")
}

func RemovePunctuation(s string) string {
	return RemoveIndent(forRemovingPunctuation.ReplaceAllString(s, " "))
}
