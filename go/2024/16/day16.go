package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
	"time"

	. "github.com/Discobluff/advent-of-code/go/utils/positions"
	. "github.com/Discobluff/advent-of-code/go/utils/search"
)

type StepPath struct {
	pos, direction Position
}

func funcNeighbors(grid []string) func(StepPath) map[StepPath]int {
	return func(step StepPath) map[StepPath]int {
		var res = make(map[StepPath]int)
		for _, direction := range DirectionsSlice {
			if direction == step.direction {
				var newPos Position = AddPositions(step.pos, direction)
				if grid[newPos.Line][newPos.Column] != '#' {
					res[StepPath{pos: newPos, direction: direction}] = 1
				}
			} else {
				if direction != OpposedDirection(step.direction) {
					res[StepPath{pos: step.pos, direction: direction}] = 1000
				}
			}
		}
		return res
	}
}
func solve(grid []string, start Position) map[StepPath]int {
	return Dijkstra(StepPath{pos: start, direction: E}, funcNeighbors(grid))
}

func shortestPaths(scores map[StepPath]int, end StepPath) int {
	var visitedPos map[Position]struct{} = make(map[Position]struct{})
	var visited map[StepPath]struct{} = make(map[StepPath]struct{})
	var nexts map[StepPath]struct{} = make(map[StepPath]struct{})
	nexts[end] = struct{}{}
	for len(nexts) > 0 {
		var position StepPath = next(nexts)
		delete(nexts, position)
		visitedPos[position.pos] = struct{}{}
		var _, ok = visited[position]
		if !ok {
			visited[position] = struct{}{}
			for _, direction := range DirectionsSlice {
				if direction != position.direction {
					if direction == OpposedDirection(position.direction) {
						var newPos = AddPositions(position.pos, direction)
						if scores[StepPath{pos: newPos, direction: position.direction}] == scores[position]-1 {
							nexts[StepPath{pos: newPos, direction: position.direction}] = struct{}{}
						}
					} else {
						var newPos = StepPath{pos: position.pos, direction: direction}
						if scores[newPos] == scores[position]-1000 {
							nexts[newPos] = struct{}{}
						}
					}
				}
			}
		}
	}
	return len(visitedPos)
}

func next(dict map[StepPath]struct{}) StepPath {
	for pos := range dict {
		return pos
	}
	var posNil = Position{Line: -1, Column: -1}
	return StepPath{pos: posNil, direction: posNil}
}

func part1(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var start Position = SearchStartLines(lines, 'S')
	var end Position = SearchStartLines(lines, 'E')
	var scores map[StepPath]int = solve(lines, start)
	var res int = -1
	for _, direction := range DirectionsSlice {
		var val, ok = scores[StepPath{pos: end, direction: direction}]
		if ok && (res == -1 || val < res) {
			res = val
		}
	}
	return res
}

func part2(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var start Position = SearchStartLines(lines, 'S')
	var end Position = SearchStartLines(lines, 'E')
	var scores map[StepPath]int = solve(lines, start)
	return shortestPaths(scores, StepPath{pos: end, direction: N})
}

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("--2024 day 16 solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(string(input)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("part2 : ", part2(string(input)))
	fmt.Println(time.Since(start))
}
