package main

import (
	"slices"
	"testing"
)

func TestGetArgs(t *testing.T) {
	t.Run("test empty args", func(t *testing.T) {
		args := []string{"", "", ""}

		_, _, _, err := getArgs(args)

		if err == nil {
			t.Error("expected error but didn't get one")
		}
	})
}

func TestPopulate(t *testing.T) {
	var rb RingBuffer
	rb.populatebuff()

	t.Run("test lowercase population", func(t *testing.T) {
		got_lower := rb.lower
		want_lower := [26]rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

		if got_lower != want_lower {
			t.Errorf("buffer was not populated with lower case letters, got %c want %c", got_lower, want_lower)
		}
	})

	t.Run("test uppercase population", func(t *testing.T) {
		got_upper := rb.upper
		want_upper := [26]rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

		if got_upper != want_upper {
			t.Errorf("buffer was not populated with upper case letters, got %c want %c", got_upper, want_upper)
		}
	})
}

func TestParseMode(t *testing.T) {
	t.Run("test encode mode", func(t *testing.T) {
		got, err := ParseMode("encode")
		if err != nil {
			t.Error(err)
		}
		want := Encode

		if got != want {
			t.Errorf("wanted %s, but got %s", got, want)
		}
	})

	t.Run("test decode mode", func(t *testing.T) {
		got, err := ParseMode("decode")
		if err != nil {
			t.Error(err)
		}
		want := Decode

		if got != want {
			t.Errorf("wanted %s, but got %s", got, want)
		}
	})

	t.Run("bad string passed in", func(t *testing.T) {
		_, err := ParseMode("bad")
		if err == nil {
			t.Error("expected error but didnt get one")
		}
	})

}

func TestCheckCase(t *testing.T) {
	var rb RingBuffer
	rb.populatebuff()

	t.Run("validate return of lowercase buffer", func(t *testing.T) {
		got, err := checkcase('i')
		if err != nil {
			t.Error(err)
		}
		want := rb.lower

		if got != want {
			t.Errorf("tested lowercase, got %v but wanted %v", got, want)
		}
	})

	t.Run("test return of uppercase buffer", func(t *testing.T) {
		got, err := checkcase('I')
		if err != nil {
			t.Error(err)
		}
		want := rb.upper

		if got != want {
			t.Errorf("tested uppercase, got %v but wanted %v", got, want)
		}
	})

	t.Run("test sending something other than an alphbetical char", func(t *testing.T) {
		_, err := checkcase('$')
		if err == nil {
			t.Error("expected an error but didnt get one")
		}
	})
}

func TestGetShiftmap(t *testing.T) {
	t.Run("validate encoding with a key", func(t *testing.T) {
		key := "abcdefghijklmnopqrstuvwxyz"
		mode := Encode

		got := get_shiftmap(key, mode)
		want := []rune{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}
		if !slices.Equal(got, want) {
			t.Errorf("got %c but wanted %c", got, want)
		}
	})

	t.Run("validate decoding with a key", func(t *testing.T) {
		key := "abcdefghijklmnopqrstuvwxyz"
		mode := Decode

		got := get_shiftmap(key, mode)
		want := []rune{0, -1, -2, -3, -4, -5, -6, -7, -8, -9, -10, -11, -12, -13, -14, -15, -16, -17, -18, -19, -20, -21, -22, -23, -24, -25}
		if !slices.Equal(got, want) {
			t.Errorf("got %c but wanted %c", got, want)
		}

	})
}

func TestApplyShift(t *testing.T) {
	input := "abc abc"
	key := "abc"

	t.Run("apply the encode shift based on the key", func(t *testing.T) {
		mode := Encode
		shiftmap := get_shiftmap(key, mode)

		got := apply_shift(input, shiftmap)
		want := "ace ace"

		if got != want {
			t.Errorf("got %s, but wanted %s", got, want)
		}
	})

	t.Run("apply the decode shift based on the key", func(t *testing.T) {
		mode := Decode
		shiftmap := get_shiftmap(key, mode)

		got := apply_shift(input, shiftmap)
		want := "aaa aaa"

		if got != want {
			t.Errorf("got %s, but wanted %s", got, want)
		}
	})
}
