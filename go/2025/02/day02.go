package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
	"time"
	"strconv"
)

func isInvalid(number int, times int)bool{
	var str = strconv.Itoa(number)
	var len = len(str)
	if len%times != 0{
		return false
	}
	for j:=range times-1{
		for i := range len/times{
			if str[i] != str[i+(j+1)*len/times]{
				return false
			}
		}
	}
	return true
}

func isInvalid1(number int)bool{
	return isInvalid(number,2)
}

func isInvalid2(number int)bool{
	for i:=2;i<=len(strconv.Itoa(number));i++{
		if isInvalid(number,i){
			return true
		}
	}
	return false
}

func solve(input string, validity func(int)bool) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var ranges = strings.Split(lines[0], ",")
	var res = 0
	for _,r := range ranges{
		var start, end int
		fmt.Sscanf(r, "%d-%d", &start, &end)
		for num := start;num<=end;num++{
			if validity(num){
				res += num
			}
		}
	}
	return res
}

func part1(input string) int {
	return solve(input, isInvalid1)
}

func part2(input string) int {
	return solve(input, isInvalid2)
}

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("--2025 day 02 solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(string(input)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("part2 : ", part2(string(input)))
	fmt.Println(time.Since(start))

}
