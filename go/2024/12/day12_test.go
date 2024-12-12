package main

import (
	_ "embed"
	"testing"
)

//go:embed input.txt
var inputDay string

//go:embed test1.txt
var input1 string

//go:embed test2.txt
var input2 string

//go:embed test3.txt
var input3 string

//go:embed test4.txt
var input4 string

//go:embed test5.txt
var input5 string

func TestPart1Input(t *testing.T) {
	result := part1(inputDay)
	expected := 1485656
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart1Input1(t *testing.T) {
	result := part1(input1)
	expected := 140
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart1Input2(t *testing.T) {
	result := part1(input2)
	expected := 1930
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart1Input3(t *testing.T) {
	result := part1(input3)
	expected := 772
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2Input(t *testing.T) {
	result := part2(inputDay)
	expected := 899196
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2Input1(t *testing.T) {
	result := part2(input1)
	expected := 80
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2Input3(t *testing.T) {
	result := part2(input3)
	expected := 436
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
func TestPart2Input4(t *testing.T) {
	result := part2(input4)
	expected := 368
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2Input5(t *testing.T) {
	result := part2(input5)
	expected := 236
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
