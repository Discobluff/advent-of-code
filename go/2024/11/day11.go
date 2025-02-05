package main

import (
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func parse(line string) map[int]int {
	var numbers []string = strings.Split(line, " ")
	var res map[int]int = make(map[int]int)
	for _, number := range numbers {
		var num, _ = strconv.Atoi(number)
		res[num] += 1
	}
	return res
}

func countDigits(number int) (int, bool) {
	var digits int = 1
	var ok bool
	var limit int = 10
	for number >= limit {
		digits++
		limit *= 10
		ok = !ok
	}
	return digits, ok
}

func split(number int, nbDigits int) (int, int) {
	var n2 int
	var length int = 1
	for range nbDigits / 2 {
		var lastDigit = number - 10*(number/10)
		n2 += length * lastDigit
		length *= 10
		number = (number - lastDigit) / 10
	}
	return number, n2

}

func blink(val int) (int, int, bool) {
	if val == 0 {
		return 1, 0, false
	}
	var nbDigits, ok = countDigits(val)
	if ok {
		var n1, n2 = split(val, nbDigits)
		return n1, n2, true
	}
	return val * 2024, 0, false
}

func sumMap(dict map[int]int) int {
	var res int
	for _, val := range dict {
		res += val
	}
	return res
}

func solve(input string, loop int) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var numbers map[int]int = parse(lines[0])
	for range loop {
		var newMap map[int]int = make(map[int]int)
		for stone, number := range numbers {
			var stone1, stone2, ok = blink(stone)
			newMap[stone1] += number
			if ok {
				newMap[stone2] += number
			}
		}
		numbers = newMap
	}
	return sumMap(numbers)
}

func part1(input string) int {
	return solve(input, 25)
}

func part2(input string) int {
	return solve(input, 75)
}

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("--2024 day 11 solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(string(input)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("part2 : ", part2(string(input)))
	fmt.Println(time.Since(start))
}
