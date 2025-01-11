package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"

	. "github.com/Discobluff/advent-of-code/go/utils/set"
)

//go:embed input.txt
var input string

func score(res rune) int {
	if 'a' <= res && res <= 'z' {
		return int(res-'a') + 1
	}
	return int(res-'A') + 27
}

func priorityBackpack(backpack string) int {
	var set = Intersect(buildItems(backpack[:len(backpack)/2]), buildItems(backpack[len(backpack)/2:]))
	var res rune = 'a'
	for r := range set {
		res = r
	}
	return score(res)
}

func buildItems(s string) Set[rune] {
	var res = DefSet[rune]()
	for _, b := range s {
		Add(res, b)
	}
	return res
}
func part1(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var res int
	for _, line := range lines {
		res += priorityBackpack(line)
	}
	return res
}

func part2(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var res int
	for i := 0; i < len(lines); i += 3 {
		var r rune = 'A'
		for a := range Intersect(Intersect(buildItems(lines[i]), buildItems(lines[i+1])), buildItems(lines[i+2])) {
			r = a
		}
		res += score(r)
	}
	return res
}

func main() {
	fmt.Println("--2022 day 03 solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(input))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("part2 : ", part2(input))
	fmt.Println(time.Since(start))
}
