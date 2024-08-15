package main

import (
	"fmt"
	"os"
)

const MAXLEN int = 26

type RingBuffer struct {
	lower [MAXLEN]rune
	upper [MAXLEN]rune
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

	rb = populatebuff(rb)
	shiftmap := get_shiftmap(key, mode, rb)

	fmt.Println(string(apply_shift(input, rb, shiftmap)))
}

func checkcase(input rune, rb RingBuffer) [MAXLEN]rune {
	if input >= 'a' && input <= 'z' {
		// lowercase
		return rb.lower
	} else if input >= 'A' && input <= 'Z' {
		// uppercase
		return rb.upper
	} else {
		// not a char
		// populate with zeros to signal to upstream that
		var notchar [MAXLEN]rune
		for i := range MAXLEN {
			notchar[i] = 0
		}
		return notchar
	}
}

// TODO: I hate passing in rb just because checkcase needs it
// There needs to be a better way to approach this.
func get_shiftmap(key string, mode string, rb RingBuffer) []rune {
	// create a slice containing alphabetical diff from a
	var shiftmap []rune
	for _, char := range key {
		base := checkcase(char, rb)
		if base[0] == 0 {
			// not a char
			// skip iteration
			continue
		}
		shift := char - base[0] // should be 'a' or 'A'
		if mode == "decode" {
			// add a negative number for decoding
			shift = shift * -1
		}
		shiftmap = append(shiftmap, shift)
	}
	return shiftmap
}

// populate shift buffer
func populatebuff(rb RingBuffer) RingBuffer {
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

func apply_shift(input string, rb RingBuffer, shiftmap []rune) []rune {
	var output []rune
	var ulcase [MAXLEN]rune

	// need to create local varible to track index of shiftmaps
	// external of loop iterator for string
	// this is because we want to increment the string by len
	// but "Skip" special chars
	var ind int
	for _, c := range input {
		currkey := shiftmap[ind%len(shiftmap)]
		ind++
		ulcase = checkcase(c, rb)
		if ulcase[0] == 0 {
			// Not a letter
			output = append(output, c)
			ind--
			continue
		}
		outind := int(c+currkey-ulcase[0]) % MAXLEN
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
