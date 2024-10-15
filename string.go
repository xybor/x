package x

import (
	"slices"
	"strconv"
)

func FormatSnowflake(v int64) string {
	if v == 0 {
		return ""
	}

	return strconv.FormatInt(v, 10)
}

func ParseSnowflake(s string) (int64, error) {
	if s == "" {
		return 0, nil
	}

	return strconv.ParseInt(s, 10, 64)
}

func IsNumber(c rune) bool {
	return c >= '0' && c <= '9'
}

func IsLowerCaseLetter(c rune) bool {
	return c >= 'a' && c <= 'z'
}

func IsUpperCaseLetter(c rune) bool {
	return c >= 'A' && c <= 'Z'
}

func IsLetter(c rune) bool {
	return IsLowerCaseLetter(c) || IsUpperCaseLetter(c)
}

func IsWhiteSpace(c rune) bool {
	return c == ' ' || c == '\t' || c == '\n'
}

func IsSpace(c rune) bool {
	return c == ' '
}

func IsUnderscore(c rune) bool {
	return c == '_'
}

func IsSpecialCharacter(c rune) bool {
	return slices.Contains([]rune{
		' ',
		'~', '`', '!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '_', '-', '+', '=',
		'[', ']', '{', '}', '\\', '|',
		':', ';', '"', '\'',
		'<', '>', ',', '.', '?', '/',
	}, c)
}
