package main

import (
	_ "embed"
	"strings"
	"testing"
)

//go:embed input.txt
var inputDay string

func TestPart1Input(t *testing.T) {
	result := part1(inputDay)
	expected := 5030
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2Input(t *testing.T) {
	result := part2(inputDay)
	expected := 1928
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestSearchStart(t *testing.T) {
	x, y := searchStart(strings.Split(strings.TrimSuffix(input, "\n"), "\n"), '^')
	expectedX, expectedY := 65, 85
	if x != expectedX || y != expectedY {
		t.Errorf("Result is incorrect, got: %d,%d, want: %d,%d.", x, y, expectedX, expectedY)
	}
}

func TestMuteString(t *testing.T) {
	result := muteString("bonjour", 2, 'k')
	expected := "bokjour"
	if result != expected {
		t.Errorf("Result is incorrect, got: %s, want: %s.", result, expected)
	}
}

func TestNextCase(t *testing.T) {
	var input1 = ".#..\n.^..\n....\n...."
	var dir1, x1, y1, ok1 = nextCase(strings.Split(strings.TrimSuffix(input1, "\n"), "\n"), 1, 1, 1)
	if dir1 != 2 || x1 != 1 || y1 != 1 || ok1 {
		t.Errorf("Result is incorrect, got: %d,%d,%d,%t, want: %d,%d,%d,%t.", dir1, x1, y1, ok1, 2, 1, 1, false)
	}
	var input2 = "....\n.^..\n....\n...."
	var dir2, x2, y2, ok2 = nextCase(strings.Split(strings.TrimSuffix(input2, "\n"), "\n"), 1, 1, 1)
	if dir2 != 1 || x2 != 0 || y2 != 1 || ok2 {
		t.Errorf("Result is incorrect, got: %d,%d,%d,%t, want: %d,%d,%d,%t.", dir2, x2, y2, ok2, 2, 1, 1, false)
	}
	var input3 = ".#.\n.^#.\n....\n...."
	var dir3, x3, y3, ok3 = nextCase(strings.Split(strings.TrimSuffix(input3, "\n"), "\n"), 2, 1, 1)
	if dir3 != 3 || x1 != 1 || y1 != 1 || ok1 {
		t.Errorf("Result is incorrect, got: %d,%d,%d,%t, want: %d,%d,%d,%t.", dir3, x3, y3, ok3, 2, 1, 1, false)
	}
	var input4 = ".^.\n...\n....\n...."
	var _, _, _, ok4 = nextCase(strings.Split(strings.TrimSuffix(input4, "\n"), "\n"), 1, 0, 1)
	if !ok4 {
		t.Errorf("Result is incorrect, got: %t, want: %t.", ok4, false)
	}
}
