package main

import (
	_ "embed"
	"slices"
	"testing"
)

//go:embed input.txt
var inputDay string

func TestPart1Input(t *testing.T) {
	result := part1(inputDay)
	expected := 4790
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2Input(t *testing.T) {
	result := part2(inputDay)
	expected := 6319
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestMiddleValue(t *testing.T) {
	result := middleValue([]int{3, 4, 5})
	expected := 4
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
	result = middleValue([]int{3, 4, 5, 0, 7, 1, 9, 12, 100})
	expected = 7
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestBuildUpdates(t *testing.T) {
	result := buildUpdates("45,29,36,78")
	expected := []int{45, 29, 36, 78}
	if !slices.Equal(result, expected) {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
