package main

import (
	_ "embed"
	"testing"
)

//go:embed input.txt
var inputDay string

func TestPart1Input(t *testing.T) {
	result := part1(inputDay)
	expected := 224529
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

// func TestPart2Input(t *testing.T) {
// 	result := part2(inputDay)
// 	expected := 0
// 	if result != expected {
// 		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
// 	}
// }

func TestCountDigits(t *testing.T) {
	var start = []int{0, 1, 9, 10, 50, 99, 100, 500, 10001, 1064842454}
	var expected = []int{1, 1, 1, 2, 2, 2, 3, 3, 5, 10}
	for i, val := range start {
		var res, ok = countDigits(val)
		if res != expected[i] || ok != (expected[i]%2 == 0) {
			t.Errorf("Result is incorrect, got: %d,%t, want: %d,%t.", res, ok, expected[i], expected[i]%2 == 0)
		}
	}
}

func TestSplit(t *testing.T) {
	var start = []int{10, 512072, 9664}
	var expected1 = []int{1, 512, 96}
	var expected2 = []int{0, 72, 64}
	for i, val := range start {
		var digits, _ = countDigits(val)
		var n1, n2 = split(val, digits)
		if n1 != expected1[i] || n2 != expected2[i] {
			t.Errorf("Result is incorrect, got: %d,%d, want: %d,%d.", n1, n2, expected1[i], expected2[i])
		}
	}

}
