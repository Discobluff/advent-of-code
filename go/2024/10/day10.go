package main

import (
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Position struct {
	line, column int
}

func up(pos Position) Position {
	pos.line--
	return pos
}

func down(pos Position) Position {
	pos.line++
	return pos
}

func left(pos Position) Position {
	pos.column--
	return pos
}

func right(pos Position) Position {
	pos.column++
	return pos
}

func solve1(grid [][]int, pos Position, visited *[][]bool) int {
	var val int = grid[pos.line][pos.column]
	if val == 9 {
		if !(*visited)[pos.line][pos.column] {
			(*visited)[pos.line][pos.column] = true
			return 1
		}
		return 0
	}
	var res int
	if pos.line > 0 && grid[pos.line-1][pos.column] == val+1 {
		res += solve1(grid, up(pos), visited)
	}
	if pos.line < len(grid)-1 && grid[pos.line+1][pos.column] == val+1 {
		res += solve1(grid, down(pos), visited)
	}
	if pos.column > 0 && grid[pos.line][pos.column-1] == val+1 {
		res += solve1(grid, left(pos), visited)
	}
	if pos.column < len(grid[0])-1 && grid[pos.line][pos.column+1] == val+1 {
		res += solve1(grid, right(pos), visited)
	}
	return res
}
func solve2(grid [][]int, pos Position) int {
	var val int = grid[pos.line][pos.column]
	if val == 9 {
		return 1
	}
	var res int
	if pos.line > 0 && grid[pos.line-1][pos.column] == val+1 {
		res += solve2(grid, up(pos))
	}
	if pos.line < len(grid)-1 && grid[pos.line+1][pos.column] == val+1 {
		res += solve2(grid, down(pos))
	}
	if pos.column > 0 && grid[pos.line][pos.column-1] == val+1 {
		res += solve2(grid, left(pos))
	}
	if pos.column < len(grid[0])-1 && grid[pos.line][pos.column+1] == val+1 {
		res += solve2(grid, right(pos))
	}
	return res
}

func parse(lines []string) [][]int {
	var res [][]int = make([][]int, len(lines))
	for i, line := range lines {
		res[i] = make([]int, len(line))
		for j, char := range line {
			res[i][j], _ = strconv.Atoi(string(char))
		}
	}
	return res
}

func createPosition(line int, column int) Position {
	var pos Position
	pos.line = line
	pos.column = column
	return pos
}

func initVisited(height int, length int) *[][]bool {
	var res [][]bool = make([][]bool, height)
	for i := range height {
		res[i] = make([]bool, length)
	}
	return &res
}

func part1(input string) int {
	var grid = parse(strings.Split(strings.TrimSuffix(input, "\n"), "\n"))
	var res int
	for i, line := range grid {
		for j, val := range line {
			if val == 0 {
				res += solve1(grid, createPosition(i, j), initVisited(len(grid), len(grid[0])))
			}
		}
	}
	return res
}

func part2(input string) int {
	var grid = parse(strings.Split(strings.TrimSuffix(input, "\n"), "\n"))
	var res int
	for i, line := range grid {
		for j, val := range line {
			if val == 0 {
				res += solve2(grid, createPosition(i, j))
			}
		}
	}
	return res
}

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("--2024 day 10 solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(string(input)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("part2 : ", part2(string(input)))
	fmt.Println(time.Since(start))
}
