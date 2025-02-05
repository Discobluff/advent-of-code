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

func TestPart1Input(t *testing.T) {
	result := part1(getInput("inputs/2024/09/input.txt"))
	expected := 6241633730082
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2Input(t *testing.T) {
	result := part2(getInput("inputs/2024/09/input.txt"))
	expected := 6265268809555
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
