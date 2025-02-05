package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
	"time"

	. "github.com/Discobluff/advent-of-code/go/utils/positions"
	. "github.com/Discobluff/advent-of-code/go/utils/set"
)

var corres = map[string]Position{"R": E, "U": N, "L": W, "D": S}
var dirs = []Position{AddPositions(N, E), AddPositions(N, W), AddPositions(S, E), AddPositions(S, W), N, S, E, W, {Line: 0, Column: 0}}

func parse(lines []string) []Position {
	var res []Position
	for _, line := range lines {
		var move string
		var count int
		fmt.Sscanf(line, "%s %d", &move, &count)
		for range count {
			res = append(res, corres[move])
		}
	}
	return res
}

func isAround(p1 Position, p2 Position) bool {
	for _, direction := range dirs {
		if AddPositions(p1, direction) == p2 {
			return true
		}
	}
	return false
}

func move(m Position, head Position, tail Position, moveHead bool) (Position, Position) {
	var newHead = head
	if moveHead {
		newHead = AddPositions(m, head)
	}
	if isAround(newHead, tail) {
		return newHead, tail
	}
	var newTail = tail
	var best = Distance(tail, newHead)
	for _, dir := range dirs {
		var newPos = AddPositions(tail, dir)
		if best > Distance(newHead, newPos) {
			best = Distance(newHead, newPos)
			newTail = newPos
		}
	}
	return newHead, newTail
}

func part1(input string) int {
	return solve(input, 2)
}

func solve(input string, size int) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var moves = parse(lines)
	var rope []Position = make([]Position, size)
	var positions = DefSet[Position]()
	for _, m := range moves {
		for i := range size - 1 {
			rope[i], rope[i+1] = move(m, rope[i], rope[i+1], i == 0)
		}
		Add(positions, rope[size-1])
	}
	return len(positions)
}

func part2(input string) int {
	return solve(input, 10)
}

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("--2022 day 09 solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(string(input)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("part2 : ", part2(string(input)))
	fmt.Println(time.Since(start))
}
