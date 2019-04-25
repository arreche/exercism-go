package hamming

import "errors"

// Distance calculates the Hamming difference between two DNA strands.
// It returns an error when the sequences have different lengths.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("the sequences have different lengths")
	}

	var count int
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
	}
	return count, nil
}
