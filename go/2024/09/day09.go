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

type Block struct {
	start, length int
}

func sum(line string) int {
	var res int
	for i := range line {
		temp, _ := strconv.Atoi(string(line[i]))
		res += temp
	}
	return res
}

func initTab(line string) []int {
	var res []int
	for i, char := range line {
		var val int
		if i%2 == 1 {
			val = -1
		} else {
			val = i / 2
		}
		var count, _ = strconv.Atoi(string(char))
		for range count {
			res = append(res, val)
		}
	}
	return res
}

func checkSum(slice []int) int {
	var res int
	for i, val := range slice {
		if val != -1 {
			res += i * val
		}
	}
	return res
}

func solve1(line string) int {
	var tab []int = initTab(line)
	var end int = len(tab) - 1
	var start int
	var res int
	for start < end {
		for tab[start] != -1 {
			res += start * tab[start]
			start++
		}
		for tab[end] == -1 {
			end--
		}
		if start < end {

			tab[start] = tab[end]
			tab[end] = -1
		}
	}
	return res
}

func maxTab(slice []int) (int, int, int) {
	var res int = -1
	var deb int = -1
	var fin int = -1
	for i, val := range slice {
		if res == val {
			fin = i
		}
		if res == -1 || res < val {
			res = val
			deb = i
			fin = i
		}
	}
	return res, deb, fin
}

func rangeSlice(slice []int, val int) (int, int) {
	var start int = -1
	var end int
	for i, v := range slice {
		if start == -1 && v == val {
			start = i
		}
		if v == val {
			end = i
		}
	}
	return start, end
}

func isEmpty(slice []int) bool {
	for _, val := range slice {
		if val != -1 {
			return false
		}
	}
	return true
}

func solve2(line string) int {
	var tab []int = initTab(line)
	var index, start, end int = maxTab(tab)
	for index > 0 {
		for i := range start {
			if isEmpty(tab[i : i+end-start+1]) {
				for j := i; j < i+end-start+1; j++ {
					tab[j] = tab[start+j-i]
					tab[start+j-i] = -1
				}
				break
			}
		}
		index--
		start, end = rangeSlice(tab, index)
	}
	return checkSum(tab)
}

func part1(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	return solve1(lines[0])
}

func part2(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	return solve2(lines[0])
}
func main() {
	fmt.Println("--2024 day 09 solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(input))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("part2 : ", part2(input))
	fmt.Println(time.Since(start))
}
