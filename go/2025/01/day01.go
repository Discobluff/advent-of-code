package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
	"time"
)

func mod(a int, b int) int{
	return ((a%b)+b)%b
}

func part1(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var dial = 50
	var res = 0
	
	for _,line := range lines{
		var val = 0
		for i := 1; i < len(line); i++ {
			val = 10*val + int(line[i]-'0')
		}
		if line[0] == 'L' {
			dial += 100 - val
		} else {
			dial += val
		}
		dial = dial%100
		if dial == 0{
			res++
		}
	}
	return res
}

func part2(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var dial = 50
	var res = 0
	var lastDial int
	for _,line := range lines{
		val := 0
		lastDial = dial
		for i := 1; i < len(line); i++ {
			val = 10*val + int(line[i]-'0')
		}
		if line[0] == 'L'{
			dial = dial-val
		}
		if line[0] == 'R'{
			dial = dial+val
		}
		if (dial>0){
			res += dial/100
		}
		if (dial<0){
			if lastDial != 0{
				res += 1
			}
			res += -dial/100

		}
		if dial == 0{
			res++
		}
		dial = mod(dial,100)
	}
	return res
}

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("--2025 day 01 solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(string(input)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("part2 : ", part2(string(input)))
	fmt.Println(time.Since(start))
}
