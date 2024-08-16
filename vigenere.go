package main

import (
	"errors"
	"fmt"
	"os"
)

const MAXLEN int = 26

type RingBuffer struct {
	lower [MAXLEN]rune
	upper [MAXLEN]rune
}

// enum for mode Arg
type Mode string

const (
	Encode Mode = "encode"
	Decode Mode = "decode"
)

func main() {
	// arg vars
	var modeArg string
	var input string
	var key string

	if len(os.Args) == 4 {
		// raw dogging it! taking stdin directly from args with no bounds checking
		// hack me bro!
		modeArg = os.Args[1]
		// TODO its better to read from a file than an arg with vigenere ciphers
		// so input should be read from a file.
		input = os.Args[2]
		key = os.Args[3]
	} else {
		// default values
		modeArg = "encode"
		input = "Please enter 1 a mode 2 string to shift and 3 a key to apply"
		key = "a"
	}

	// validate mode input
	mode, err := ParseMode(modeArg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	shiftmap := get_shiftmap(key, mode)
	cipher_out := apply_shift(input, shiftmap)

	fmt.Println(cipher_out)
}

func ParseMode(mode string) (Mode, error) {
	switch mode {
	case string(Encode):
		return Encode, nil
	case string(Decode):
		return Decode, nil
	default:
		return "", errors.New("invalid mode: must be 'encode' or 'decode'")
	}
}

func checkcase(input rune) [MAXLEN]rune {
	// this is realistically the only function that needs access to the RingBuffer
	// Lets initialize it here
	// even though we'll be reinitializing on each call
	// this struct is just 2 arrays (sequential so not a slice)
	// minimal overhead
	var rb RingBuffer
	rb = populatebuff(rb)

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

func get_shiftmap(key string, mode Mode) []rune {
	// create a slice containing alphabetical diff from a
	var shiftmap []rune
	for _, char := range key {
		base := checkcase(char)
		if base[0] == 0 {
			// not a char
			// skip iteration
			continue
		}
		shift := char - base[0] // should be 'a' or 'A'
		if mode == Decode {
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

func apply_shift(input string, shiftmap []rune) string {
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
		ulcase = checkcase(c)
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
	return string(output)
}
