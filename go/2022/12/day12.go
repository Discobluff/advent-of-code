package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
	"time"

	. "github.com/Discobluff/advent-of-code/go/utils/positions"
	. "github.com/Discobluff/advent-of-code/go/utils/search"
	// . "github.com/Discobluff/advent-of-code/go/utils/set"
)

func isValid(grid [][]byte, pos Position) bool {
	return pos.Line >= 0 && pos.Column >= 0 && pos.Line < len(grid) && pos.Column < len(grid[0])
}

func funcNeighbors(grid [][]byte) func(Position) map[Position]int {
	return func(pos Position) map[Position]int {
		var res map[Position]int = make(map[Position]int)
		for _, dir := range Directions {
			var newPos = AddPositions(dir, pos)
			if isValid(grid, newPos) {
				var evalPos = grid[pos.Line][pos.Column]
				var evalNewPos = grid[newPos.Line][newPos.Column]
				if evalPos+1 == evalNewPos || evalPos >= evalNewPos {
					res[newPos] = 1
				}
			}
		}
		return res
	}
}

func parse(lines []string) [][]byte {
	var res [][]byte = make([][]byte, len(lines))
	for i, line := range lines {
		res[i] = make([]byte, len(line))
		for j, char := range line {
			res[i][j] = byte(char)
		}
	}
	return res
}

func part1(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var grid = parse(lines)
	var start = SearchStartGrid(grid, 'S')
	var end = SearchStartGrid(grid, 'E')
	fmt.Println(start, end)
	grid[start.Line][start.Column] = 'a'
	grid[end.Line][end.Column] = 'z'
	return Dijkstra(start, funcNeighbors(grid))[end]
}

func part2(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var grid = parse(lines)
	var start = SearchStartGrid(grid, 'S')
	var end = SearchStartGrid(grid, 'E')
	grid[start.Line][start.Column] = 'a'
	grid[end.Line][end.Column] = 'z'
	var min = -1
	for l, line := range grid {
		for c, char := range line {
			if char == 'a' {
				var calc = Dijkstra(Position{Line: l, Column: c}, funcNeighbors(grid))[end]
				if calc != 0 && (min == -1 || calc < min) {
					min = calc
				}
			}
		}
	}
	return min
}

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("--2024 day 05 solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(string(input)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("part2 : ", part2(string(input)))
	fmt.Println(time.Since(start))
}
