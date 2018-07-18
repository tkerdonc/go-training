// This package implements an hamming distance calculation function
// aimed at DNA strands

package hamming

import (
	"errors"
	"strings"
)

/*  Distance returns the hamming distance between two dna strands
    passed as string parameters. Returns -1 if the strands are of
	different lengths.
    parameters :
		a : strand as a string
		b : strand as a string
*/
func Distance(a, b string) (distance int, distanceError error) {
	if len(a) != len(b) {
		distance = -1
		distanceError = errors.New("Distance expects strands to be of equal lengths.")
	} else {
		distanceError = nil
		distance = 0
		var aReader = strings.NewReader(a)
		var bReader = strings.NewReader(b)
		aRune, _, aErr := aReader.ReadRune()
		for aErr == nil {
			bRune, _, _ := bReader.ReadRune()
			if aRune != bRune {
				distance += 1
			}
			aRune, _, aErr = aReader.ReadRune()
		}
	}
	return
}
