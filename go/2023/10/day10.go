package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed test.txt
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

func touchSide(Case [2]int, height int, length int) bool {
	return Case[0] == 0 || Case[1] == 0 || Case[0] == height-1 || Case[1] == length-1
}

func touchSideCases(sliceCase [][2]int, height int, length int) bool {
	for _, Case := range sliceCase {
		if touchSide(Case, height, length) {
			return true
		}
	}
	return false
}

func isPresent(sliceCase [][2]int, Case [2]int) bool {
	for _, elem := range sliceCase {
		if elem == Case {
			return true
		}
	}
	return false
}

func getNeighbours(Case [2]int, height int, length int) [][2]int {
	var res [][2]int
	var c1, c2, c3, c4 = decaleCase(Case, 0, 1), decaleCase(Case, 0, -1), decaleCase(Case, 1, 1), decaleCase(Case, 1, -1)
	if valable(c1, length, height) {
		res = append(res, c1)
	}
	if valable(c2, length, height) {
		res = append(res, c2)
	}
	if valable(c3, length, height) {
		res = append(res, c3)
	}
	if valable(c4, length, height) {
		res = append(res, c4)
	}
	return res
}

func getSliceCases(grid [][]int, val int) [][2]int {
	var res [][2]int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == val {
				var Case [2]int
				Case[0] = i
				Case[1] = j
				res = append(res, Case)
			}
		}
	}
	return res
}

func buildGroupCases(grid [][]int, start [2]int, height int, length int, val int) [][]int {
	// var res [][2]int
	grid[start[0]][start[1]] = val
	for _, c := range getNeighbours(start, height, length) {
		if grid[c[0]][c[1]] == -1 {
			buildGroupCases(grid, c, height, length, val)
		}
	}
	return grid
}

//1111
//1001
//1001
//1111

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

func parcours1(lines []string, start [2]int, previous [2]int, next [2]int, tab *[][]int) {
	if next != start {
		if (*tab)[next[0]][next[1]] == -1 {
			(*tab)[next[0]][next[1]] = (*tab)[previous[0]][previous[1]] + 1
		} else {
			(*tab)[next[0]][next[1]] = min((*tab)[next[0]][next[1]], (*tab)[previous[0]][previous[1]]+1)
		}
		var symbol = lines[next[0]][next[1]]
		parcours1(lines, start, next, getNextCase(symbol, next, previous), tab)
	}
}

func maxTab(tab [][]int) int {
	var res int = -1
	for i := range tab {
		for j := range tab[i] {
			if res < tab[i][j] {
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
	parcours1(lines, start, start, next1, &tab)
	parcours1(lines, start, start, next2, &tab)
	return maxTab(tab)
}

func parcours2(lines []string, start [2]int, previous [2]int, next [2]int, tab *[][]int) {
	if next != start {
		(*tab)[next[0]][next[1]] = 1
		var symbol = lines[next[0]][next[1]]
		parcours2(lines, start, next, getNextCase(symbol, next, previous), tab)
	}
}

func searchEnclosed(tab *[][]int) {
	// -1 : we don't know, 1 : tile of the loop, 0 : tile not of the loop not enclosed, 2 : tile not of the loop
}

func isPresentSlice(grid [][]int, elem int) ([2]int, bool) {
	var res [2]int
	res[0] = -1
	res[1] = -1
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == elem {
				res[0] = i
				res[1] = j
				return res, true
			}
		}
	}
	return res, false
}

func part2(lines []string) int {
	var start [2]int = search(lines, 'S')
	var tab = initTab(len(lines[0]), len(lines))
	var next1, _ = getCases(getSymbolStart(lines, start), start)
	tab[start[0]][start[1]] = 1
	parcours2(lines, start, start, next1, &tab)
	var cas [2]int
	cas[0] = 6
	cas[1] = 2
	var l [][2]int
	l = append(l, cas)
	var groupsCases [][][2]int
	var count int = 2
	var Case, ok = isPresentSlice(tab, -1)
	for ok {
		tab = buildGroupCases(tab, Case, len(lines), len(lines[0]), count)
		groupsCases = append(groupsCases, getSliceCases(tab, count))
		count++
		Case, ok = isPresentSlice(tab, -1)
	}
	var res int
	var i int = 2
	for _, group := range groupsCases {
		if !touchSideCases(group, len(lines), len(lines[0])) {
			fmt.Println(i, len(group))
			res += len(group)
		}
		i++
	}
	printTab(tab)
	return res
}

func main() {
	var input = strings.TrimSuffix(inputDay, "\n")
	var lines = strings.Split(input, "\n")
	// fmt.Println(part1(lines))
	fmt.Println(part2(lines))

}
