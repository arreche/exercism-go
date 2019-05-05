// Package letter computes the frequency of letters in texts using parallel computation.
package letter

// ConcurrentFrequency returns the total frequency of each letter in a list of texts
func ConcurrentFrequency(s []string) FreqMap {
	freq := FreqMap{}

	results := make(chan FreqMap, 10)

	for _, data := range s {
		go func(data string) {
			results <- Frequency(data)
		}(data)
	}

	for range s {
		for k, v := range <-results {
			freq[k] += v
		}
	}

	return freq
}
