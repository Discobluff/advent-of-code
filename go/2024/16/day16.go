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

var Scores [][]int

func isValid(grid []string, pos Position) bool {
	return grid[pos.Line][pos.Column] != '#'
}

func best(scores map[StepPath]int, pos1 StepPath, pos2 StepPath, score int) int {
	var _, ok = scores[pos1]
	if !ok {
		return scores[pos2] + score
	}
	if scores[pos1] < scores[pos2]+score {
		return scores[pos1]
	}
	return scores[pos2] + score
}

func insert(tab []StepPath, scores map[StepPath]int, step StepPath) []StepPath {
	var index int = -1
	for i, pos := range tab {
		if scores[pos] > scores[step] {
			index = i
			break
		}
	}
	if index != -1 {
		var res []StepPath
		res = append(res, tab[:index]...)
		res = append(res, step)
		res = append(res, tab[index:]...)
		return res
	}
	return append(tab, step)

}

func solve(grid []string, start Position) map[StepPath]int {
	var scores map[StepPath]int = make(map[StepPath]int)
	scores[StepPath{pos: start, direction: E}] = 0
	var nexts []StepPath = make([]StepPath, 1)
	nexts[0] = StepPath{pos: start, direction: E}
	var visited map[StepPath]struct{} = make(map[StepPath]struct{})
	for len(nexts) > 0 {
		var position = nexts[0]
		nexts = nexts[1:]
		var _, ok = visited[position]
		if !ok {
			visited[position] = struct{}{}
			for _, direction := range DirectionsSlice {
				if direction != OpposedDirection(position.direction) {
					if direction == position.direction {
						var newPos = AddPositions(position.pos, direction)
						if isValid(grid, newPos) {
							scores[StepPath{pos: newPos, direction: direction}] = best(scores, StepPath{pos: newPos, direction: direction}, position, 1)
							nexts = insert(nexts, scores, StepPath{pos: newPos, direction: direction})
						}
					} else {
						var newPos = StepPath{pos: position.pos, direction: direction}
						scores[newPos] = best(scores, newPos, position, 1000)
						nexts = insert(nexts, scores, newPos)
					}
				}
			}
		}
	}
	return scores
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
	fmt.Println("--2024 day 16 solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(input))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("part2 : ", part2(input))
	fmt.Println(time.Since(start))
}
