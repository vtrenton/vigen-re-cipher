package main

import (
	"testing"
)

func TestRingBuffPopulate(t *testing.T) {
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
