package main

import (
	"testing"
)

func TestSumDigits(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{input: 123, expected: 6},
		{input: 999999, expected: 54},
		{input: 0, expected: 0},
		{input: -456, expected: 15},
	}
	for _, test := range tests {
		result := SumDigits(test.input)
		if result != test.expected {
			t.Errorf("SumDigits(%d) = %d; expected %d", test.input, result, test.expected)

		}

	}
}
