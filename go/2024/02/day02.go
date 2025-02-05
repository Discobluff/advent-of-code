package main

import (
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func parse(lines []string) []int {
	var res []int = make([]int, len(lines))
	for i, line := range lines {
		res[i], _ = strconv.Atoi(line)
	}
	return res
}

func signe(x int) int {
	if x >= 0 {
		return 1
	}
	return -1
}

func isSafe(numbers []int, ignore int) bool {
	var sens int
	for i := range len(numbers) - 1 {
		if i != ignore && !(ignore == len(numbers)-1 && i == ignore-1) {
			var current int = i
			var next int = i + 1
			if i == ignore-1 {
				next = ignore + 1
			}
			var diff int = numbers[current] - numbers[next]
			if sens == 0 {
				sens = signe(diff)
			}
			if signe(diff) != sens {
				return false
			}
			diff = sens * diff
			if diff < 1 || diff > 3 {
				return false
			}
		}
	}
	return true
}
func isSafe2(numbers []int) bool {
	for i := -1; i < len(numbers); i++ {
		if isSafe(numbers, i) {
			return true
		}
	}
	return false
}

func differences(numbers []int) []int {
	var res []int = make([]int, len(numbers)-1)
	for i := range len(numbers) - 1 {
		res[i] = res[i+1] - res[i]
	}
	return res
}

func part1(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var res int
	for _, line := range lines {
		if isSafe(parse(strings.Split(line, " ")), -1) {
			res++
		}
	}
	return res
}

func part2(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var res int
	for _, line := range lines {
		if isSafe2(parse(strings.Split(line, " "))) {
			res++
		}
	}
	return res
}

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("--2024 day 02 solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(string(input)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("part2 : ", part2(string(input)))
	fmt.Println(time.Since(start))

}
