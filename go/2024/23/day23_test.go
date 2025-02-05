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
	result := part1(getInput("inputs/2024/23/input.txt"))
	expected := 1352
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2Input(t *testing.T) {
	result := part2(getInput("inputs/2024/23/input.txt"))
	expected := "dm,do,fr,gf,gh,gy,iq,jb,kt,on,rg,xf,ze"
	if result != expected {
		t.Errorf("Result is incorrect, got: %s, want: %s.", result, expected)
	}
}
