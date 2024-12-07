package main

import (
	_ "embed"
	"slices"
	"testing"
)

//go:embed input.txt
var inputDay string

func TestPart1Input(t *testing.T) {
	result := part1(inputDay)
	expected := 2314935962622
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2Input(t *testing.T) {
	result := part2(inputDay)
	expected := 401477450831495
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestCountDigits(t *testing.T) {
	var start = [10]int{0, 1, 9, 50, 99, 100, 1000, 60542, 9999999999, 64813278512468}
	var resultsExpected = [10]int{1, 1, 1, 2, 2, 3, 4, 5, 10, 14}
	for i, number := range start {
		var result = countDigits(number)
		if result != resultsExpected[i] {
			t.Errorf("Result is incorrect for %d, got: %d, want: %d.", start[i], result, resultsExpected[i])
		}
	}
}

func TestConcatenate(t *testing.T) {
	var start = [8][2]int{{450, 680}, {1, 0}, {50, 6}, {6, 48952}, {231659, 0}, {5648321, 32189853}, {6597456, 32}, {0, 321}}
	var resultsExpected = [8]int{450680, 10, 506, 648952, 2316590, 564832132189853, 659745632, 321}
	for i, numbers := range start {
		var result = concatenate(numbers[0], numbers[1])
		if result != resultsExpected[i] {
			t.Errorf("Result is incorrect for %d || %d, got: %d, want: %d.", start[i][0], start[i][1], result, resultsExpected[i])
		}
	}
}

func TestParse(t *testing.T) {
	var start = [4]string{"6153340711: 5 6 395 3 3 676 2 5 6 2", "192: 17 8 14", "33: 1", "21037: 9 7 18 13"}
	var numbersExpected = [4]int{6153340711, 192, 33, 21037}
	var slicesExepected = [4][]int{{5, 6, 395, 3, 3, 676, 2, 5, 6, 2}, {17, 8, 14}, {1}, {9, 7, 18, 13}}
	for i, str := range start {
		var result1, result2 = parse(str)
		if result1 != numbersExpected[i] {
			t.Errorf("Result is incorrect for the parsing of %s, got: %d, want: %d.", start[i], result1, numbersExpected[i])
		}
		if !slices.Equal(result2, slicesExepected[i]) {
			t.Errorf("Result is incorrect for the parsing of %s, got: %v, want: %v.", start[i], result2, slicesExepected[i])
		}
	}
}
