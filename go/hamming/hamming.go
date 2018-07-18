// This package implements an hamming distance calculation function
// aimed at DNA strands

package hamming

import (
	"errors"
)

/*	Distance returns the hamming distance between two dna strands
	passed as string parameters. Returns -1 if the strands are of
	different lengths.
	parameters :
		a : strand as a string
		b : strand as a string
*/
func Distance(a, b string) (int, error) {
	distance := 0
	if len(a) != len(b) {
		var distanceError = errors.New("Distance expects strands to be of equal lengths.")
		return distance, distanceError
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance += 1
		}
	}

	return distance, nil
}
