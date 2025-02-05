package main

import (
	_ "embed"
	"maps"
	"os"
	"testing"
)

func getInput(path string) string {
	data, _ := os.ReadFile(path)
	return string(data)
}

func TestPart1Input(t *testing.T) {
	result := part1(getInput("inputs/2024/11/input.txt"))
	expected := 224529
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2Input(t *testing.T) {
	result := part2(getInput("inputs/2024/11/input.txt"))
	expected := 266820198587914
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

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

func TestParse(t *testing.T) {
	var result = parse("48 56 123 560023 1")
	var expected = map[int]int{48: 1, 56: 1, 123: 1, 560023: 1, 1: 1}
	if !maps.Equal(result, expected) {
		t.Errorf("Result is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestSumMap(t *testing.T) {
	var result = sumMap(map[int]int{43: 2, 0: 98, 156: 29, 33: 1})
	var expected = 130
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestBlinck(t *testing.T) {
	var s11, _, ok1 = blink(0)
	if ok1 || s11 != 1 {
		t.Errorf("Result is incorrect, got: %d,%t, want: %d,%t.", s11, ok1, 1, false)
	}
	var s21, _, ok2 = blink(1)
	if ok2 || s21 != 2024 {
		t.Errorf("Result is incorrect, got: %d,%t, want: %d,%t.", s21, ok2, 2024, false)
	}
	var s31, s32, ok3 = blink(5568)
	if !ok3 || s31 != 55 || s32 != 68 {
		t.Errorf("Result is incorrect, got: %d,%d,%t, want: %d,%d,%t.", s31, s32, ok3, 55, 68, true)
	}
	var s41, s42, ok4 = blink(6948700036)
	if !ok4 || s41 != 69487 || s42 != 36 {
		t.Errorf("Result is incorrect, got: %d,%d,%t, want: %d,%d,%t.", s41, s42, ok4, 69487, 36, true)
	}
	var s51, _, ok5 = blink(8645321)
	if ok5 || s51 != 2024*8645321 {
		t.Errorf("Result is incorrect, got: %d,%t, want: %d,%t.", s51, ok5, 2024*8645321, false)
	}
}
