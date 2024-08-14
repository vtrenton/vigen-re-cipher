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
	var input string
	var key string

	// initialize shift buffer
	var rb RingBuffer
	rb = popbuff(rb)

	if len(os.Args) == 3 {
		// raw dogging it! taking stdin directly from args with no bounds checking
		// hack me bro!
		input = os.Args[1]
		key = os.Args[2]
	} else {
		// default values
		input = "Please enter 1 a string to shift and 2 a key to apply"
		key = "a"
	}

	fmt.Println("This is where the output will go")
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

func get_keymap(key string) []int {
	// get len of key in runes
	// create a slice containing alphabetical diff from a
	// ascii can help here
	// return a slice of integers.
}
