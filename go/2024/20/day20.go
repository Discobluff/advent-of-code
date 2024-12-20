package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"

	. "github.com/Discobluff/advent-of-code/go/utils/positions"
)

//go:embed input.txt
var input string

func isValidSize(grid []string, pos Position) bool {
	return pos.Line >= 0 && pos.Column >= 0 && pos.Line < len(grid) && pos.Column < len(grid[0])
}

func isValidFair(grid []string, pos Position) bool {
	return isValidSize(grid, pos) && Eval(grid, pos) != '#'
}

func getTimeFair(grid []string, start Position, lastDirection Position, end Position) int {
	if start == end {
		return 1
	}
	for _, direction := range DirectionsSlice {
		if direction != OpposedDirection(lastDirection) {
			var newPos = AddPositions(start, direction)
			if isValidFair(grid, newPos) {
				return 1 + getTimeFair(grid, newPos, direction, end)
			}
		}
	}
	return 0
}

func getPositionsFair(grid []string, start Position, lastDirection Position, end Position) []Position {
	var length int = getTimeFair(grid, start, lastDirection, end)
	var res []Position = make([]Position, length)
	for i := range length - 1 {
		for _, direction := range DirectionsSlice {
			if direction != OpposedDirection(lastDirection) {
				var newPos = AddPositions(start, direction)
				if isValidFair(grid, newPos) {
					res[i] = start
					start = newPos
					lastDirection = direction
					break
				}
			}
		}
	}
	res[length-1] = end
	return res
}

func part1(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var start Position = SearchStartLines(lines, 'S')
	var end Position = SearchStartLines(lines, 'E')
	var positions []Position = getPositionsFair(lines, start, S, end)
	var res int
	for i, p1 := range positions {
		for j := i + 1; j < len(positions); j++ {
			var p2 Position = positions[j]
			if Distance(p1, p2) == 2 && (p1.Line == p2.Line || p1.Column == p2.Column) {
				if j-i-2 >= 100 {
					res++
				}
			}
		}
	}
	return res
}

func getFirstDirection(grid []string, start Position) Position {
	for _, direction := range DirectionsSlice {
		if Eval(grid, AddPositions(start, direction)) == '.' {
			return direction
		}
	}
	return DefPosition(-1, -1)
}

func part2(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var start Position = SearchStartLines(lines, 'S')
	var end Position = SearchStartLines(lines, 'E')
	var direction Position = getFirstDirection(lines, start)
	var positions []Position = getPositionsFair(lines, start, direction, end)
	var res int
	for i, p1 := range positions {
		for j := i + 2; j < len(positions); j++ {
			var p2 Position = positions[j]
			if Distance(p1, p2) <= 20 {
				if j-i-Distance(p1, p2) >= 100 {
					res++
				}
			}
		}
	}
	return res
}

func main() {
	fmt.Println("--2024 day 20 solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(input))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("part2 : ", part2(input))
	fmt.Println(time.Since(start))
}
