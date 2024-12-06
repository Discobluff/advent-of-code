package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
	"time"
)

//go:embed input.txt
var input string

type Case struct {
	direction []int
	visited   bool
}

func searchStart(grid []string, char byte) (int, int) {
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == char {
				return i, j
			}
		}
	}
	return -1, -1
}

// 0 left, 1 up, 2 right, 3 bottom
func nextCase(grid []string, direction int, x int, y int) (int, int, int, bool) { // direction,nextX,nextY,isOutside
	if (direction == 0 && y == 0) || (direction == 1 && x == 0) || (direction == 2 && y == len(grid[0])-1) || (direction == 3 && x == len(grid)-1) {
		return -1, -1, -1, true
	}
	var directions [4][2]int = [4][2]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}
	var xShift int = x + directions[direction][0]
	var yShift int = y + directions[direction][1]
	if grid[xShift][yShift] == '#' {
		if xShift < 0 || yShift < 0 || xShift >= len(grid) || yShift >= len(grid[0]) {
			return -1, -1, -1, true
		}
		return (direction + 1) % 4, x, y, false
	}
	return direction, xShift, yShift, false
}

func initCase(visit bool) Case {
	var newCase Case
	newCase.direction = make([]int, 0)
	newCase.visited = visit
	return newCase
}

func initGrid(height int, length int) [][]Case {
	var res [][]Case = make([][]Case, height)
	for i := range height {
		res[i] = make([]Case, length)
		for j := range length {
			res[i][j] = initCase(false)
		}
	}
	return res
}

func solve(lines []string) (int, bool) {
	var grid = initGrid(len(lines), len(lines[0]))
	var direction int = 1
	var x, y = searchStart(lines, '^')
	var res int
	var leaveGrid bool
	var loop bool
	for !leaveGrid && !loop {
		if !grid[x][y].visited {
			res++
		}
		grid[x][y].visited = true
		grid[x][y].direction = append(grid[x][y].direction, direction)
		direction, x, y, leaveGrid = nextCase(lines, direction, x, y)
		if !leaveGrid {
			loop = slices.Contains(grid[x][y].direction, direction)
		}
	}
	return res, loop
}

func part1(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var res, _ = solve(lines)
	return res
}

func muteString(str string, index int, char byte) string {
	var res []byte = make([]byte, len(str))
	for j := range str {
		if j == index {
			res[j] = char
		} else {
			res[j] = str[j]
		}
	}
	return string(res)
}

// func cop(lines []string) []string {

// }

func part2(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var x, y = searchStart(lines, '^')
	var res int
	// var cop []string = copys(lines)

	// fmt.Println("non")
	// fmt.Println("bart", slices.Equal(lines, cop))
	for i := range lines {
		for j := range lines[i] {
			if x != i || y != j {
				var char = lines[i][j]
				if char != '#' {
					lines[i] = muteString(lines[i], j, '#')
					var _, ok = solve(lines)
					if ok {
						res++
					}
					lines[i] = muteString(lines[i], j, char)
					// if !slices.Equal(lines, cop) {
					// 	fmt.Println("alerte")
					// }
					// fmt.Println(slices.Equal(lines, cop))
				}
			}
		}
	}
	return res
}

func main() {
	fmt.Println("--2024 day 06 solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(input))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("part2 : ", part2(input))
	fmt.Println(time.Since(start))
}
