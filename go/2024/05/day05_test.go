package main

import (
	_ "embed"
	"os"
	"slices"
	"testing"
)

func getInput(path string) string {
	data, _ := os.ReadFile(path)
	return string(data)
}

func TestPart1Input(t *testing.T) {
	result := part1(getInput("../../../inputs/2024/05/input.txt"))
	expected := 4790
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2Input(t *testing.T) {
	result := part2(getInput("../../../inputs/2024/05/input.txt"))
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
