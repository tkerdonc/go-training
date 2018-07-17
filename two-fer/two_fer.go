// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package twofer should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package twofer

import "fmt"

// Returns the following formatted string depending on the entryString parameter
// "One for <entryString>, one for me."
// parameters:
// * entryString : string to use in the formatting, replaced by "you" if a
// zero-length string is passed
func ShareWith(entryString string) string {
	if entryString == "" {
		entryString = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", entryString)
}
