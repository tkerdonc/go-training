// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Exercism package provinding a function adding a gigasecond to an
// input date
package gigasecond

import "time"

const GIGASECOND int = 1000000000;

// AddGigaSecond adds a gigasecond to an input time.
// parameters
// * inputTime : time.Time structure to which this function will add a
//				 Gigasecond
func AddGigasecond(inputTime time.Time) time.Time {
	gigasecondDuration := time.Duration(GIGASECOND) * time.Second
	return inputTime.Add(gigasecondDuration)
}
