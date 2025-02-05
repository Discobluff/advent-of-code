package main

import (
	_ "embed"
	"os"
	"slices"
	"testing"
)

func getInput(path string) string {
	data, _ := os.ReadFile(path)
	return string(data)
}

func TestPart1Input(t *testing.T) {
	result := part1(getInput("inputs/2024/08/input.txt"))
	expected := 305
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2Input(t *testing.T) {
	result := part2(getInput("inputs/2024/08/input.txt"))
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
	var pos [][2]Position = [][2]Position{{initPosition(3, 4), initPosition(5, 5)}, {initPosition(2, 1), initPosition(9, 15)}, {initPosition(0, 0), initPosition(1, 3)}}
	var distances []int = []int{2, 0, 3}
	var expected [][2]Position = [][2]Position{{initPosition(1, 3), initPosition(7, 6)}, {initPosition(9, 15), initPosition(2, 1)}, {initPosition(-2, -6), initPosition(3, 9)}}
	for i, positions := range pos {
		var res1, res2 = antinode(positions[0], positions[1], distances[i])
		if res1.line != expected[i][0].line || res1.column != expected[i][0].column || res2.line != expected[i][1].line || res2.column != expected[i][1].column {
			t.Errorf("Result is incorrect, got: %d,%d|%d,%d, want: %d,%d|%d,%d.", res1.line, res1.column, res2.line, res2.column, expected[i][0].line, expected[i][0].column, expected[i][1].line, expected[i][1].column)
		}
	}
}

func TestInitMap(t *testing.T) {
	var dict = initMap([]string{"..4.", ".4.a", "a.A.", "...A"})
	if !slices.Equal(dict['4'], []Position{initPosition(0, 2), initPosition(1, 1)}) {
		t.Errorf("Result is incorrect for char: %c, got: %v, want: %v", '4', dict['4'], []Position{initPosition(0, 2), initPosition(1, 1)})
	}
	if !slices.Equal(dict['a'], []Position{initPosition(1, 3), initPosition(2, 0)}) {
		t.Errorf("Result is incorrect for char: %c, got: %v, want: %v", '4', dict['a'], []Position{initPosition(0, 2), initPosition(1, 1)})
	}
	if !slices.Equal(dict['A'], []Position{initPosition(2, 2), initPosition(3, 3)}) {
		t.Errorf("Result is incorrect for char: %c, got: %v, want: %v", '4', dict['A'], []Position{initPosition(0, 2), initPosition(1, 1)})
	}
}

func TestInitGrid(t *testing.T) {
	var grid = initGrid(2, 3)
	if len(grid) != 3 || len(grid[0]) != 2 {
		t.Errorf("Result is incorrect for size, got: %d,%d, want: %d,%d", len(grid), len(grid[0]), 3, 2)
	}
	for i := range 3 {
		for j := range 2 {
			if grid[i][j] {
				t.Errorf("Result is incorrect for grid at %d,%d", i, j)
			}
		}
	}
}
