package main

import (
	_ "embed"
	"fmt"
	"strings"
)

type couple [2]int

//go:embed input.txt
var inputDay string

func part1(lines []string) int {
	var sum int
	for _, line := range lines {
		var find bool
		var res couple
		for _, char := range line {
			if char >= 48 && char <= 57 {
				if !find {
					find = true
					res[0] = int(char - 48)
				}
				res[1] = int(char - 48)
			}
		}
		sum += 10*res[0] + res[1]
	}
	return sum
}

func part2(lines []string) int {
	var sum int
	numbers := [10]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for _, line := range lines {
		var find bool
		var res couple
		for i := 0; i < len(line); i++ {
			if line[i] >= 48 && line[i] <= 57 {
				if !find {
					find = true
					res[0] = int(line[i] - 48)
				}
				res[1] = int(line[i] - 48)
			}
			for j, number := range numbers {
				if i+len(number) <= len(line) {
				}
				if i+len(number) <= len(line) && line[i:i+len(number)] == number {
					if !find {
						find = true
						res[0] = j
					}
					res[1] = j
					break
				}
			}
		}
		sum += 10*res[0] + res[1]
	}
	return sum
}

func main() {
	var input = strings.TrimSuffix(inputDay, "\n")
	var lines = strings.Split(input, "\n")
	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}
