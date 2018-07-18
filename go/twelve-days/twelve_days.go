//This package sings the twelve days christmas carol
package twelve

import "strings"
import "fmt"

var ORDINALS = map[int]string{
	1:  "first",
	2:  "second",
	3:  "third",
	4:  "fourth",
	5:  "fifth",
	6:  "sixth",
	7:  "seventh",
	8:  "eighth",
	9:  "ninth",
	10: "tenth",
	11: "eleventh",
	12: "twelfth",
}

var PRESENTS = map[int]string{
	1:  " a Partridge in a Pear Tree",
	2:  " two Turtle Doves",
	3:  " three French Hens",
	4:  " four Calling Birds",
	5:  " five Gold Rings",
	6:  " six Geese-a-Laying",
	7:  " seven Swans-a-Swimming",
	8:  " eight Maids-a-Milking",
	9:  " nine Ladies Dancing",
	10: " ten Lords-a-Leaping",
	11: " eleven Pipers Piping",
	12: " twelve Drummers Drumming",
}

const AND_STRING = " and"

// Verse takes the verse number as an integer parameter, and returns the
// corresponding verse as a string.
func Verse(verseIndex int) string {
	verse := ""

	verseTemplate := "On the %s day of Christmas my true love gave to me%s."

	if verseIndex < 13 && verseIndex > 0 {
		ordinal := ORDINALS[verseIndex]
		var enumerationBuilder = strings.Builder{}

		for i := verseIndex; i > 0; i-- {
			enumerationBuilder.WriteRune(',')
			if verseIndex > 1 && i == 1 {
				enumerationBuilder.WriteString(AND_STRING)
			}
			present := PRESENTS[i]
			enumerationBuilder.WriteString(present)
		}

		verse = fmt.Sprintf(verseTemplate, ordinal, enumerationBuilder.String())
	}
	return verse
}

// Song returns the entire 12 verses of the twelve day christmas carol
func Song() string {
	var songBuilder = strings.Builder{}
	for i := 1; i < 13; i++ {
		verse := Verse(i)
		songBuilder.WriteString(verse)
		songBuilder.WriteRune('\n')
	}
	return songBuilder.String()
}
