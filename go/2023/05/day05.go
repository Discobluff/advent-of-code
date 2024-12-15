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

func parse(lines []string) [][][3]int {
	var index int = -1
	var res [][][3]int = make([][][3]int, 0)
	for i := 1; i < len(lines); i++ {
		if lines[i] == "" {
			index += 1
			var newTab [][3]int
			res = append(res, newTab)
		} else {
			if !isPresent(':', lines[i]) {
				res[index] = append(res[index], [3]int(getValues(lines[i])))
			}
		}
	}
	return res
}

func calcul(mapRange [3]int, val int) (int, bool) {
	if mapRange[1] <= val && val < mapRange[1]+mapRange[2] {
		return mapRange[0] + val - mapRange[1], true
	}
	return val, false
}

func calculSlice(tabMap [][3]int, val int) int {
	for _, mapRange := range tabMap {
		var newVal, ok = calcul(mapRange, val)
		if ok {
			return newVal
		}
	}
	return val
}

func part1(lines []string) int {
	var seeds []int = getSeeds(lines[0])
	var tabRanges [][][3]int = parse(lines)
	for _, tabRange := range tabRanges {
		for j, seed := range seeds {
			seeds[j] = calculSlice(tabRange, seed)
		}
	}
	return minSlice(seeds)
}

func main() {
	var input = strings.TrimSuffix(inputDay, "\n")
	var lines = strings.Split(input, "\n")
	fmt.Println(part1(lines))
}
