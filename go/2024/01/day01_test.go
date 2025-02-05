package main

import (
	_ "embed"
	"maps"
	"math/rand/v2"
	"os"
	"slices"
	"testing"
)

func getInput(path string) string {
	data, _ := os.ReadFile(path)
	return string(data)
}

func TestPart1Input(t *testing.T) {
	result := part1(getInput("../../../inputs/2024/01/input.txt"))
	expected := 2815556
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2Input(t *testing.T) {
	result := part2(getInput("../../../inputs/2024/01/input.txt"))
	expected := 23927637
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestSortSlice(t *testing.T) {
	var length int = rand.IntN(60000)
	var slice []int = make([]int, length)
	for i := range length {
		slice[i] = rand.Int()
	}
	var result = sortTab(slice)
	slices.Sort(slice)
	if !slices.Equal(result, slice) {
		t.Errorf("Result is incorrect")
	}
}

func TestParse(t *testing.T) {
	var r1, r2 = parse([]string{"45   932165", "2316497   5"})
	var expected1, expected2 = []int{45, 2316497}, []int{932165, 5}
	if !slices.Equal(r1, expected1) || !slices.Equal(r2, expected2) {
		t.Errorf("Result is incorrect, got:%v,%v, want:%v,%v", r1, r2, expected1, expected2)
	}
}

func TestOccurrence(t *testing.T) {
	var result = occurrence([]int{45, 32, 6, 9, 45, 6, 6})
	var expected = map[int]int{45: 2, 32: 1, 6: 3, 9: 1}
	if !maps.Equal(result, expected) {
		t.Errorf("Result is incorrect, got:%v, want:%v", result, expected)
	}
}
