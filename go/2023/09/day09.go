package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var inputDay string

func convertStringToInt(str string) int {
	var res int
	var opp int = 1
	for _, char := range str {
		if char == 45 {
			opp = -1
		} else {
			res = 10*res + int(char-'0')
		}
	}
	return opp * res
}

func convertStringToSliceInt(str []string) []int {
	var res []int = make([]int, len(str), len(str))
	for i, chain := range str {
		if chain != "" {
			res[i] = convertStringToInt(chain)
		}
	}
	return res
}

func isZero(tab []int) bool {
	for _, elem := range tab {
		if elem != 0 {
			return false
		}
	}
	return true
}

func buildNextTab(tab []int) []int {
	var res = make([]int, len(tab)-1)
	for i := range len(tab) - 1 {
		res[i] = tab[i+1] - tab[i]
	}
	return res
}

func solve1(numbers []int) int {
	if isZero(numbers) {
		return 0
	}
	return numbers[len(numbers)-1] + solve1(buildNextTab(numbers))
}

func solve2(numbers []int) int {
	if isZero(numbers) {
		return 0
	}
	var val = solve2(buildNextTab(numbers))
	return numbers[0] - val
}

func part1(lines []string) int {
	var res int
	for _, line := range lines {
		res += solve1(convertStringToSliceInt(strings.Split(line, " ")))
	}
	return res
}

func part2(lines []string) int {
	var res int
	for _, line := range lines {
		res += solve2(convertStringToSliceInt(strings.Split(line, " ")))
	}
	return res
}

func main() {
	var input = strings.TrimSuffix(inputDay, "\n")
	var lines = strings.Split(input, "\n")
	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}
