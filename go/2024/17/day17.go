package main

import (
	_ "embed"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func combo(registers []int, val int) int {
	if val >= 0 && val <= 3 {
		return val
	}
	if val >= 4 && val <= 6 {
		return registers[val-4]
	}
	return 7

}

func instructions(index *int, i1 int, i2 int, registers *[]int) int {
	if i1 == 0 {
		var res = (*registers)[0] / int(math.Pow(2., float64(combo(*registers, i2))))
		(*registers)[0] = res
	}
	if i1 == 1 {
		var res = (*registers)[1] ^ i2
		(*registers)[1] = res
	}
	if i1 == 2 {
		var res = combo(*registers, i2) % 8
		(*registers)[1] = res
	}
	if i1 == 3 {
		if (*registers)[0] != 0 {
			(*index) = i2
			return -1
		}
	}
	if i1 == 4 {
		var res = (*registers)[1] ^ (*registers)[2]
		(*registers)[1] = res
	}
	if i1 == 5 {
		var res = combo(*registers, i2) % 8
		(*index) = (*index) + 2
		return res
	}
	if i1 == 6 {
		var res = (*registers)[0] / int(math.Pow(2., float64(combo(*registers, i2))))
		(*registers)[1] = res
	}
	if i1 == 7 {
		var res = (*registers)[0] / int(math.Pow(2., float64(combo(*registers, i2))))
		(*registers)[2] = res
	}
	(*index) = (*index) + 2
	return -1

}

func parseRegister(line string) int {
	var res, _ = strconv.Atoi(strings.Split(line, ": ")[1])
	return res
}

func parseInstructions(line string) []int {
	var instructions = strings.Split(strings.Split(line, ": ")[1], ",")
	var res []int = make([]int, len(instructions))
	for i, number := range instructions {
		var numb, _ = strconv.Atoi(number)
		res[i] = numb
	}
	return res
}

func part1(input string) string {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var registers []int = []int{parseRegister(lines[0]), parseRegister(lines[1]), parseRegister(lines[2])}
	var instructionsSlice []int = parseInstructions(lines[4])
	var index int
	var res []byte
	for index < len(instructionsSlice)-1 {
		var val = instructions(&index, instructionsSlice[index], instructionsSlice[index+1], &registers)
		if val != -1 {
			res = append(res, byte(val+48))
			res = append(res, ',')
		}
	}
	return string(res[:len(res)-1])
}

func part2(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var instructionsSlice []int = parseInstructions(lines[4])
	// For the input.txt
	//The instructions are repeated from the start if A!=0 when we reach the end of the instructions
	// At each repetition, we write only one time in the output
	//Since there are len(instructionsSlice) instructions we want len(instructionsSlice) repetitions
	// At each repetition, A is changed only one time and become A/8
	// So we can write A=a_0+8*a_1+...+a_15*8^15
	// So we need to find those a_i in [|0,7|]
	// We start from the end
	// Try the next digit of A, and check if it is the right answer which is wrote in the output
	var aPossibles []int = []int{4}
	for i := range len(instructionsSlice) - 1 {
		var nexts []int
		for _, newA := range aPossibles {
			var b int
			var c int
			var a int
			for a_i := range 8 {
				a = a_i + newA*8
				b = a_i
				b = b ^ 1
				c = a / int(math.Pow(2., float64(b)))
				b = b ^ 5
				b = b ^ c
				// a = a / 8
				if b%8 == instructionsSlice[len(instructionsSlice)-i-2] {
					nexts = append(nexts, a)
				}
			}
		}
		aPossibles = nexts
	}
	return slices.Min(aPossibles)
}

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("--2024 day 17 solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(string(input)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("part2 : ", part2(string(input)))
	fmt.Println(time.Since(start))
}
