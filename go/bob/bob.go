// This package implements a teenager.
package bob

import (
	"regexp"
)

/* Hey implements the teenager's interaction function, generating its
 response string depending on the remark addressed to it.
 parameters
 * remark : input string the teenager is responding to
			Each of the following case triggers a different response:
							- zero length or silent remark
							- shouting (all upper case)
							- question (ends with a '?')
							- shouting and question
							- none of the above */
func Hey(remark string) string {
	silentCharacterRegex := regexp.MustCompile("[\t \n\r]")
	nonSilences := silentCharacterRegex.ReplaceAllString(remark, "")
	if len(nonSilences) == 0 {
		// This is a non localized teenager, responses are hardcoded
		return "Fine. Be that way!"
	}
	isQuestion, _ := regexp.MatchString(".*\\?[ \t\n\r]*$", remark)
	hasLowerCase, _ := regexp.MatchString("[a-z]", remark)
	hasUpperCase, _ := regexp.MatchString("[A-Z]", remark)
	isShouting := hasUpperCase && !hasLowerCase
	if isShouting && isQuestion {
		return "Calm down, I know what I'm doing!"
	} else if isShouting {
		return "Whoa, chill out!"
	} else if isQuestion {
		return "Sure."
	} else {
		return "Whatever."
	}
}
