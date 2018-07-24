// This package implements a secret handshake generation function. These
// handshakes consist in a list of strings and depend on the binary
// representation of an integer.
// Here are the actions allowed in a handshake, and the bits
// activating them.
// 1 = wink
// 10 = double blink
// 100 = close your eyes
// 1000 = jump
// 10000 = Reverse the order of the operations in the secret handshake.

package secret

// reverse reverses an array of strings.
func reverse(input []string) []string {
	var reversedArray = make([]string, len(input))

	for i := 0; i < len(input); i++ {
		reversedArray[len(input)-1-i] = input[i]
	}

	return reversedArray
}

// Handshake generates the string slice corresponding to the actions in
// a handshake, depending on the input integer parameter.
func Handshake(i uint) []string {
	var actions = make([]string, 0)

	if i&1 == 1 {
		actions = append(actions, "wink")
	}
	if i>>1&1 == 1 {
		actions = append(actions, "double blink")
	}
	if i>>2&1 == 1 {
		actions = append(actions, "close your eyes")
	}
	if i>>3&1 == 1 {
		actions = append(actions, "jump")
	}
	if i>>4&1 == 1 {
		actions = reverse(actions)
	}

	return actions
}
