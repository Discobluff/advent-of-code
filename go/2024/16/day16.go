package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"

	. "github.com/Discobluff/advent-of-code/go/utils/positions"
)

//go:embed test1.txt
var input string

type StepPath struct {
	pos, direction Position
}

var Scores [][]int

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

func best(scores [][]int, pos1 Position, pos2 Position, score int) int {
	if scores[pos1.Line][pos1.Column] == -1 {
		return scores[pos2.Line][pos2.Column] + score
	}
	if scores[pos1.Line][pos1.Column] < scores[pos2.Line][pos2.Column]+score {
		return scores[pos1.Line][pos1.Column]
	}
	return scores[pos2.Line][pos2.Column] + score
}

func display(scores [][]int) {
	for _, line := range scores {
		fmt.Println(line)
	}
}

func insert(tab []StepPath, scores [][]int, step StepPath) []StepPath {
	var index int = -1
	for i, pos := range tab {
		if scores[pos.pos.Line][pos.pos.Column] > scores[step.pos.Line][step.pos.Column] {
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

func solve(grid []string, start Position) [][]int {
	var scores [][]int = make([][]int, len(grid))
	for i, line := range grid {
		scores[i] = make([]int, len(line))
		for j := range line {
			scores[i][j] = -1
		}
	}
	scores[start.Line][start.Column] = 0
	var nexts []StepPath = make([]StepPath, 1)
	nexts[0] = StepPath{pos: start, direction: E}
	var visited map[Position]struct{} = make(map[Position]struct{})
	for len(nexts) > 0 {
		var position = nexts[0]
		nexts = nexts[1:]
		var _, ok = visited[position.pos]
		if !ok {
			visited[position.pos] = struct{}{}
			for _, direction := range DirectionsSlice {
				if direction != OpposedDirection(position.direction) {
					var newPos = AddPositions(position.pos, direction)
					if isValid(grid, newPos) {
						var newScore int
						if direction == position.direction {
							newScore = best(scores, newPos, position.pos, 1)
						} else {
							newScore = best(scores, newPos, position.pos, 1001)
						}
						scores[newPos.Line][newPos.Column] = newScore
						nexts = insert(nexts, scores, StepPath{pos: newPos, direction: direction})
					}
				}
			}
		}
	}
	return scores
}

func next(dict map[Position]struct{}) Position {
	for pos := range dict {
		return pos
	}
	return Position{Line: -1, Column: -1}
}
func displayV(height int, length int, dict map[Position]struct{}) {
	var tab [][]int = make([][]int, height)
	for i := range height {
		tab[i] = make([]int, length)
	}
	for pos := range dict {
		tab[pos.Line][pos.Column] = 1
	}
	for _, line := range tab {
		fmt.Println(line)
	}
}

func getScore(scores [][]int, pos Position) int {
	return scores[pos.Line][pos.Column]
}

func explore2(scores [][]int, end Position, grid []string) int {
	var visited map[Position]struct{} = make(map[Position]struct{})
	var nexts map[Position]struct{} = make(map[Position]struct{})
	nexts[end] = struct{}{}
	for len(nexts) > 0 {
		var pos = next(nexts)
		delete(nexts, pos)
		var _, ok = visited[pos]
		if !ok {
			visited[pos] = struct{}{}
			var score = getScore(scores, pos)
			for _, direction := range DirectionsSlice {
				var newPos = AddPositions(pos, direction)
				var scoreNew = getScore(scores, newPos)
				if direction == S {

					var ter = AddPositions(newPos, W)
					if scoreNew != -1 && (scoreNew == score-1 || scoreNew == score-1001 || (isValid(grid, ter) && getScore(scores, ter) == score-2)) {
						nexts[newPos] = struct{}{}
						fmt.Println("ouba", (isValid(grid, ter) && getScore(scores, ter) == score-2))
					}
				} else {
					if scoreNew != -1 && (scoreNew == score-1 || scoreNew == score-1001) {
						nexts[newPos] = struct{}{}
					}
				}
			}
		}
	}
	fmt.Println(visited)
	displayV(15, 15, visited)
	return len(visited)
}

func part1(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var start Position = SearchStartLines(lines, 'S')
	var end Position = SearchStartLines(lines, 'E')
	var scores [][]int = solve(lines, start)
	return scores[end.Line][end.Column]
}

func part2(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var start Position = SearchStartLines(lines, 'S')
	var end Position = SearchStartLines(lines, 'E')
	var scores [][]int = solve(lines, start)
	display(scores)
	// displayV()
	return explore2(scores, end, lines)
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
