package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var inputDay string

func initTab(length int, height int) [][]int {
	var tab [][]int = make([][]int, height)
	for i := range height {
		tab[i] = make([]int, length)
		for j := range length {
			tab[i][j] = -1
		}
	}
	return tab
}

func getSymbolStart(lines []string, Case [2]int) byte {
	var north, east, south, west bool
	if lines[Case[0]][Case[1]+1] == '-' || lines[Case[0]][Case[1]+1] == '7' || lines[Case[0]][Case[1]+1] == 'J' {
		east = true
	}
	if lines[Case[0]][Case[1]-1] == '-' || lines[Case[0]][Case[1]-1] == 'F' || lines[Case[0]][Case[1]-1] == 'L' {
		west = true
	}
	if lines[Case[0]+1][Case[1]] == '|' || lines[Case[0]+1][Case[1]] == 'L' || lines[Case[0]+1][Case[1]] == 'J' {
		south = true
	}
	if lines[Case[0]-1][Case[1]] == '|' || lines[Case[0]-1][Case[1]] == 'F' || lines[Case[0]-1][Case[1]] == '7' {
		north = true
	}
	if north {
		if south {
			return '|'
		}
		if east {
			return 'L'
		}
		if west {
			return 'J'
		}
	}
	if south {
		if east {
			return 'F'
		}
		if west {
			return '7'
		}
	}
	return '-'

}

func decaleCase(Case [2]int, index int, decalage int) [2]int {
	var res [2]int
	res[0] = Case[0]
	res[1] = Case[1]
	res[index] += decalage
	return res
}

func getCases(symbol byte, caseSymbol [2]int) ([2]int, [2]int) {
	if symbol == '|' {
		return decaleCase(caseSymbol, 0, -1), decaleCase(caseSymbol, 0, 1)
	}
	if symbol == '-' {
		return decaleCase(caseSymbol, 1, -1), decaleCase(caseSymbol, 1, 1)
	}
	if symbol == '7' {
		return decaleCase(caseSymbol, 1, -1), decaleCase(caseSymbol, 0, 1)
	}
	if symbol == 'F' {
		return decaleCase(caseSymbol, 1, 1), decaleCase(caseSymbol, 0, 1)
	}
	if symbol == 'L' {
		return decaleCase(caseSymbol, 0, -1), decaleCase(caseSymbol, 1, 1)
	}
	//J
	return decaleCase(caseSymbol, 1, -1), decaleCase(caseSymbol, 0, -1)
}

func getNextCase(symbol byte, caseSymbole [2]int, previousCase [2]int) [2]int {
	case1, case2 := getCases(symbol, caseSymbole)
	if case1 == previousCase {
		return case2
	}
	if case2 == previousCase {
		return case1
	}
	var nilCase [2]int
	nilCase[0] = -1
	nilCase[1] = -1
	return nilCase
}

func valable(Case [2]int, length int, height int) bool {
	return !(Case[0] < 0 || Case[1] < 0 || Case[0] >= height || Case[1] >= length)
}

func search(lines []string, char byte) [2]int {
	var res [2]int
	res[0] = -1
	res[1] = -1
	for i := range len(lines) {
		for j := range len(lines[i]) {
			if lines[i][j] == char {
				res[0] = i
				res[1] = j
				return res
			}
		}
	}
	return res
}

func printTab(lines [][]int) {
	for _, line := range lines {
		fmt.Println(line)
	}
}

func parcours(lines []string, start [2]int, previous [2]int, next [2]int, tab *[][]int) {
	if next != start {
		if (*tab)[next[0]][next[1]] == -1 {
			(*tab)[next[0]][next[1]] = (*tab)[previous[0]][previous[1]] + 1
		} else {
			(*tab)[next[0]][next[1]] = min((*tab)[next[0]][next[1]], (*tab)[previous[0]][previous[1]]+1)
		}
		var symbol = lines[next[0]][next[1]]
		parcours(lines, start, next, getNextCase(symbol, next, previous), tab)
	}
}

func maxTab(tab [][]int) int {
	var res int = -1
	for i := range tab {
		for j := range tab[i] {
			if res == -1 || res < tab[i][j] {
				res = tab[i][j]
			}
		}
	}
	return res
}

func part1(lines []string) int {
	var start [2]int = search(lines, 'S')
	var tab = initTab(len(lines[0]), len(lines))
	var next1, next2 = getCases(getSymbolStart(lines, start), start)
	tab[start[0]][start[1]] = 0
	parcours(lines, start, start, next1, &tab)
	parcours(lines, start, start, next2, &tab)
	return maxTab(tab)
}

func main() {
	var input = strings.TrimSuffix(inputDay, "\n")
	var lines = strings.Split(input, "\n")
	fmt.Println(part1(lines))

}
