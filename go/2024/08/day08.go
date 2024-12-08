package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var input string

type Position struct {
	line, column int
}

func initPosition(line int, column int) Position {
	var res Position
	res.line = line
	res.column = column
	return res
}

func initMap(lines []string) map[byte][]Position {
	var res map[byte][]Position = make(map[byte][]Position)
	for i := range lines {
		for j := range lines[i] {
			var char byte = lines[i][j]
			if char != '.' {
				res[char] = append(res[char], initPosition(i, j))
			}
		}
	}
	return res
}

func initGrid(length int, height int) [][]bool {
	var res [][]bool = make([][]bool, height)
	for i := range height {
		res[i] = make([]bool, length)
	}
	return res
}

func antinode1(pos1 Position, pos2 Position) (Position, Position) {
	return initPosition(2*pos1.line-pos2.line, 2*pos1.column-pos2.column), initPosition(2*pos2.line-pos1.line, 2*pos2.column-pos1.column)
}

func antinode2(pos1 Position, pos2 Position, length int, height int) []Position {
	var res []Position //= make([]Position, 0)
	var i = 1
	var pos = initPosition(i*pos1.line-(i-1)*pos2.line, i*pos1.column-(i-1)*pos2.column)
	for isValid(pos, height, length) {
		res = append(res, pos)
		i += 1
		pos = initPosition(i*pos1.line-(i-1)*pos2.line, i*pos1.column-(i-1)*pos2.column)
	}
	i = 1
	pos = initPosition(i*pos2.line-(i-1)*pos1.line, i*pos2.column-(i-1)*pos1.column)
	for isValid(pos, height, length) {
		res = append(res, pos)
		i += 1
		pos = initPosition(i*pos2.line-(i-1)*pos1.line, i*pos2.column-(i-1)*pos1.column)
	}
	return res
}

func isValid(pos Position, height int, length int) bool {
	return pos.line < height && pos.line >= 0 && pos.column < length && pos.column >= 0
}

func countTrue(grid [][]bool) int {
	var res int
	for _, line := range grid {
		for _, val := range line {
			if val {
				res++
			}
		}
	}
	return res
}

func part1(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var dict map[byte][]Position = initMap(lines)
	var length = len(lines[0])
	var height = len(lines)
	var grid [][]bool = initGrid(length, height)
	var res int
	for _, positions := range dict {
		for i := range positions {
			for j := i + 1; j < len(positions); j++ {
				var pos1, pos2 = antinode1(positions[i], positions[j])
				if isValid(pos1, height, length) && !grid[pos1.line][pos1.column] {
					grid[pos1.line][pos1.column] = true
					res++
				}
				if isValid(pos2, height, length) && !grid[pos2.line][pos2.column] {
					grid[pos2.line][pos2.column] = true
					res++
				}
			}
		}
	}
	return res
}

func part2(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var dict map[byte][]Position = initMap(lines)
	var length = len(lines[0])
	var height = len(lines)
	var grid [][]bool = initGrid(length, height)
	var res int
	for _, positions := range dict {
		for i := range positions {
			for j := i + 1; j < len(positions); j++ {
				var slice = antinode2(positions[i], positions[j], length, height)
				for _, pos := range slice {
					if isValid(pos, height, length) && !grid[pos.line][pos.column] {
						grid[pos.line][pos.column] = true
						res++
					}
				}
			}
		}
	}
	return res
}

func main() {
	fmt.Println("--2024 day 08 solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(input))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("part2 : ", part2(input))
	fmt.Println(time.Since(start))
}
