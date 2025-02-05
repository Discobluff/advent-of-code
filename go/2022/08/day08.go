package main

import (
	_ "embed"
	"strconv"
	"strings"

	. "github.com/Discobluff/advent-of-code/go/utils/positions"
)

func parse(lines []string) [][]int {
	var res [][]int = make([][]int, len(lines))
	for i, line := range lines {
		res[i] = make([]int, len(line))
		for j, char := range line {
			res[i][j], _ = strconv.Atoi(string(char))
		}
	}
	return res
}

func isEdge(grid [][]int, pos Position) bool {
	return pos.Column == 0 || pos.Line == 0 || pos.Line == len(grid)-1 || pos.Column == len(grid[0])-1
}

func explore(grid [][]int, start Position, direction Position, val int) bool {
	if isEdge(grid, start) {
		return true
	}
	var newPos = AddPositions(start, direction)
	return grid[newPos.Line][newPos.Column] < val && explore(grid, newPos, direction, val)
}

func solve1(grid [][]int) int {
	var res int
	for i, line := range grid {
		for j := range line {
			var pos = Position{Line: i, Column: j}
			for _, direction := range DirectionsSlice {
				if explore(grid, pos, direction, grid[pos.Line][pos.Column]) {
					res++
					break
				}
			}
		}
	}
	return res
}

func explore2(grid [][]int, start Position, direction Position, val int) int {
	if isEdge(grid, start) {
		return 0
	}
	var newPos = AddPositions(start, direction)
	if grid[newPos.Line][newPos.Column] >= val {
		return 1
	}
	return 1 + explore2(grid, newPos, direction, val)
}

func solve2(grid [][]int) int {
	var res int = 1
	for i, line := range grid {
		for j := range line {
			var pos = Position{Line: i, Column: j}
			var temp int = 1
			for _, direction := range DirectionsSlice {
				temp *= explore2(grid, pos, direction, grid[pos.Line][pos.Column])
			}
			res = max(res, temp)
		}
	}
	return res
}

func part1(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	return solve1(parse(lines))
}

func part2(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	return solve2(parse(lines))
}
