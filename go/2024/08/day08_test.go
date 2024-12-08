package main

import (
	_ "embed"
	"testing"
)

// //go:embed input.txt
var inputDay string

func TestPart1Input(t *testing.T) {
	result := part1(inputDay)
	expected := 305
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2Input(t *testing.T) {
	result := part2(inputDay)
	expected := 1150
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestInitPosition(t *testing.T) {
	var pos = initPosition(8, 9)
	if pos.line != 8 || pos.column != 9 {
		t.Errorf("Result is incorrect, got: %d,%d, want: %d,%d.", pos.line, pos.column, 8, 9)
	}
}

func TestIsValide(t *testing.T) {
	if isValid(initPosition(-1, 5), 4, 6) {
		t.Errorf("Result is incorrect, got: %t, want: %t.", true, false)
	}
	if isValid(initPosition(0, 6), 4, 6) {
		t.Errorf("Result is incorrect, got: %t, want: %t.", true, false)
	}
	if !isValid(initPosition(1, 5), 4, 6) {
		t.Errorf("Result is incorrect, got: %t, want: %t.", false, true)
	}
}

func TestAntinode(t *testing.T) {
	var pos1, pos2 = antinode1(initPosition(3, 4), initPosition(5, 5))
	if pos1.line != 1 || pos1.column != 3 || pos2.line != 7 || pos2.column != 6 {
		t.Errorf("Result is incorrect, got: %d,%d|%d,%d, want: %d,%d|%d,%d.", pos1.line, pos1.column, pos2.line, pos2.column, 1, 3, 7, 6)
	}
}
