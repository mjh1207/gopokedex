package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	} {
		{
			input: " hello world ",
			expected: []string {"hello", "world"},
		},
		{
			input: "Hello my name is Mitch",
			expected: []string{"hello", "my", "name", "is", "mitch"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Length of actual cleanInput: %v does not match expected length: %v", len(actual), len(c.expected))
			t.Fail()
		}
		
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("actual: %s does not match expected: %s", word, expectedWord)
				t.Fail()
			}
		}
	}
}
