//This package implements the Luhn checksum in order to validate a
// numeric string.
package luhn

import "regexp"

/*  Valid implements the Luhn checksum function. Given the target
    string as a parameter, it returns the result of the algorithm as
	a boolean.
	The algorithm iterates over the string and multiplies every
	odd digit by two, then substracts 9 to it if the result was over 9.
	The sum of all these digits is computed, the input is considered
	valid if said sum is divisible by 10.  */
func Valid(input string) (ok bool) {
	ok = false
	digitSum := 0

	spaceStrippingRegex := regexp.MustCompile(" ")
	input = spaceStrippingRegex.ReplaceAllString(input, "")
	isNumeric, _ := regexp.MatchString("^[0-9]", input)

	if len(input) > 1 && isNumeric {
		for i := len(input) - 1; i >= 0; i-- {
			currentInt := int(input[i]) - '0'
			//Matches every second rune starting from the right
			//might be easier to than comparing (len(input) - i)%2 to 0
			if (len(input)-1-i)%2 == 1 {
				currentInt = (currentInt * 2)
				if currentInt > 9 {
					currentInt -= 9
				}
			}

			// Summing digit in the same iteration as doublings
			digitSum += currentInt
		}

		//performing final %10 check
		ok = digitSum%10 == 0

	}
	return
}
