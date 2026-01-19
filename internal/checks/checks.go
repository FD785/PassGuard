package checks

import (
	"regexp"
	"strings"
)

func HasUpper(p string) bool {
	return regexp.MustCompile(`[A-Z]`).MatchString(p)
}

func HasLower(p string) bool {
	return regexp.MustCompile(`[a-z]`).MatchString(p)
}

func HasNumber(p string) bool {
	return regexp.MustCompile(`[0-9]`).MatchString(p)
}

func HasSymbol(p string) bool {
	return regexp.MustCompile(`[^a-zA-Z0-9]`).MatchString(p)
}

func IsCommon(p string) bool {
	common := []string{"password", "admin", "welcome", "qwerty", "letmein", "123456"}
	for _, w := range common {
		if strings.Contains(strings.ToLower(p), w) {
			return true
		}
	}
	return false
}
