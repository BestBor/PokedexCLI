package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: 		"	hello  world ",
			expected:	[]string{"hello", "world"},
		},
		{
			input:		"this is    an example",
			expected:	[]string{"this", "is", "an", "example"},
		},
		{
			input: 		"another example 456 123 123456",
			expected:	[]string{"another", "example", "456", "123", "123456"},
		},
	}


	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(c.expected){
			t.Errorf("for %v: expected: %d, got: %d", c.input, len(c.expected), len(actual))
			continue
		}
		
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("For input: %v, expected word %s at index %d, got %s", c.input, c.expected[i], i, actual[i])
			}
		}
	}
}
