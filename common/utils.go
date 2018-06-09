package common

import "strings"

// Index takes in an array of strings, and returns the index of the second argument
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// Includes checks if a given slice includes the requested argument
func Includes(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

// Reverse reverses a string
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// TrimSuffix trims the last of suffix from a string
func TrimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}

// Substring returns only parts of a string
func Substring(s string, start int, end int) string {
	runes := []rune(s)
	stripped := runes[start:end]
	return string(stripped)
}
