package scrabble

import "unicode"

// score computes the scrabble score of a rune
func score(r rune) int {
	switch r {
	case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
		return 1
	case 'D', 'G':
		return 2
	case 'B', 'C', 'M', 'P':
		return 3
	case 'F', 'H', 'V', 'W', 'Y':
		return 4
	case 'K':
		return 5
	case 'J', 'X':
		return 8
	case 'Q', 'Z':
		return 10
	default:
		return 0
	}
}

// Score computes the scrabble score of a word
// i.g. cabbage => 3 + 2*1 + 2*3 + 2 + 1 = 14
func Score(w string) int {
	var total int
	for _, r := range w {
		total += score(unicode.ToUpper(r))
	}
	return total
}
