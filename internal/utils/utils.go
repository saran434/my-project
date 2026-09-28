package utils

import "strings"

func TrimAndLower(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}

func IsBlank(s string) bool {
	return strings.TrimSpace(s) == ""
}
