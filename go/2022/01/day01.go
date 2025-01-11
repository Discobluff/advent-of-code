package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"

	. "github.com/Discobluff/advent-of-code/go/utils/set"
)

//go:embed input.txt
var input string

func maxInt(a int, b int) bool {
	return a <= b
}

func solve(input string, first int) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n\n")
	var calories = DefSet[int]()
	for _, elve := range lines {
		var calorieTotal int
		for _, calorie := range strings.Split(elve, "\n") {
			var conv, _ = strconv.Atoi(calorie)
			calorieTotal += conv
		}
		Add(calories, calorieTotal)
	}
	var res int
	for range first {
		var t, _ = Max(calories, maxInt)
		res += t
		Remove(calories, t)
	}
	return res

}

func part1(input string) int {
	return solve(input, 1)
}

func part2(input string) int {
	return solve(input, 3)
}

func main() {
	fmt.Println("--2022 day 01 solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(input))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("part2 : ", part2(input))
	fmt.Println(time.Since(start))
}
