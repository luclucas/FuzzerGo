package main

import ("fmt"
"math/rand/v2"
)

func main() {
	fmt.Print(fuzzer(100, 32, 32))
}

func fuzzer(maxLength int, charStart int, charRange int) string {
	var stringLength = rand.IntN(maxLength)
	var out = ""

	for i := 0; i < stringLength; i++ {
		out += string(rune(rand.IntN(charRange) + charStart))
	}
	return out
}