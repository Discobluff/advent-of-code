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

func antinode(pos1 Position, pos2 Position, distance int) (Position, Position) {
	return initPosition(distance*pos1.line-(distance-1)*pos2.line, distance*pos1.column-(distance-1)*pos2.column), initPosition(distance*pos2.line-(distance-1)*pos1.line, distance*pos2.column-(distance-1)*pos1.column)
}

func antinode2(pos1 Position, pos2 Position, length int, height int, grid *[][]bool) int {
	var i int = 1
	var res int
	var position1, position2 = initPosition(i*pos1.line-(i-1)*pos2.line, i*pos1.column-(i-1)*pos2.column), initPosition(i*pos2.line-(i-1)*pos1.line, i*pos2.column-(i-1)*pos1.column)
	for isValid(position1, height, length) || isValid(position2, height, length) {
		if isValid(position1, height, length) && !(*grid)[position1.line][position1.column] {
			(*grid)[position1.line][position1.column] = true
			res++
		}
		if isValid(position2, height, length) && !(*grid)[position2.line][position2.column] {
			(*grid)[position2.line][position2.column] = true
			res++
		}
		i += 1
		position1, position2 = antinode(pos1, pos2, i)
	}
	return res
}

func isValid(pos Position, height int, length int) bool {
	return pos.line < height && pos.line >= 0 && pos.column < length && pos.column >= 0
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
				var pos1, pos2 = antinode(positions[i], positions[j], 2)
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
				res += antinode2(positions[i], positions[j], length, height, &grid)
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
