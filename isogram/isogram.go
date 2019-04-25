package isogram

import "unicode"

// IsIsogram determines if a word or phrase is an isogram.
func IsIsogram(s string) bool {
	occurrences := make(map[rune]bool, len(s))

	for _, r := range s {

		if r == ' ' || r == '-' {
			continue
		}

		char := unicode.ToUpper(r)

		if occurrences[char] {
			return false
		}

		occurrences[char] = true
	}

	return true
}
