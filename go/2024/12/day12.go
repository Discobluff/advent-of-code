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

var N Position = Position{line: -1, column: 0}
var S Position = Position{line: 1, column: 0}
var E Position = Position{line: 0, column: 1}
var W Position = Position{line: 0, column: -1}
var directions []Position = []Position{N, S, E, W}
var directionsLine []Position = []Position{N, S}
var directionsColumn []Position = []Position{E, W}

func isValid(position Position, height int, length int) bool {
	return position.line >= 0 && position.column >= 0 && position.line < height && position.column < length
}

func addPosition(pos1 Position, pos2 Position) Position {
	return Position{line: pos1.line + pos2.line, column: pos1.column + pos2.column}
}

func valueMatrix(grid []string, position Position) byte {
	return grid[position.line][position.column]
}

func explore(grid []string, plant byte, position Position, visited *map[Position]struct{}) {
	(*visited)[position] = struct{}{}
	for _, direction := range directions {
		var nextPosition Position = addPosition(position, direction)
		if isValid(nextPosition, len(grid), len(grid[0])) && valueMatrix(grid, nextPosition) == plant {
			var _, ok = (*visited)[nextPosition]
			if !ok {
				explore(grid, plant, nextPosition, visited)
			}
		}
	}
}

func initSlice(length int, height int) [][]bool {
	var res [][]bool = make([][]bool, height)
	for i := range height {
		res[i] = make([]bool, length)
	}
	return res
}

func area(region map[Position]struct{}) int {
	return len(region)
}

func perimeter(region map[Position]struct{}) int {
	var res int
	for position := range region {
		res += 4
		for _, direction := range directions {
			var _, ok = region[addPosition(position, direction)]
			if ok {
				res--
			}
		}
	}
	return res
}

type Line struct {
	Position
	length int
}

type Column struct {
	Position
	length int
}

func corners(region map[Position]struct{}) (Position, Position) {
	var minLine, maxLine, minColumn, maxColumn int = -1, -1, -1, -1
	for key := range region {
		if minLine == -1 || minLine > key.line {
			minLine = key.line
		}
		if maxLine == -1 || maxLine < key.line {
			maxLine = key.line
		}
		if minColumn == -1 || minColumn > key.column {
			minColumn = key.column
		}
		if maxColumn == -1 || maxColumn < key.column {
			maxColumn = key.column
		}
	}
	return Position{line: minLine, column: minColumn}, Position{line: maxLine, column: maxColumn}
}

func sides(region map[Position]struct{}) int {
	//Parcours horizontal pour créer des lignes
	var corner1, corner2 = corners(region)
	var lines []Line
	for line := corner1.line; line <= corner2.line; line++ {
		for _, direction := range directionsLine {
			var deb = Position{line: 0, column: 0}
			var continueLine = false
			for column := corner1.column; column <= corner2.column; column++ {
				var _, ok1 = region[Position{line: line, column: column}]
				if ok1 {
					var newPos = addPosition(Position{line: line, column: column}, direction)
					var _, ok = region[newPos]
					if !ok {
						if !continueLine {
							continueLine = true
							deb = newPos
							lines = append(lines, Line{Position: deb, length: 0})
						}
						lines[len(lines)-1].length += 1
					} else {
						continueLine = false
					}
				} else {
					continueLine = false
				}
			}
		}
	}
	//Parcours vertical pour créer des colonnes
	var columns []Column
	for column := corner1.column; column <= corner2.column; column++ {
		for _, direction := range directionsColumn {
			var deb = Position{line: 0, column: 0}
			var continueColumn = false
			for line := corner1.line; line <= corner2.line; line++ {
				var _, ok1 = region[Position{line: line, column: column}]
				if ok1 {
					var newPos = addPosition(Position{line: line, column: column}, direction)
					var _, ok = region[newPos]
					if !ok {
						if !continueColumn {
							continueColumn = true
							deb = newPos
							columns = append(columns, Column{Position: deb, length: 0})
						}
						columns[len(columns)-1].length += 1
					} else {
						continueColumn = false
					}
				} else {
					continueColumn = false
				}
			}
		}
	}
	return len(lines) + len(columns)
}

func part1(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var parcourus [][]bool = initSlice(len(lines[0]), len(lines))
	var res int
	for i, line := range lines {
		for j := range line {
			if !parcourus[i][j] {
				var visited map[Position]struct{} = make(map[Position]struct{})
				explore(lines, lines[i][j], Position{line: i, column: j}, &visited)
				res += area(visited) * perimeter(visited)
				for key := range visited {
					parcourus[key.line][key.column] = true
				}
			}
		}
	}
	return res
}

func part2(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var parcourus [][]bool = initSlice(len(lines[0]), len(lines))
	var res int
	for i, line := range lines {
		for j := range line {
			if !parcourus[i][j] {
				var visited map[Position]struct{} = make(map[Position]struct{})
				explore(lines, lines[i][j], Position{line: i, column: j}, &visited)
				var side = sides(visited)
				res += area(visited) * side
				for key := range visited {
					parcourus[key.line][key.column] = true
				}
			}
		}
	}
	return res
}

func main() {
	fmt.Println("--2024 day 12 solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(input))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("part2 : ", part2(input))
	fmt.Println(time.Since(start))
}
