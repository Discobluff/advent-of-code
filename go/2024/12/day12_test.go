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
	result := part1(getInput("inputs/2024/12/input.txt"))
	expected := 1485656
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart1Input1(t *testing.T) {
	result := part1(getInput("inputs/2024/12/test1.txt"))
	expected := 140
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart1Input2(t *testing.T) {
	result := part1(getInput("inputs/2024/12/test2.txt"))
	expected := 1930
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart1Input3(t *testing.T) {
	result := part1(getInput("inputs/2024/12/test3.txt"))
	expected := 772
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2Input(t *testing.T) {
	result := part2(getInput("inputs/2024/12/input.txt"))
	expected := 899196
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2Input1(t *testing.T) {
	result := part2(getInput("inputs/2024/12/test1.txt"))
	expected := 80
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2Input3(t *testing.T) {
	result := part2(getInput("inputs/2024/12/test3.txt"))
	expected := 436
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
func TestPart2Input4(t *testing.T) {
	result := part2(getInput("inputs/2024/12/test4.txt"))
	expected := 368
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2Input5(t *testing.T) {
	result := part2(getInput("inputs/2024/12/test5.txt"))
	expected := 236
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
