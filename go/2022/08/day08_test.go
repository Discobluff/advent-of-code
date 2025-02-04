package main

import (
	_ "embed"
	"testing"
)

//go:embed input.txt
var inputDay string

//go:embed test.txt
var inputTest string

func TestPart1Input(t *testing.T) {
	result := part1(inputDay)
	expected := 1695
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2Input(t *testing.T) {
	result := part2(inputDay)
	expected := 287040
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart1Test(t *testing.T) {
	result := part1(inputTest)
	expected := 21
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := part2(inputTest)
	expected := 8
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
