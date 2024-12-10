package main

import (
	_ "embed"
	"testing"
)

// //go:embed input.txt
var inputDay string

func TestPart1Input(t *testing.T) {
	result := part1(inputDay)
	expected := 776
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2Input(t *testing.T) {
	result := part2(inputDay)
	expected := 1657
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
