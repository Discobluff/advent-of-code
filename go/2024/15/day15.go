package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
	"time"

	. "github.com/Discobluff/advent-of-code/go/utils/positions"
)

//go:embed input.txt
var input string

type WiderObject struct {
	b1, b2 byte
}

func gpsCoords(grid [][]byte, char byte) int {
	var res int
	for i, line := range grid {
		for j := range line {
			if grid[i][j] == char {
				res += 100*i + j
			}
		}
	}
	return res
}

func parse1(input string) ([][]byte, []Position) {
	var inputSplit []string = strings.Split(strings.TrimSuffix(input, "\n"), "\n\n")
	var linesGrid []string = strings.Split(inputSplit[0], "\n")
	var grid [][]byte = make([][]byte, len(linesGrid))
	for i, line := range linesGrid {
		grid[i] = []byte(line)
	}
	var linesMoves []string = strings.Split(inputSplit[1], "\n")
	var positions []Position = make([]Position, len(linesMoves)*len(linesMoves[0]))
	for i, line := range linesMoves {
		for j := range line {
			positions[i*len(linesMoves[0])+j] = Directions[line[j]]
		}
	}
	return grid, positions
}

func val(grid [][]byte, pos Position) byte {
	return grid[pos.Line][pos.Column]
}

func write(grid *[][]byte, pos Position, value byte) {
	(*grid)[pos.Line][pos.Column] = value
}

func part1(input string) int {
	var grid, positions = parse1(input)
	var posRobot Position = SearchStartGrid(grid, '@')
	for _, pos := range positions {
		var newPos = AddPositions(posRobot, pos)
		if val(grid, newPos) == '.' {
			write(&grid, posRobot, '.')
			posRobot = newPos
			write(&grid, posRobot, '@')
		} else {
			if val(grid, newPos) == 'O' {
				for val(grid, newPos) == 'O' {
					newPos = AddPositions(newPos, pos)
				}
				if val(grid, newPos) == '.' {
					write(&grid, posRobot, '.')
					posRobot = AddPositions(posRobot, pos)
					write(&grid, posRobot, '@')
					for writePos := AddPositions(posRobot, pos); writePos != AddPositions(newPos, pos); writePos = AddPositions(writePos, pos) {
						write(&grid, writePos, 'O')
					}
				}
			}
		}

	}
	return gpsCoords(grid, 'O')
}

func parse2(input string) ([][]byte, []Position) {
	var inputSplit []string = strings.Split(strings.TrimSuffix(input, "\n"), "\n\n")
	var linesGrid []string = strings.Split(inputSplit[0], "\n")
	var grid [][]byte = make([][]byte, len(linesGrid))
	var replacements = map[byte]WiderObject{'@': {b1: '@', b2: '.'}, '.': {b1: '.', b2: '.'}, 'O': {b1: '[', b2: ']'}, '#': {b1: '#', b2: '#'}}
	for i, line := range linesGrid {
		grid[i] = make([]byte, 2*len(line))
		for j := 0; j < 2*len(line); j += 2 {
			grid[i][j] = replacements[line[j/2]].b1
			grid[i][j+1] = replacements[line[j/2]].b2
		}
	}
	var linesMoves []string = strings.Split(inputSplit[1], "\n")
	var positions []Position = make([]Position, len(linesMoves)*len(linesMoves[0]))
	for i, line := range linesMoves {
		for j := range line {
			positions[i*len(linesMoves[0])+j] = Directions[line[j]]
		}
	}
	return grid, positions
}

func findBoxes(grid [][]byte, box Position, direction int) []Position {
	var nextBoxes []Position
	nextBoxes = append(nextBoxes, box)
	var boxesEncountered []Position
	for len(nextBoxes) > 0 {
		var currentBox = nextBoxes[0]
		nextBoxes = nextBoxes[1:]
		if !slices.Contains(boxesEncountered, currentBox) {
			boxesEncountered = append(boxesEncountered, currentBox)
			if grid[currentBox.Line+direction][currentBox.Column] == '[' {
				nextBoxes = append(nextBoxes, Position{Line: currentBox.Line + direction, Column: currentBox.Column})
			}
			if grid[currentBox.Line+direction][currentBox.Column] == ']' {
				nextBoxes = append(nextBoxes, Position{Line: currentBox.Line + direction, Column: currentBox.Column - 1})
			}
			if grid[currentBox.Line+direction][currentBox.Column+1] == '[' {
				nextBoxes = append(nextBoxes, Position{Line: currentBox.Line + direction, Column: currentBox.Column + 1})
			}
		}
	}
	return boxesEncountered
}

func movePossible(grid [][]byte, boxes []Position, direction Position) bool {
	for _, box := range boxes {
		if val(grid, AddPositions(box, direction)) == '#' || val(grid, AddPositions(AddPositions(box, E), direction)) == '#' {
			return false
		}
	}
	return true
}

func move(grid *[][]byte, boxes []Position, direction Position) {
	for _, box := range boxes {
		write(grid, box, '.')
		write(grid, AddPositions(box, E), '.')
		write(grid, AddPositions(box, direction), '[')
		write(grid, AddPositions(AddPositions(box, E), direction), ']')
	}
}

func part2(input string) int {
	var grid, positions = parse2(input)
	var posRobot Position = SearchStartGrid(grid, '@')
	for _, pos := range positions {
		var newPos = AddPositions(posRobot, pos)
		if val(grid, newPos) == '.' {
			write(&grid, posRobot, '.')
			posRobot = newPos
			write(&grid, posRobot, '@')
		} else {
			if pos == E || pos == W {
				if val(grid, newPos) == ']' || val(grid, newPos) == '[' {
					for val(grid, newPos) == ']' || val(grid, newPos) == '[' {
						newPos = AddPositions(newPos, pos)
					}
					if val(grid, newPos) == '.' {
						write(&grid, posRobot, '.')
						posRobot = AddPositions(posRobot, pos)
						write(&grid, posRobot, '@')
						var writeByte byte = ']'
						if pos == E {
							writeByte = '['
						}
						for writePos := AddPositions(posRobot, pos); writePos != AddPositions(newPos, pos); writePos = AddPositions(writePos, pos) {
							write(&grid, writePos, writeByte)
							if writeByte == '[' {
								writeByte = ']'
							} else {
								writeByte = '['
							}
						}
					}
				}
			} else {
				//Handle the move of the boxes to the north or the south
				var boxes []Position
				if val(grid, newPos) == '[' {
					boxes = findBoxes(grid, newPos, pos.Line)
				}
				if val(grid, newPos) == ']' {
					boxes = findBoxes(grid, AddPositions(newPos, W), pos.Line)
				}
				if len(boxes) > 0 && movePossible(grid, boxes, pos) {
					slices.Reverse(boxes)
					move(&grid, boxes, pos)
					write(&grid, posRobot, '.')
					posRobot = newPos
					write(&grid, posRobot, '@')
				}
			}
		}
	}
	return gpsCoords(grid, '[')
}

func main() {
	fmt.Println("--2024 day 15 solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(input))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("part2 : ", part2(input))
	fmt.Println(time.Since(start))
}
