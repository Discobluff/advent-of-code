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

func mix(value int, secret int) int {
	return value ^ secret
}

func prune(secret int) int {
	return secret % 16777216
}

func nextSecretNumber(secretNumber int) int {
	secretNumber = prune(mix(secretNumber*64, secretNumber))
	secretNumber = prune(mix(secretNumber/32, secretNumber))
	secretNumber = prune(mix(secretNumber*2048, secretNumber))
	return secretNumber
}

func secretNumbers(start int, count int) []int {
	var res []int = make([]int, count)
	for i := range count {
		res[i] = start
		start = nextSecretNumber(start)
	}
	return res
}

func part1(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var res int
	for _, line := range lines {
		var secret, _ = strconv.Atoi(line)
		var numbers = secretNumbers(secret, 2001)
		res += numbers[len(numbers)-1]
	}
	return res
}

func differenceDigitNumbers(numbers []int) []int {
	var res []int = make([]int, len(numbers)-1)
	for i := range len(numbers) - 1 {
		res[i] = numbers[i+1]%10 - numbers[i]%10
	}
	return res
}

func part2(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var numbersSlice [][]int = make([][]int, len(lines))
	var differences [][]int = make([][]int, len(lines))
	for i, line := range lines {
		var secret, _ = strconv.Atoi(line)
		numbersSlice[i] = secretNumbers(secret, 2000)
		differences[i] = differenceDigitNumbers(numbersSlice[i])
	}
	var maxi int
	for i := -9; i < 10; i++ {
		for j := -9; j < 10; j++ {
			for k := -9; k < 10; k++ {
				for l := -9; l < 10; l++ {
					var suite []int = []int{i, j, k, l}
					var val int
					for n, difference := range differences {
						for m := range len(difference) - 4 {
							if slices.Equal(suite, difference[m:m+4]) {
								val += numbersSlice[n][m+4] % 10
								break
							}
						}
					}
					if val > maxi {
						maxi = val
					}
				}
			}
		}
	}
	return maxi
}

func main() {
	fmt.Println("--2024 day 22 solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(input))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("part2 : ", part2(input))
	fmt.Println(time.Since(start))
}
