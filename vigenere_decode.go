package main

import (
	"fmt"
	"os"
)

// lets use a RingBuffer since this is a shiftcipher
const MAXLEN int = 26

type RingBuffer struct {
	lower [MAXLEN]rune
	upper [MAXLEN]rune
}

func main() {
	// arg vars
	//var input string
	var key string

	// initialize shift buffer
	var rb RingBuffer
	rb = popbuff(rb)

	if len(os.Args) == 3 {
		// raw dogging it! taking stdin directly from args with no bounds checking
		// hack me bro!
		//input = os.Args[1]
		key = os.Args[2]
	} else {
		// default values
		//input = "Please enter 1 a string to shift and 2 a key to apply"
		key = "a"
	}

	shiftarr := get_shiftmap(key)
	for _, x := range shiftarr {
		fmt.Printf("%d\n", x)
	}
}

// populate shift buffer
func popbuff(rb RingBuffer) RingBuffer {
	for i := 0; i < MAXLEN; i++ {
		rb.lower[i] = rune('a' + i)
	}

	for i := 0; i < MAXLEN; i++ {
		rb.upper[i] = rune('A' + i)
	}

	return RingBuffer{
		lower: rb.lower,
		upper: rb.upper,
	}
}

func get_shiftmap(key string) []int {
	// create a slice containing alphabetical diff from a
	var shiftmap []int
	for _, char := range key {
		shiftmap = append(shiftmap, int(char-checkcase(char)))
	}
	return shiftmap
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
