package luhn

import (
	"strconv"
	"strings"
)

// Valid determines whether a number is valid per the Luhn formula.
func Valid(s string) bool {
	s = strings.TrimSpace(s)

	if len(s) <= 1 {
		return false
	}

	var sum int
	even := false

	for i := range s {
		k := len(s) - i - 1

		char := string(s[k])
		if char == " " {
			continue
		}

		d, err := strconv.Atoi(char)
		if err != nil {
			return false
		}

		if even {
			even = false
			d *= 2
			if d > 9 {
				d -= 9
			}
		} else {
			even = true
		}

		sum += d
	}

	if sum%10 != 0 {
		return false
	}

	return true
}
