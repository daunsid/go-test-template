package main

import "testing"

func TestCalculate(t *testing.T) {
	if Calculate(20) != 22 {
		t.Error("Expected 20 + 2 to equal 22")
	}
}

func TestTableCalculate(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
		{0, 2},
	}
	for _, test := range tests {
		if output := Calculate(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, output, test.expected)
		}
	}
}
