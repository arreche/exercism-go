package diffsquares

// SquareOfSum calculates the square of the sum of the first N natural numbers
// E.g. (1 + 2 + ... + 10)² = 55² = 3025.
func SquareOfSum(n int) int {
	sum := n * (n + 1) / 2
	return sum * sum
}

// SumOfSquares calculates the sum of the squares of the first N natural numbers
// E.g. 1² + 2² + ... + 10² = 385.
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference finds the difference between the square of the sum and the sum of the squares of the first N natural numbers.
// E.g. 3025 - 385 = 2640.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
