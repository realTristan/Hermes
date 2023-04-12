package hermes

import (
	"regexp"
	"strings"
)

// Check if a string is all alphabetic
func _IsAlphaNum(s string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(s)
}

// Remove double spaces from a string
func _RemoveDoubleSpaces(s string) string {
	for strings.Contains(s, "  ") {
		s = strings.Replace(s, "  ", " ", -1)
	}
	return s
}

// Check if an int is in an array
func _ContainsInt(array []int, value int) bool {
	for i := range array {
		if array[i] == value {
			return true
		}
	}
	return false
}

// Check if a string is in an array
func _ContainsString(array []string, value string) bool {
	for i := range array {
		if array[i] == value {
			return true
		}
	}
	return false
}

// Check if a string contains another string (case insensitive)
func ContainsIgnoreCase(s1 string, s2 string) bool {
	return strings.Contains(strings.ToLower(s1), strings.ToLower(s2))
}
