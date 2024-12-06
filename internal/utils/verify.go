package utils

import (
	"regexp"
	"strings"
)

func IsValidMethod(method string) bool {
	methods := []string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS", "TRACE"}

	for _, m := range methods {
		if strings.Contains(method, m) {
			return true
		}
	}
	return false
}

func IsValidPath(path string) bool {
	re := regexp.MustCompile(`^[A-Za-z0-9_/\\-]+$`)
	return re.MatchString(path)
}
