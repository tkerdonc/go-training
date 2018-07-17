// This package implements the core functionnality of communication
// sciences, transforming words into acronyms
package acronym

import "strings"

// Generates and acronym our of a string
// parameters
// * inputString : base string for the acronym. Will be split into words
// 				   using ' ' or '-' as a separator. Each word will then
//				   be used to compute the returned acronym
func Abbreviate(inputString string) string {
	// Who knows when we would like another separator...
	separators := [2]string{" ", "-"}
	acronym := ""
	words := []string{inputString}
	splittedWords :=[]string{}

	// Iterate the separators and words, creating a flattened list
	for _, separator := range(separators) {
		splittedWords =[]string{}
		for _, word := range(words) {
			splittedWords = append(splittedWords, strings.Split(word, separator)...)
		}
		words = splittedWords
	}

	// Pick the first letter
	for _, word := range(words) {
		acronym += (word[:1])
	}

	return strings.ToUpper(acronym)
}
