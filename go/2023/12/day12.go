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
		res = 10*res + int(char-'0')
	}
	return res
}

func convertStringToIntSlice(strSlice []string) []int {
	var res []int = make([]int, len(strSlice))
	for i, str := range strSlice {
		res[i] = convertStringToInt(str)
	}
	return res
}

func initTab(line string) []int {
	var res []int = make([]int, len(line))
	for i, char := range line {
		if char == '?' {
			res[i] = -1
		}
		if char == '#' {
			res[i] = 1
		}
	}
	return res
}

func isVal(tab []int, start int, length int, val int) bool {
	for i := 0; i < length; i++ {
		if i+start >= len(tab) || tab[i+start] != val {
			return false
		}
	}
	return true
}

func valable(tab []int, numbers []int) bool {
	var number int
	for i := 0; i < len(tab); i++ {
		if tab[i] == 1 {
			if number >= len(numbers) {
				return false
			}
			if isVal(tab, i, numbers[number], 1) && (i+numbers[number] == len(tab) || tab[i+numbers[number]] == 0) {
				i = i + numbers[number]
				number++
			} else {
				return false
			}
		}
	}
	return number == len(numbers)
}

func changeValueSlice(tab []int, index int, value int) []int {
	tab[index] = value
	return tab
}

func copySlice(tab []int) []int {
	var res []int = make([]int, len(tab))
	for i, elem := range tab {
		res[i] = elem
	}
	return res
}

func browseAllConfig(tab []int, index int, numbers []int) int {
	if index == len(tab) {
		if valable(tab, numbers) {
			return 1
		}
		return 0
	}
	if tab[index] == -1 {
		return browseAllConfig(changeValueSlice(copySlice(tab), index, 0), index+1, numbers) + browseAllConfig(changeValueSlice(copySlice(tab), index, 1), index+1, numbers)
	}
	return browseAllConfig(tab, index+1, numbers)
}

func solve1(line string) int {
	var numbers []int = convertStringToIntSlice(strings.Split(strings.Split(line, " ")[1], ","))
	var tab []int = initTab(strings.Split(line, " ")[0])
	// fmt.Println(valable(tab, numbers), isVal(tab, 4, 3, 1))
	return browseAllConfig(tab, 0, numbers)
}

func part1(lines []string) int {
	var res int
	for _, line := range lines {
		res += solve1(line)
	}
	return res
}

func main() {
	var input = strings.TrimSuffix(inputDay, "\n")
	var lines = strings.Split(input, "\n")
	fmt.Println(part1(lines))
	// var numbers []int
	// numbers = append(numbers, 1)
	// var tab []int
	// tab = append(tab, 0)
	// tab = append(tab, 1)
	// tab = append(tab, 1)
	// tab = append(tab, 0)
	// fmt.Println(tab, numbers)
	// fmt.Println(valable(tab, numbers))

}
