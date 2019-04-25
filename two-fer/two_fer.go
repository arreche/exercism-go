// Package twofer implements a single function to solve the Two Fer problem.
package twofer

// ShareWith accept a name and returns a string in the following cases:
// - If the given name is "Alice", the result should be "One for Alice, one for me."
// - If no name is given, the result should be "One for you, one for me."
func ShareWith(name string) string {
	if "" == name {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
