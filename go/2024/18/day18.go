package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"

	. "github.com/Discobluff/advent-of-code/go/utils/positions"
	. "github.com/Discobluff/advent-of-code/go/utils/search"
)

//go:embed input.txt
var input string

func parse(line string) Position {
	var numbers = strings.Split(line, ",")
	var x, _ = strconv.Atoi(numbers[0])
	var y, _ = strconv.Atoi(numbers[1])
	return DefPosition(y, x)
}

func isValid(size int, walls map[Position]struct{}, pos Position) bool {
	var _, ok = walls[pos]
	return pos.Line >= 0 && pos.Column >= 0 && pos.Line < size && pos.Column < size && !ok
}

func isValid2(size int, walls []Position, limit int, pos Position) bool {
	if !(pos.Line >= 0 && pos.Column >= 0 && pos.Line < size && pos.Column < size) {
		return false
	}
	for i := 0; i < limit; i++ {
		if pos == walls[i] {
			return false
		}
	}
	return true
}

func funcNeighbors(size int, walls map[Position]struct{}) func(Position) map[Position]int {
	return func(pos Position) map[Position]int {
		var res map[Position]int = make(map[Position]int)
		for _, direction := range DirectionsSlice {
			var newPos Position = AddPositions(pos, direction)
			if isValid(size, walls, newPos) {
				res[newPos] = 1
			}
		}
		return res
	}
}

func funcNeighbors2(size int, walls []Position, limit int) func(Position) map[Position]int {
	return func(pos Position) map[Position]int {
		var res map[Position]int = make(map[Position]int)
		for _, direction := range DirectionsSlice {
			var newPos Position = AddPositions(pos, direction)
			if isValid2(size, walls, limit, newPos) {
				res[newPos] = 1
			}
		}
		return res
	}
}

func part1(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var size, _ = strconv.Atoi(lines[0])
	var limit, _ = strconv.Atoi(lines[1])
	var walls map[Position]struct{} = make(map[Position]struct{})
	for _, line := range lines[2 : limit+2] {
		walls[parse(line)] = struct{}{}
	}
	var scores = Dijkstra(DefPosition(0, 0), funcNeighbors(size, walls))
	return scores[DefPosition(size-1, size-1)]
}

func good(walls []Position, size int, index int) bool {
	var scores = Dijkstra(DefPosition(0, 0), funcNeighbors2(size, walls, index))
	var _, ok = scores[DefPosition(size-1, size-1)]
	return ok
}

func part2(input string) string {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var size, _ = strconv.Atoi(lines[0])
	var deb = 2
	var fin = len(lines)
	var walls []Position = make([]Position, fin-deb)
	for i, line := range lines[2:] {
		walls[i] = parse(line)
	}
	for true {
		var index = (deb + fin) / 2
		var ok = good(walls, size, index)
		if ok && !good(walls, size, index+1) {
			return lines[index+2]
		}
		if ok {
			deb = index
		} else {
			fin = index
		}
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
