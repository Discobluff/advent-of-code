package main

import (
	_ "embed"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func mergeTab(tab1 []int, tab2 []int) []int {
	var res []int = make([]int, len(tab1)+len(tab2))
	var i int
	var j int
	for i+j < len(tab1)+len(tab2) {
		if i >= len(tab1) {
			res[i+j] = tab2[j]
			j++
		} else {
			if j >= len(tab2) {
				res[i+j] = tab1[i]
				i++
			} else {
				if tab1[i] <= tab2[j] {
					res[i+j] = tab1[i]
					i++
				} else {
					res[i+j] = tab2[j]
					j++
				}
			}
		}
	}
	return res
}

func splitTab(numbers []int) ([]int, []int) {
	return numbers[:len(numbers)/2], numbers[len(numbers)/2:]
}

func sortTab(numbers []int) []int {
	if len(numbers) < 2 {
		return numbers
	}
	var s1, s2 []int = splitTab(numbers)
	return mergeTab(sortTab(s1), sortTab(s2))
}

func parse(lines []string) ([]int, []int) {
	var s1 []int = make([]int, len(lines))
	var s2 []int = make([]int, len(lines))
	for i, line := range lines {
		var numbers = strings.Split(line, "   ")
		s1[i], _ = strconv.Atoi(numbers[0])
		s2[i], _ = strconv.Atoi(numbers[1])
	}
	return s1, s2
}

func occurrence(numbers []int) map[int]int {
	var res map[int]int = make(map[int]int)
	for _, number := range numbers {
		res[number]++
	}
	return res
}

func part1(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var s1, s2 []int = parse(lines)
	s1 = sortTab(s1)
	s2 = sortTab(s2)
	var res int
	for i := range s1 {
		res += int(math.Abs(float64(s1[i] - s2[i])))
	}
	return res
}

func part2(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var s1, s2 []int = parse(lines)
	var res int
	var occurrences map[int]int = occurrence(s2)
	for _, val := range s1 {
		res += val * occurrences[val]
	}
	return res
}

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("--2024 day 01 solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(string(input)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("part2 : ", part2(string(input)))
	fmt.Println(time.Since(start))
}
