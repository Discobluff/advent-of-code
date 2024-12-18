package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"

	. "github.com/Discobluff/advent-of-code/go/utils/positions"
)

//go:embed input.txt
var input string

func parse(line string) Position {
	var numbers = strings.Split(line, ",")
	var x, _ = strconv.Atoi(numbers[0])
	var y, _ = strconv.Atoi(numbers[1])
	return Position{Line: y, Column: x}
}

func isValidWalls(size int, walls map[Position]struct{}) func(Position) bool {
	return func(pos Position) bool {
		var _, ok = walls[pos]
		return pos.Line >= 0 && pos.Column >= 0 && pos.Line < size && pos.Column < size && !ok
	}
}

func initScore(size int) [][]int {
	var res [][]int = make([][]int, size)
	for i := range size {
		res[i] = make([]int, size)
		for j := range size {
			res[i][j] = -1
		}
	}
	return res
}

func part1(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var size, _ = strconv.Atoi(lines[0])
	var limit, _ = strconv.Atoi(lines[1])
	var walls map[Position]struct{} = make(map[Position]struct{})
	for _, line := range lines[2 : limit+2] {
		walls[parse(line)] = struct{}{}
	}
	var scores = initScore(size)
	Dijkstra(Position{Line: 0, Column: 0}, isValidWalls(size, walls), DefaultBest, &scores)
	return GetScore(Position{Line: size - 1, Column: size - 1}, scores)
}

func part2(input string) string {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var size, _ = strconv.Atoi(lines[0])
	var limitIndex int = 2
	var walls map[Position]struct{} = make(map[Position]struct{})
	for true {
		walls[parse(lines[limitIndex])] = struct{}{}
		var scores [][]int = initScore(size)
		Dijkstra(DefPosition(0, 0), isValidWalls(size, walls), DefaultBest, &scores)
		if GetScore(DefPosition(size-1, size-1), scores) == -1 {
			return lines[limitIndex]
		}
		limitIndex++
	}
	return ""
}

func main() {
	fmt.Println("--2024 day 18 solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(input))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("part2 : ", part2(input))
	fmt.Println(time.Since(start))
}
