package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input string

func buildMap(dict *map[int][]int, line string) {
	var split []string = strings.Split(line, "|")
	var n1, _ = strconv.Atoi(split[0])
	var n2, _ = strconv.Atoi(split[1])
	if _, exists := (*dict)[n1]; !exists {
		(*dict)[n1] = make([]int, 0)
	}
	(*dict)[n1] = append((*dict)[n1], n2)
}

func buildUpdates(line string) []int {
	var res []int
	for _, number := range strings.Split(line, ",") {
		var numb, _ = strconv.Atoi(number)
		res = append(res, numb)
	}
	return res
}

func middleValue(slice []int) int {
	return slice[(len(slice)-1)/2]
}

func isCorrect(update []int, dict map[int][]int) bool {
	for i, val1 := range update {
		for j := i + 1; j < len(update); j++ {
			if !slices.Contains(dict[val1], update[j]) {
				return false
			}
		}
	}
	return true
}

func parse(lines []string) (map[int][]int, [][]int) {
	var dict map[int][]int = make(map[int][]int)
	var updates [][]int
	for _, line := range lines {
		if len(line) > 3 && line[2] == '|' {
			buildMap(&dict, line)
		} else {
			if line != "" {
				updates = append(updates, buildUpdates(line))
			}
		}
	}
	return dict, updates
}

func part1(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	dict, updates := parse(lines)
	var res int
	for _, update := range updates {
		if isCorrect(update, dict) {
			res += middleValue(update)
		}
	}
	return res
}

func reorder(update []int, dict map[int][]int) []int {
	var res []int = make([]int, len(update))
	var index int
	var indexs []int
	for index != len(res) {
		for i, val1 := range update {
			if !slices.Contains(indexs, i) {
				var right bool = true
				for j, val2 := range update {
					if j != i && !slices.Contains(indexs, j) && !slices.Contains(dict[val1], val2) {
						right = false
					}
				}
				if right {
					res[index] = val1
					indexs = append(indexs, i)
					index++
					break
				}
			}
		}
	}
	return res
}

func part2(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	dict, updates := parse(lines)
	var res int
	for _, update := range updates {
		if !isCorrect(update, dict) {
			res += middleValue(reorder(update, dict))
		}
	}
	return res
}

func main() {
	fmt.Println("--2024 day 05 solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(input))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("part2 : ", part2(input))
	fmt.Println(time.Since(start))
}
