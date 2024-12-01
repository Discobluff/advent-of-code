package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input string

func minSlice(numbers []int, ignoredIndex []bool) (int, int) {
	var res int = -1
	var index int = -1
	for i, val := range numbers {
		if !ignoredIndex[i] {
			if res == -1 || val < res {
				res = val
				index = i
			}
		}
	}
	return res, index

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

func occurrence(numbers []int, number int) int {
	var res int
	for _, val := range numbers {
		if val == number {
			res += 1
		}
	}
	return res
}

func part1(lines []string) int {
	var s1, s2 []int = parse(lines)
	var ignoredIndex1 []bool = make([]bool, len(s1))
	var ignoredIndex2 []bool = make([]bool, len(s1))
	var res int
	for range s1 {
		var min1, i1 = minSlice(s1, ignoredIndex1)
		var min2, i2 = minSlice(s2, ignoredIndex2)
		ignoredIndex1[i1] = true
		ignoredIndex2[i2] = true
		res += int(math.Abs(float64(min1 - min2)))
	}
	return res
}

func part2(lines []string) int {
	var s1, s2 []int = parse(lines)
	var res int
	for _, val := range s1 {
		res += val * occurrence(s2, val)
	}
	return res
}

func main() {
	var input = strings.TrimSuffix(input, "\n")
	var lines = strings.Split(input, "\n")
	fmt.Println("--2024 day 1 solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(lines))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("part2 : ", part2(lines))
	fmt.Println(time.Since(start))
}
