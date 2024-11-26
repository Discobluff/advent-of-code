package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"
)

//go:embed input.txt
var inputDay string

func isPresent(list []int, elem int) bool {
	for _, val := range list {
		if val == elem {
			return true
		}
	}
	return false
}

func intersect(list1 []int, list2 []int) []int {
	var res []int
	for _, elem1 := range list1 {
		if isPresent(list2, elem1) {
			res = append(res, elem1)
		}
	}
	return res
}

func convertStringToInt(str string) int {
	var res int
	for _, char := range str {
		res += 10*res + int(char-48)
	}
	return res
}

func recomposeHand(hand []string) []int {
	var res []int
	for _, number := range hand {
		if number != "" {
			res = append(res, convertStringToInt(number))
		}
	}
	return res
}

func parseHand(hand string) []int {
	return recomposeHand(strings.Split(hand, " "))
}

func parseLine(line string) ([]int, []int) {
	line = strings.Split(line, ": ")[1]
	return parseHand(strings.Split(line, " | ")[0]), parseHand(strings.Split(line, " | ")[1])
}

func parse(lines []string) ([][]int, [][]int) {
	var res1 [][]int
	var res2 [][]int
	for _, line := range lines {
		tab1, tab2 := parseLine(line)
		res1 = append(res1, tab1)
		res2 = append(res2, tab2)
	}
	return res1, res2
}

func score(tab []int) int {
	return int(math.Pow(2., float64(len(tab)-1)))
}

func sumSlice(tab []int) int {
	var sum int
	for _, elem := range tab {
		sum += elem
	}
	return sum
}

func part1(lines []string) int {
	tab1, tab2 := parse(lines)
	var sum int
	for i := range tab1 {
		sum += score(intersect(tab1[i], tab2[i]))
	}
	return sum
}

func part2(lines []string) int {
	var compte []int = make([]int, len(lines), len(lines))
	tab1, tab2 := parse(lines)
	for i := range len(lines) {
		compte[i] = 1
	}
	for i := range len(lines) {
		var nbSame int = len(intersect(tab1[i], tab2[i]))
		for j := i + 1; j <= i+nbSame; j++ {
			compte[j] += compte[i]
		}
	}
	return sumSlice(compte)
}

func main() {
	var input = strings.TrimSuffix(inputDay, "\n")
	var lines = strings.Split(input, "\n")
	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}
