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
	result := part1(getInput("inputs/2024/14/input.txt"))
	expected := 233709840
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart1Input1(t *testing.T) {
	result := part1(getInput("inputs/2024/14/test.txt"))
	expected := 12
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2Input(t *testing.T) {
	result := part2(getInput("inputs/2024/14/input.txt"))
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
