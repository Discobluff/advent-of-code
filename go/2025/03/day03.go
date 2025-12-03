package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
	"time"
)

func max(line string, start int, end int)(int,int){
	var res int = int(line[start]-'0')
	var index int = start
	for i:=start+1;i<=end;i++{
		if int(line[i]-'0') > res{
			res = int(line[i])-'0'
			index = i
		}
	}
	return res,index
}

func maxVoltage(line string, size int)int{
	var res int
	var firstStart int
	for i:=range size{
		var digit, index = max(line,firstStart, len(line)-1-(size-1)+i)
		firstStart = index+1
		res = 10*res+digit
	}
	return res
}

func solve(input string, part2 bool)int{
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var res int
	var size int = 2
	if part2{
		size = 12
	}
	for _,line := range lines{
		res += maxVoltage(line, size)
	}
	return res
}

func part1(input string) int {
	return solve(input, false)
}

func part2(input string) int {
	return solve(input, true)
}

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("--2025 day 03 solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(string(input)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("part2 : ", part2(string(input)))
	fmt.Println(time.Since(start))
}
