package main

import (
	_ "embed"
	"fmt"
	"slices"
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

func getTimeFair(grid []string, start Position, lastDirection Position, end Position) []Position {
	if start == end {
		return []Position{end}
	}
	for _, direction := range DirectionsSlice {
		if direction != OpposedDirection(lastDirection) {
			var newPos = AddPositions(start, direction)
			if isValidFair(grid, newPos) {
				return append(getTimeFair(grid, newPos, direction, end), start)
			}
		}
	}
	return []Position{}
}
func part1(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var start Position = SearchStartLines(lines, 'S')
	var end Position = SearchStartLines(lines, 'E')
	var positions []Position = getTimeFair(lines, start, S, end)
	slices.Reverse(positions)
	var res int
	for i, p1 := range positions {
		for j, p2 := range positions {
			if j > i && Abs(p1.Line-p2.Line) == 2 && p1.Column == p2.Column {
				// if Eval(lines, DefPosition(min(p1.Line, p2.Line)+1, p1.Column)) == '#' {
				if j-i-2 >= 100 {
					res++
				}
				// }
			}
			if j > i && Abs(p1.Column-p2.Column) == 2 && p1.Line == p2.Line {
				// if Eval(lines, DefPosition(p1.Line, min(p1.Column, p2.Column)+1)) == '#' {
				if j-i-2 >= 100 {
					res++
				}
				// }
			}
		}
	}
	return res
}

func funcNeighbors(grid []string, end Position) func(Position) map[Position]int {
	return func(pos Position) map[Position]int {
		var res map[Position]int = make(map[Position]int)
		for _, direction := range DirectionsSlice {
			var newPos Position = AddPositions(pos, direction)
			if newPos == end || (newPos.Column >= min(pos.Column, end.Column) && newPos.Column <= max(pos.Column, end.Column) && newPos.Line >= min(pos.Line, end.Line) && newPos.Line <= max(pos.Line, end.Line) && Eval(grid, newPos) == '#') {
				res[newPos] = 1
			}
		}
		return res
	}
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
	var positions []Position = getTimeFair(lines, start, direction, end)
	slices.Reverse(positions)
	var res int
	for i, p1 := range positions {
		for j := i + 2; j < len(positions); j++ {
			var p2 Position = positions[j]
			if j > i && Distance(p1, p2) <= 20 {
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
