package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Gotta catch them all",
			expected: []string{"gotta", "catch", "them", "all"},
		},
	}

	for _, c := range cases {
		testname := fmt.Sprintf("Testing input of: %v", c.input)

		t.Run(testname, func(t *testing.T) {
			actual := cleanInput(c.input)

			if len(actual) != len(c.expected) {
				t.Errorf("actual slice not the same length as expected slice")
			}

			for i := range actual {
				word := actual[i]
				expectedWord := c.expected[i]
				fmt.Printf("Word %d is %v \n", i, word)
				fmt.Printf("Expected %d is %v \n", i, expectedWord)

				if word != expectedWord {
					t.Errorf("test failed")
				}
			}
		})

	}
}
