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

func countDigits(number int) int {
	var res int = 1
	var prod int = 10
	for prod <= number {
		res++
		prod *= 10
	}
	return res
}

func concatenate(number1 int, number2 int) int {
	return int(math.Pow(10., float64(countDigits(number2))))*number1 + number2
}

func parse(line string) (int, []int) {
	var splitLine = strings.Split(line, ": ")
	var splitNumbers = strings.Split(splitLine[1], " ")
	var tabNumbers []int = make([]int, len(splitNumbers))
	for i, number := range splitNumbers {
		tabNumbers[i], _ = strconv.Atoi(number)
	}
	var res, _ = strconv.Atoi(splitLine[0])
	return res, tabNumbers
}

func solve(number int, numbers []int, index int, currentResult int, part2 bool) bool {
	if currentResult > number {
		return false
	}
	if index == len(numbers) {
		return number == currentResult
	}
	return solve(number, numbers, index+1, currentResult+numbers[index], part2) || solve(number, numbers, index+1, currentResult*numbers[index], part2) || (part2 && solve(number, numbers, index+1, concatenate(currentResult, numbers[index]), part2))
}

func parcours(lines []string, part2 bool) int {
	var res int
	for _, line := range lines {
		var number, numbers = parse(line)
		if solve(number, numbers, 0, 0, part2) {
			res += number
		}
	}
	return res
}

func part1(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	return parcours(lines, false)
}

func part2(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	return parcours(lines, true)
}

func main() {
	// var s = string(500)
	fmt.Println("--2024 day 07 solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(input))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("part2 : ", part2(input))
	fmt.Println(time.Since(start))
}
