package validators

import (
	"regexp"
	"strings"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

func IsEmail(s string) bool {
	s = strings.TrimSpace(s)
	if s == "" {
		return false
	}
	return emailRegex.MatchString(s)
}
