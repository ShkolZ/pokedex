package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	type myCase struct {
		input    string
		expected []string
	}

	cases := []myCase{
		{input: " hello   world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  vaflya suka",
			expected: []string{"vaflya", "suka"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("slice length doesnt match")
			fmt.Printf("%v != %v", actual, c.expected)
			return
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("words do not match")

				return
			}
		}
	}
}

func TestCache(t *testing.T) {

}
