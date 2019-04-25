// Package grains implements functions to solve the Wheat and Chessboard Problem.
// http://mathworld.wolfram.com/WheatandChessboardProblem.html
package grains

import (
	"errors"
)

// Square calculates how many grains were on each square. 2^(n-1)
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("squares must be between 1 and 64")
	}
	return 1 << uint(n-1), nil
}

// Total calculates the total number of grains. 2^64-1
func Total() uint64 {
	return 1<<64 - 1
}
