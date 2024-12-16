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

type StepPath struct {
	pos, direction Position
}

var ScoreMin int = -1

func isValid(grid []string, pos Position) bool {
	return grid[pos.Line][pos.Column] != '#'
}

func copy(dict map[Position]struct{}) map[Position]struct{} {
	var res = make(map[Position]struct{})
	for key := range dict {
		res[key] = struct{}{}
	}
	return res
}

func solve(grid []string, encounteredPositions map[Position]struct{}, currentPos Position, lastDirection Position, end Position, score int) {
	// fmt.Println(encounteredPositions, currentPos, score)
	if currentPos == end {
		if ScoreMin == -1 || score < ScoreMin {
			ScoreMin = score
		}
	} else {
		var _, ok = encounteredPositions[currentPos]
		if !ok && !(score > ScoreMin && ScoreMin != -1) {
			encounteredPositions[currentPos] = struct{}{}
			for _, direction := range Directions {
				if direction != OpposedDirection(lastDirection) {
					var newPos = AddPositions(currentPos, direction)
					if isValid(grid, newPos) {
						if direction == lastDirection {
							solve(grid, copy(encounteredPositions), newPos, direction, end, score+1)
						} else {
							solve(grid, copy(encounteredPositions), newPos, direction, end, score+1001)
						}

					}
				}
			}
		}
	}
}

func part1(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var start Position = SearchStartLines(lines, 'S')
	var end Position = SearchStartLines(lines, 'E')
	var encounteredPositions map[Position]struct{} = make(map[Position]struct{})
	solve(lines, encounteredPositions, start, E, end, 0)
	return ScoreMin
}

// func part2(input string) int {
// 	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
// 	var res int
// 	return res
// }

func main() {
	fmt.Println("--2024 day 16 solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(input))
	fmt.Println(time.Since(start))
	// start = time.Now()
	// fmt.Println("part2 : ", part2(input))
	// fmt.Println(time.Since(start))
}
