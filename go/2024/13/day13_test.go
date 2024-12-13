package main

import (
	_ "embed"
	"testing"
)

//go:embed input.txt
var inputDay string

//go:embed test.txt
var input1 string

func TestPart1Input(t *testing.T) {
	result := part1(inputDay)
	expected := 33481
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart1Input1(t *testing.T) {
	result := part1(input1)
	expected := 480
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2Input(t *testing.T) {
	result := part2(inputDay)
	expected := 92572057880885
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
