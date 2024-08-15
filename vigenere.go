package main

import (
	"fmt"
	"os"
)

const MAXLEN int = 26

type RingBuffer struct {
	lower    [MAXLEN]rune
	upper    [MAXLEN]rune
	shiftmap []rune
}

func main() {
	// arg vars

	// TODO: Enum for mode!
	var mode string
	var input string
	var key string

	// initialize shift buffer
	var rb RingBuffer

	if len(os.Args) == 4 {
		// raw dogging it! taking stdin directly from args with no bounds checking
		// hack me bro!
		mode = os.Args[1]
		// TODO its better to read from a file than an arg with vigenere ciphers
		// so input should be read from a file.
		input = os.Args[2]
		key = os.Args[3]
	} else {
		// default values
		mode = "encode"
		input = "Please enter 1 a mode 2 string to shift and 3 a key to apply"
		key = "a"
	}

	shiftmap := get_shiftmap(key, mode)
	rb = populatebuff(rb, shiftmap)

	fmt.Println(string(apply_shift(input, rb)))
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
		return 0
	}
}

func get_shiftmap(key string, mode string) []rune {
	// create a slice containing alphabetical diff from a
	var shiftmap []rune
	for _, char := range key {
		base := checkcase(char)
		if base == 0 {
			// not a char
			// skip iteration
			continue
		}
		shift := char - base
		if mode == "decode" {
			// add a negative number for decoding
			shift = shift * -1
		}
		shiftmap = append(shiftmap, shift)
	}
	return shiftmap
}

// populate shift buffer
func populatebuff(rb RingBuffer, shiftmap []rune) RingBuffer {
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

func apply_shift(input string, rb RingBuffer) []rune {
	var output []rune
	var ulcase [MAXLEN]rune

	// need to create local varible to track index of shiftmaps
	// external of loop iterator for string
	// this is because we want to increment the string by len
	// but "Skip" special chars
	var ind int
	for _, c := range input {
		currkey := rb.shiftmap[ind%len(rb.shiftmap)]
		ind++
		// TODO figure out case
		// this is a mess that rechecks case.
		// need to find a more elegant DRY way to handle this
		if checkcase(c) == 'a' {
			ulcase = rb.lower
		} else if checkcase(c) == 'A' {
			ulcase = rb.upper
		} else {
			// Not a letter
			output = append(output, c)
			ind--
			continue
		}
		outind := int(c+currkey-checkcase(c)) % MAXLEN
		// in the case of decode we need to add back MAXLEN
		// So we don't have a negative index
		if outind < 0 {
			outind = outind + MAXLEN
		}
		outchar := ulcase[outind]

		output = append(output, outchar)
	}
	return output
}
