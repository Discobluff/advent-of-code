package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var inputDay string

// 0 : line, 1 : column, 2 : direction
type direction [3]int

func isPoint(char byte) bool {
	return char == 46
}

// 0 = left, 1 = right, 2 = up, 3 = bottom
func propagate(grid []string, dir direction, laser *[][]bool) (int, int, byte) {
	var char byte
	line, column, sens := decomposeDirection(dir)
	if sens == 0 {
		var res int = column + 1
		for j := column; j >= 0 && isPoint(grid[line][j]); j-- {
			(*laser)[line][j] = true
			res = j
		}
		res = max(0, res-1)
		char = grid[line][res]
		return line, res, char
	}
	if sens == 1 {
		var res int = column - 1
		for j := column; j <= len(grid[line])-1 && isPoint(grid[line][j]); j++ {
			(*laser)[line][j] = true
			res = j
		}
		res = min(res+1, len(grid[line])-1)
		char = grid[line][res]
		return line, res, char
	}
	if sens == 2 {
		var res = line + 1
		for i := line; i >= 0 && isPoint(grid[i][column]); i-- {
			(*laser)[i][column] = true
			res = i
		}
		res = max(0, res-1)
		char = grid[res][column]
		return res, column, char
	}
	var res = line - 1
	for i := line; i <= len(grid)-1 && isPoint((grid[i][column])); i++ {
		(*laser)[i][column] = true
		res = i
	}
	res = min(res+1, len(grid)-1)
	char = grid[res][column]
	return res, column, char
}

func createDirection(line int, column int, sens int) direction {
	var res direction
	res[0] = line
	res[1] = column
	res[2] = sens
	return res
}

func decomposeDirection(dir direction) (int, int, int) {
	return dir[0], dir[1], dir[2]
}

func countTrue(laser [][]bool) int {
	var res int
	for i := range laser {
		for j := range laser[i] {
			if laser[i][j] {
				res += 1
			}
		}
	}
	return res
}

func isPresent(directionList []direction, directionToCheck direction) bool {
	for _, dir := range directionList {
		if dir == directionToCheck {
			return true
		}
	}
	return false
}

func printGrid(laser [][]bool) {
	for _, line := range laser {
		for _, b := range line {
			if b {
				fmt.Printf("1 ")
			} else {
				fmt.Printf("0 ")
			}
		}
		fmt.Printf("\n")
	}
}

func solve(grid []string, start direction) int {
	laser := make([][]bool, len(grid))
	for i := range laser {
		laser[i] = make([]bool, len(grid[0]))
	}
	var directionEncountered []direction = make([]direction, 0, 1)
	var directionToBrowse []direction = make([]direction, 1, 1)
	directionToBrowse[0] = start
	for len(directionToBrowse) != 0 {
		if !isPresent(directionEncountered, directionToBrowse[0]) {
			var sens int = directionToBrowse[0][2]
			directionEncountered = append(directionEncountered, directionToBrowse[0])
			line, column, symbol := propagate(grid, directionToBrowse[0], &laser)
			laser[line][column] = true
			// Symbole /
			if symbol == 47 {
				if sens == 0 && line+1 <= len(grid)-1 {
					directionToBrowse = append(directionToBrowse, createDirection(line+1, column, 3))
				}
				if sens == 1 && line-1 >= 0 {
					directionToBrowse = append(directionToBrowse, createDirection(line-1, column, 2))
				}
				if sens == 2 && column+1 <= len(grid[line])-1 {
					directionToBrowse = append(directionToBrowse, createDirection(line, column+1, 1))
				}
				if sens == 3 && column-1 >= 0 {
					directionToBrowse = append(directionToBrowse, createDirection(line, column-1, 0))
				}
			}
			// Symbole \
			if symbol == 92 {
				if sens == 0 && line-1 >= 0 {
					directionToBrowse = append(directionToBrowse, createDirection(line-1, column, 2))
				}
				if sens == 1 && line+1 <= len(grid)-1 {
					directionToBrowse = append(directionToBrowse, createDirection(line+1, column, 3))
				}
				if sens == 2 && column-1 >= 0 {
					directionToBrowse = append(directionToBrowse, createDirection(line, column-1, 0))
				}
				if sens == 3 && column+1 <= len(grid[line])-1 {
					directionToBrowse = append(directionToBrowse, createDirection(line, column+1, 1))
				}
			}
			// Symbole -
			if symbol == 45 {
				// Si le sens est horizontal on ne change rien
				if sens == 0 && column-1 >= 0 {
					directionToBrowse = append(directionToBrowse, createDirection(line, column-1, 0))
				}
				if sens == 1 && column+1 <= len(grid[line])-1 {
					directionToBrowse = append(directionToBrowse, createDirection(line, column+1, 1))
				} else {
					if column+1 <= len(grid[line])-1 {
						directionToBrowse = append(directionToBrowse, createDirection(line, column+1, 1))
					}
					if column-1 >= 0 {
						directionToBrowse = append(directionToBrowse, createDirection(line, column-1, 0))
					}
				}
			}
			// Symbole |
			if symbol == 124 {
				// Si le sens est vertical on ne change rien
				if sens == 2 && line-1 >= 0 {
					directionToBrowse = append(directionToBrowse, createDirection(line-1, column, 2))
				}
				if sens == 3 && line+1 <= len(grid)-1 {
					directionToBrowse = append(directionToBrowse, createDirection(line+1, column, 3))
				} else {
					if line-1 >= 0 {
						directionToBrowse = append(directionToBrowse, createDirection(line-1, column, 2))
					}
					if line+1 <= len(grid)-1 {
						directionToBrowse = append(directionToBrowse, createDirection(line+1, column, 3))
					}
				}
			}
		}
		directionToBrowse = directionToBrowse[1:]
	}
	return countTrue(laser)

}

func part1(lines []string) int {
	return solve(lines, createDirection(0, 0, 1))
}

func main() {
	var input = strings.TrimSuffix(inputDay, "\n")
	var lines = strings.Split(input, "\n")
	fmt.Println(part1(lines))
}
