package main

import (
	_ "embed"
	"testing"
)

//go:embed input.txt
var inputDay string

//go:embed test.txt
var inputTest string

func TestPart1InputTest(t *testing.T) {
	result := part1(inputTest)
	expected := 22
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart1Input(t *testing.T) {
	result := part1(inputDay)
	expected := 324
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2InputTest(t *testing.T) {
	result := part2(inputTest)
	expected := "6,1"
	if result != expected {
		t.Errorf("Result is incorrect, got: %s, want: %s.", result, expected)
	}
}

func TestPart2Input(t *testing.T) {
	result := part2(inputDay)
	expected := "46,23"
	if result != expected {
		t.Errorf("Result is incorrect, got: %s, want: %s.", result, expected)
	}
}
