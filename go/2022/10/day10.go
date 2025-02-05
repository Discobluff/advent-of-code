package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
	"time"
)

func part1(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var res int
	var register int = 1
	var cycle int = 1
	for _, line := range lines {
		if cycle%40 == 20 {
			res += register * cycle
		}
		cycle += 1
		if line != "noop" {
			var d int
			fmt.Sscanf(line, "addx %d", &d)
			if cycle%40 == 20 {
				res += register * cycle
			}
			register += d
			cycle += 1
		}
	}
	return res
}

func part2(input string) string {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var grid [][]byte = make([][]byte, 6)
	for i := range 6 {
		grid[i] = make([]byte, 40)
	}
	var register int = 1
	var cycle int
	for _, line := range lines {
		if register-1 <= cycle%40 && cycle%40 <= register+1 {
			grid[cycle/40][cycle%40] = '#'
		} else {
			grid[cycle/40][cycle%40] = ' '
		}
		cycle += 1
		if line != "noop" {
			var add int
			fmt.Sscanf(line, "addx %d", &add)
			if register-1 <= cycle%40 && cycle%40 <= register+1 {
				grid[cycle/40][cycle%40] = '#'
			} else {
				grid[cycle/40][cycle%40] = ' '
			}
			cycle += 1
			register += add

		}
	}
	var res string
	for _, line := range grid {
		res += string(line) + "\n"
	}
	return res
}

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("--2024 day 05 solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(string(input)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("part2 : ")
	fmt.Println(part2(string(input)))
	fmt.Println(time.Since(start))
}
