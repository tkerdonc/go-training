// This package implements the core functionnality of communication
// sciences, transforming words into acronyms
package acronym

import "strings"

// Abbreviate generates an acronym representing a string parameter.
// * inputString : base string for the acronym. Will be split into words
//				   using ' ' or '-' as a separator. Each word will then
//                 be used to compute the returned acronym
func Abbreviate(inputString string) string {
	// Who knows when we would like another separator...
	separators := map[rune]struct{}{' ': {}, '-': {}}
	var lastRune = ' ' // Initialize as a separator to match first rune
	var acronymBuilder = strings.Builder{}

	for _, currentRune := range inputString {
		// Check that last rune was a separator
		if _, isSeparator := separators[lastRune]; isSeparator {
			// Check that current rune is not a separator
			if _, isSeparator := separators[currentRune]; !isSeparator {
				acronymBuilder.Grow(4) // Rune size would be 4 bytes right ?
				acronymBuilder.WriteRune(currentRune)
			}
		}
		lastRune = currentRune
	}

	return strings.ToUpper(acronymBuilder.String())
}
