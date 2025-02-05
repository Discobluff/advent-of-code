package main

import (
	_ "embed"
	"os"
	"testing"
)

//go:embed input.txt
var inputDay string

//go:embed test.txt
var inputTest string

func getInput(path string) string {
	data, _ := os.ReadFile(path)
	return string(data)
}

func TestPart1Input(t *testing.T) {
	result := part1(getInput("inputs/2022/08"))
	expected := 1695
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2Input(t *testing.T) {
	result := part2(getInput("inputs/2022/08"))
	expected := 287040
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

// func TestPart1Test(t *testing.T) {
// 	result := part1(inputTest)
// 	expected := 21
// 	if result != expected {
// 		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
// 	}
// }

// func TestPart2(t *testing.T) {
// 	result := part2(inputTest)
// 	expected := 8
// 	if result != expected {
// 		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
// 	}
// }
