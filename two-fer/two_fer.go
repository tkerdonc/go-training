// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package twofer should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package twofer

import "fmt"

// ShareWith returns the following formatted string depending on the 
// person parameter.
// "One for <person>, one for me."
// parameters:
// * person : string to use in the formatting, replaced by "you" if a
//            zero-length string is passed.
func ShareWith(person string) string {
	template := "One for %s, one for me."
	if len(person) == 0 {
		person = "you"
	}
	return fmt.Sprintf(template, person)
}
