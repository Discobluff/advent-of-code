package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input string

type Interval struct {
	a, b int
}

func overlap(i1, i2 Interval) bool {
	return i1.a <= i2.a && i1.b >= i2.a
}

func include(i1, i2 Interval) bool {
	return i1.a >= i2.a && i1.b <= i2.b
}

func createInverval(s string) Interval {
	var split = strings.Split(s, "-")
	var n1, _ = strconv.Atoi(split[0])
	var n2, _ = strconv.Atoi(split[1])
	return Interval{a: n1, b: n2}
}

func solve(input string, fun func(Interval, Interval) bool) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var res int
	for _, line := range lines {
		var split = strings.Split(line, ",")
		var i1 = createInverval(split[0])
		var i2 = createInverval(split[1])
		if fun(i1, i2) || fun(i2, i1) {
			res++
		}
	}
	return res
}

func part1(input string) int {
	return (solve(input, include))
}

func part2(input string) int {
	return (solve(input, overlap))
}

func main() {
	fmt.Println("--2022 day 04 solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(input))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("part2 : ", part2(input))
	fmt.Println(time.Since(start))
}
