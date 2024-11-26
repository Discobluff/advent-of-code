package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var inputDay string

func convertStringToInt(str string) int {
	var res int
	for _, char := range str {
		res = res*10 + int(char-'0')
	}
	return res
}

func convertSliceStringToSliceInt(strSlice []string) []int {
	var res []int
	for _, number := range strSlice {
		res = append(res, convertStringToInt(number))
	}
	return res
}

func getSeeds(str string) []int {
	return convertSliceStringToSliceInt(strings.Split(strings.Split(str, ": ")[1], " "))
}

func getValues(str string) []int {
	return convertSliceStringToSliceInt(strings.Split(str, " "))
}

func buildRange(start int, len int) []int {
	var res []int = make([]int, len, len)
	for i := start; i-start < len; i++ {
		res[i-start] = i
	}
	return res
}

func buildRanges(tab []int) ([]int, []int) {
	return buildRange(tab[1], tab[2]), buildRange(tab[0], tab[2])
}

func addToDict(str string, dict *map[int]int) {
	tab1, tab2 := buildRanges(getValues(str))
	for i := range len(tab1) {
		(*dict)[tab1[i]] = tab2[i]
	}
}

func minSlice(slice []int) int {
	var res int = -1
	for _, elem := range slice {
		if res == -1 || elem < res {
			res = elem
		}
	}
	return res
}

func isPresent(char byte, str string) bool {
	for i := range str {
		if str[i] == char {
			return true
		}
	}
	return false
}

func parse(lines []string) []map[int]int {
	var index int = -1
	var dictTab []map[int]int
	for i := 1; i < len(lines); i++ {
		if lines[i] == "" {
			index += 1
			dictTab = append(dictTab, make(map[int]int))
		} else {
			if !isPresent(':', lines[i]) {
				addToDict(lines[i], &dictTab[index])
			}
		}
	}
	return dictTab
}

func part1(lines []string) int {
	var seeds []int = getSeeds(lines[0])
	var dictTab = parse(lines)
	var locationSeeds []int = make([]int, len(seeds), len(seeds))
	return 0
	for i, seed := range seeds {
		var location int = seed
		for _, dict := range dictTab {
			value, exists := dict[location]
			if exists {
				location = value
			}
		}
		locationSeeds[i] = location
	}
	return minSlice(locationSeeds)
}

func main() {
	var input = strings.TrimSuffix(inputDay, "\n")
	var lines = strings.Split(input, "\n")
	fmt.Println(part1(lines))
}
