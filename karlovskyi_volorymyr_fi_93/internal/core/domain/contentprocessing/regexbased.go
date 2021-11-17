package contentprocessing

import (
	"bytes"
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

func RemoveIndentBytes(p []byte) []byte {
	p = bytes.Trim(p," ")
	return indentPattern.ReplaceAllLiteral(p, []byte{' '})
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

func ReplaceNotAllowedBytes(p []byte) []byte {
	return allowedPattern.ReplaceAllLiteral(p, []byte{' '})
}

func RemovePunctuation(s string) string {
	return RemoveIndent(forRemovingPunctuation.ReplaceAllString(s, " "))
}

func RemovePunctuationBytes(p []byte) []byte {
	return RemoveIndentBytes(forRemovingPunctuation.ReplaceAllLiteral(p, []byte{' '}))
}