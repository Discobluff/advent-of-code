package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
	"time"
)

func parcours(line string) int {
	var sum int
	for i := 0; i < len(line)-4; i++ {
		if line[i:i+4] == "mul(" {
			var number1 int
			var number2 int
			for j := i + 4; j < len(line) && isDigit(line[j]); j++ {
				number1 = 10*number1 + int(line[j]-'0')
				i = j
			}
			if i+1 < len(line) && line[i+1] == ',' {
				for j := i + 2; j < len(line) && isDigit(line[j]); j++ {
					number2 = 10*number2 + int(line[j]-'0')
					i = j
				}
				if i+1 < len(line) && line[i+1] == ')' {
					sum += number1 * number2
				}
			}

		}
	}
	return sum
}

func parcours2(line string, enab bool) (int, bool) {
	var sum int
	var enable bool = enab
	for i := 0; i < len(line); i++ {
		if i+4 < len(line) && line[i:i+4] == "mul(" {
			if enable {
				var number1 int
				for j := i + 4; j < len(line) && isDigit(line[j]); j++ {
					number1 = 10*number1 + int(line[j]-'0')
					i = j
				}
				if i+1 < len(line) && line[i+1] == ',' {
					var number2 int
					for j := i + 2; j < len(line) && isDigit(line[j]); j++ {
						number2 = 10*number2 + int(line[j]-'0')
						i = j
					}
					if i+1 < len(line) && line[i+1] == ')' {
						sum += number1 * number2
						i = i + 1
					}
				}
			} else {
				i += 3
			}
		}
		if i+4 < len(line) && line[i:i+4] == "do()" {
			enable = true
			i += 3
		}
		if i+7 < len(line) && line[i:i+7] == "don't()" {
			enable = false
			i += 6
		}
	}
	return sum, enable
}

func isDigit(char byte) bool {
	return char >= 48 && char <= 58
}

func test(line string) {
	var v int
	var w int
	fmt.Sscanf(line, "mul(%d,%d)", &v, &w)
	fmt.Println(v)
}

func part1(lines []string) int {
	var res int
	for _, line := range lines {
		res += parcours(line)
	}
	return res
}

func part2(lines []string) int {
	var res int
	var enable bool = true
	for _, line := range lines {
		sum, enab := parcours2(line, enable)
		res += sum
		enable = enab
	}
	return res
}

func main() {
	input1, _ := os.ReadFile("input.txt")
	var input = strings.TrimSuffix(string(input1), "\n")
	var lines = strings.Split(input, "\n")
	fmt.Println("--2024 day 03 solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(lines))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("part2 : ", part2(lines))
	fmt.Println(time.Since(start))
	test(lines[0])
}
