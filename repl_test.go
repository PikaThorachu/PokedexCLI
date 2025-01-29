package main

import (
	"testing"
)

func TestClearnInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "  Charmander  Bulbasaur  PIKACHU  ",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "When he encountered maize for the first time, he thought it incredibly corny.",
			expected: []string{"when", "he", "encountered", "maize", "for", "the", "first", "time,", "he", "thought", "it", "incredibly", "corny."},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("cleanInput(%s) == %v, expected %v", c.input, actual, c.expected)
		}
		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("cleanInput(%s) == %v, expected %v", c.input, actual, c.expected)
			}
		}
	}
}
