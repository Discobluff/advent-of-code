package main

import (
	_ "embed"
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

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

func solve(lines []string, grid [][]Case, direction int, x int, y int, res int) (int, bool) {
	if slices.Contains(grid[x][y].direction, direction) {
		return res, true
	}
	var leaveGrid bool
	var newRes int = res
	if !grid[x][y].visited {
		newRes = res + 1
	}
	grid[x][y].visited = true
	grid[x][y].direction = append(grid[x][y].direction, direction)
	direction, x, y, leaveGrid = nextCase(lines, direction, x, y)
	if leaveGrid {
		return newRes, false
	}
	return solve(lines, grid, direction, x, y, newRes)
}

func part1(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var grid = initGrid(len(lines), len(lines[0]))
	var x, y = searchStart(lines, '^')
	var res, _ = solve(lines, grid, 1, x, y, 0)
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

func part2(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var grid = initGrid(len(lines), len(lines[0]))
	var x, y = searchStart(lines, '^')
	var dirS = 1
	var xS, yS = searchStart(lines, '^')
	var res int
	var leaveGrid bool
	var loop bool
	var direction int = 1
	var nextX int
	var nextY int
	var nextDir int
	var parcourus [][2]int
	for !leaveGrid && !loop {
		if slices.Contains(grid[x][y].direction, direction) {
			return res
		}
		grid[x][y].visited = true
		grid[x][y].direction = append(grid[x][y].direction, direction)
		nextDir, nextX, nextY, leaveGrid = nextCase(lines, direction, x, y)
		if leaveGrid {
			return res
		}
		if lines[nextX][nextY] != '#' && !slices.Contains(parcourus, [2]int{nextX, nextY}) {
			parcourus = append(parcourus, [2]int{nextX, nextY})
			var line = lines[nextX]
			lines[nextX] = muteString(lines[nextX], nextY, '#')
			var _, ok = solve(lines, initGrid(len(lines), len(lines[0])), dirS, xS, yS, 0)
			if ok {
				res++
			}
			lines[nextX] = line
		}
		direction = nextDir
		x = nextX
		y = nextY
	}
	return res
}

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("--2024 day 06 solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(string(input)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("part2 : ", part2(string(input)))
	fmt.Println(time.Since(start))
}
