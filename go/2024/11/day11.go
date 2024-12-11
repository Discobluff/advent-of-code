package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//go:embed test.txt
var input string

func parse(line string) []int {
	var numbers []string = strings.Split(line, " ")
	var res []int = make([]int, len(numbers))
	for i, number := range numbers {
		res[i], _ = strconv.Atoi(number)
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

func blinkNumber(number int, times int) int {
	// fmt.Println(number, times)
	if times == 0 {
		return 1
	}
	if number == 0 {
		return blinkNumber(1, times-1)
	}
	var digits, ok = countDigits(number)
	if ok {
		var n1, n2 int = split(number, digits)
		return blinkNumber(n1, times-1) + blinkNumber(n2, times-1)
	}
	return blinkNumber(number*2024, times-1)
}

// func blink(numbers []int) []int {
// 	var res []int
// 	for _, val := range numbers {
// 		if val == 0 {
// 			res = append(res, 1)
// 		} else {
// 			var nbDigits, ok = countDigits(val)
// 			if ok {
// 				var n1, n2 = split(val, nbDigits)
// 				res = append(res, n1)
// 				res = append(res, n2)
// 			} else {
// 				res = append(res, val*2024)
// 			}
// 		}
// 	}
// 	return res
// }

func solve(input string, loop int) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var numbers []int = parse(lines[0])
	var res int
	for _, number := range numbers {
		res += blinkNumber(number, loop)
	}
	return res
}

func part1(input string) int {
	return solve(input, 25)
}

func part2(input string) int {
	return solve(input, 75)
}

func main() {
	fmt.Println("--2024 day 11 solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(input))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("part2 : ", part2(input))
	fmt.Println(time.Since(start))
}
