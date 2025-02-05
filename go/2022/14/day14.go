package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
	"time"

	. "github.com/Discobluff/advent-of-code/go/utils/positions"
)

func sgn(a int) int {
	if a == 0 {
		return 0
	}
	if a > 0 {
		return 1
	}
	return -1
}

func getSize(lines []string) (int, int) {
	var maxc, maxl int
	for _, line := range lines {
		for _, coords := range strings.Split(line, " -> ") {
			var l, c int
			fmt.Sscanf(coords, "%d,%d", &c, &l)
			if l > maxl {
				maxl = l
			}
			if c > maxc {
				maxc = c
			}
		}
	}
	return maxl + 1, maxc + 1
}

func parse(lines []string) [][]byte {
	var height, length = getSize(lines)
	var grid [][]byte = make([][]byte, height)
	for i := range height {
		grid[i] = make([]byte, length)
		for j := range length {
			grid[i][j] = '.'
		}
	}
	for _, line := range lines {
		var split = strings.Split(line, " -> ")
		for i := 0; i < len(split)-1; i++ {
			var l1, c1, l2, c2 int
			fmt.Sscanf(split[i], "%d,%d", &c1, &l1)
			fmt.Sscanf(split[i+1], "%d,%d", &c2, &l2)
			for l1 != l2 || c1 != c2 {
				grid[l1][c1] = '#'
				l1 += sgn(l2 - l1)
				c1 += sgn(c2 - c1)
			}
			grid[l1][c1] = '#'
		}
	}
	return grid
}

func parse2(lines []string) [][]byte {
	var grid = parse(lines)
	grid = append(grid, make([]byte, len(grid[0])))
	grid = append(grid, make([]byte, len(grid[0])))
	for i := range grid[0] {
		grid[len(grid)-2][i] = '.'
		grid[len(grid)-1][i] = '#'
	}
	return grid
}

func eval(grid [][]byte, pos Position) byte {
	return grid[pos.Line][pos.Column]
}

func move(grid [][]byte, sand Position) (bool, Position) {
	if sand.Column >= len(grid[0])-1 {
		grid = addColumn(grid)
	}
	if sand.Line >= len(grid)-1 {
		return false, Position{Line: 0, Column: 0}
	}
	var sandS = AddPositions(sand, S)
	if eval(grid, sandS) == '.' {
		return move(grid, sandS)
	}
	var sandSW = AddPositions(sand, AddPositions(S, W))
	if eval(grid, sandSW) == '.' {
		return move(grid, sandSW)
	}
	var sandSE = AddPositions(sand, AddPositions(S, E))
	if eval(grid, sandSE) == '.' {
		return move(grid, sandSE)
	}
	return true, sand
}

func printt(grid [][]byte) {
	for _, line := range grid {
		for _, char := range line[494:] {
			fmt.Printf(string(char))
		}
		fmt.Printf("\n")
	}
	fmt.Println("----------------------------")
}

func part1(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var grid = parse(lines)
	var start = Position{Line: 0, Column: 500}
	grid[start.Line][start.Column] = '+'
	var ok bool = true
	var res int
	var sand = start
	for ok {
		res++
		ok, sand = move(grid, start)
		grid[sand.Line][sand.Column] = 'o'
	}
	return res - 1
}

func addColumn(grid [][]byte) [][]byte {
	for i := range grid {
		if i < len(grid)-1 {
			grid[i] = append(grid[i], '.')
		} else {
			grid[i] = append(grid[i], '#')
		}
	}
	return grid
}

func part2(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var grid = parse2(lines)
	var start = Position{Line: 0, Column: 500}
	grid[start.Line][start.Column] = '+'
	var ok bool = true
	var res int
	var sand = start
	for ok {
		res++
		_, sand = move(grid, start)
		grid[sand.Line][sand.Column] = 'o'
		ok = sand != start
	}
	return res
}

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("--2022 day 14 solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(string(input)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("part2 : ", part2(string(input)))
	fmt.Println(time.Since(start))
}
