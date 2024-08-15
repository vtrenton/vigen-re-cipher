package main

import (
	"fmt"
	"os"
)

// lets use a RingBuffer since this is a shiftcipher
const MAXLEN int = 26

type RingBuffer struct {
	lower    [MAXLEN]rune
	upper    [MAXLEN]rune
	shiftmap []rune
}

func main() {
	// arg vars

	// TODO: Enum for this!
	var mode string
	var input string
	var key string

	// initialize shift buffer
	var rb RingBuffer

	if len(os.Args) == 4 {
		// raw dogging it! taking stdin directly from args with no bounds checking
		// hack me bro!
		mode = os.Args[1]
		input = os.Args[2]
		key = os.Args[3]
	} else {
		// default values
		input = "Please enter 1 a mode 2  string to shift and 3 a key to apply"
		key = "a"
	}

	shiftmap := get_shiftmap(key, mode)
	rb = popbuff(rb, shiftmap)

	apply_shift(input, rb)
	//for _, x := range shiftmap {
	//	fmt.Printf("%d\n", x)
	//}
}

func checkcase(input rune) rune {
	if input >= 'a' && input <= 'z' {
		// lowercase
		return 'a'
	} else if input >= 'A' && input <= 'Z' {
		// uppercase
		return 'A'
	} else {
		// not a number
		// calling function should subtract itself so the shift is 0
		return input
	}
}

func get_shiftmap(key string, mode string) []rune {
	// create a slice containing alphabetical diff from a
	var shiftmap []rune
	for _, char := range key {
		shift := char - checkcase(char)
		if mode == "decode" {
			// add a negative number for decoding
			shift = shift * -1
		}
		shiftmap = append(shiftmap, shift)
	}
	return shiftmap
}

// populate shift buffer
func popbuff(rb RingBuffer, shiftmap []rune) RingBuffer {
	for i := 0; i < MAXLEN; i++ {
		rb.lower[i] = rune('a' + i)
	}

	for i := 0; i < MAXLEN; i++ {
		rb.upper[i] = rune('A' + i)
	}

	return RingBuffer{
		lower:    rb.lower,
		upper:    rb.upper,
		shiftmap: shiftmap,
	}
}

func apply_shift(input string, rb RingBuffer) {
	//(plaintext letter + shift letter - 'a') % MAXLEN
	for i, c := range input {
		currkey := rb.shiftmap[i%len(rb.shiftmap)]
		combo := c + currkey
		encoded := int(combo)%MAXLEN + int(checkcase(c))
		fmt.Printf("%c\n", rune(encoded))
	}
}
