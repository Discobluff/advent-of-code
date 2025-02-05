package main

import (
	_ "embed"
	"os"
	"testing"
)

func getInput(path string) string {
	data, _ := os.ReadFile(path)
	return string(data)
}

func TestPart1InputTest(t *testing.T) {
	result := part1(getInput("../../../inputs/2024/18/test.txt"))
	expected := 22
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart1Input(t *testing.T) {
	result := part1(getInput("../../../inputs/2024/18/input.txt"))
	expected := 324
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2InputTest(t *testing.T) {
	result := part2(getInput("../../../inputs/2024/18/test.txt"))
	expected := "6,1"
	if result != expected {
		t.Errorf("Result is incorrect, got: %s, want: %s.", result, expected)
	}
}

func TestPart2Input(t *testing.T) {
	result := part2(getInput("../../../inputs/2024/18/input.txt"))
	expected := "46,23"
	if result != expected {
		t.Errorf("Result is incorrect, got: %s, want: %s.", result, expected)
	}
}
