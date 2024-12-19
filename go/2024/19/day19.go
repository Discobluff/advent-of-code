package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var input string

func start(t1 string, t2 string) bool {
	for i := range t1 {
		if i >= len(t2) || t1[i] != t2[i] {
			return false
		}
	}
	return true
}

func possible1(towels []string, towel string, dict map[string]bool) bool {
	var result, ok = dict[towel]
	if ok {
		return result
	}
	var res bool
	for _, t := range towels {
		if start(t, towel) {
			res = res || possible1(towels, towel[len(t):], dict)
		}
	}
	dict[towel] = res
	return res

}

func possible2(towels []string, towel string, dict map[string]int) int {
	var result, ok = dict[towel]
	if ok {
		return result
	}
	var res int
	for _, t := range towels {
		if start(t, towel) {
			res += possible2(towels, towel[len(t):], dict)
		}
	}
	dict[towel] = res
	return res

}

func part1(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var towels []string = strings.Split(lines[0], ", ")
	var res int
	var dict map[string]bool = make(map[string]bool)
	dict[""] = true
	for _, towel := range lines[2:] {
		if possible1(towels, towel, dict) {
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
		res += possible2(towels, towel, dict)
	}
	return res
}

func main() {
	fmt.Println("--2024 day 19 solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(input))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("part2 : ", part2(input))
	fmt.Println(time.Since(start))
}
