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

func TestPart1Input1(t *testing.T) {
	result := part1(getInput("inputs/2024/15/test1.txt"))
	expected := 2028
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart1Input2(t *testing.T) {
	result := part1(getInput("inputs/2024/15/test2.txt"))
	expected := 10092
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart1Input(t *testing.T) {
	result := part1(getInput("inputs/2024/15/input.txt"))
	expected := 1509863
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2Input2(t *testing.T) {
	result := part2(getInput("inputs/2024/15/test2.txt"))
	expected := 9021
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2Input(t *testing.T) {
	result := part2(getInput("inputs/2024/15/input.txt"))
	expected := 1548815
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
