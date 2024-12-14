package main

import (
	_ "embed"
	"testing"
)

//go:embed input.txt
var inputDay string

//go:embed test.txt
var input1 string

func TestPart1Input(t *testing.T) {
	result := part1(inputDay)
	expected := 233709840
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart1Input1(t *testing.T) {
	result := part1(input1)
	expected := 12
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2Input(t *testing.T) {
	result := part2(inputDay)
	expected := 6620
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestParse(t *testing.T) {
	result := parse("p=3526,-42 v=7,65310")
	var posExpected = Position{x: 3526, y: -42}
	var velExpected = Position{x: 7, y: 65310}
	var expected = Robot{position: posExpected, velocity: velExpected}
	if result != expected {
		t.Errorf("Result is incorrect, got %v, want:%v", result, expected)
	}
}

func TestParseSize(t *testing.T) {
	r1, r2 := parseSize("9965,412302")
	if r1 != 9965 || r2 != 412302 {
		t.Errorf("Result is incorrect, got %d,%d, want:%d,%d", r1, r2, 9965, 412302)
	}
}
