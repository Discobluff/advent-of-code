package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
	"time"
)

func start(t1 string, t2 string) bool {
	if len(t1) > len(t2) {
		return false
	}
	for i := range t1 {
		if t1[i] != t2[i] {
			return false
		}
	}
	return true
}

func possible(towels []string, towel string, dict map[string]int) int {
	var result, ok = dict[towel]
	if ok {
		return result
	}
	var res int
	for _, t := range towels {
		if start(t, towel) {
			res += possible(towels, towel[len(t):], dict)
		}
	}
	dict[towel] = res
	return res

}

func part1(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var towels []string = strings.Split(lines[0], ", ")
	var res int
	var dict map[string]int = make(map[string]int)
	dict[""] = 1
	for _, towel := range lines[2:] {
		if possible(towels, towel, dict) != 0 {
			res++
		}
	}
	return res
}

// google code jam
// google kickstart
// competitiv programming
func part2(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var towels []string = strings.Split(lines[0], ", ")
	var res int
	var dict map[string]int = make(map[string]int)
	dict[""] = 1
	for _, towel := range lines[2:] {
		res += possible(towels, towel, dict)
	}
	return res
}

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("--2024 day 19 solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(string(input)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("part2 : ", part2(string(input)))
	fmt.Println(time.Since(start))
}
