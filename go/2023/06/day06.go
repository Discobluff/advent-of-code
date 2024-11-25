package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var inputDay string

func getDistance(time int, timeAccel int) int {
	return (time - timeAccel) * timeAccel
}

func getCountGreaterDistance(time int, record int) int {
	var count int
	for i := range time + 1 {
		if getDistance(time, i) > record {
			count += 1
		}
	}
	return count
}

func convertStringToInt(str string) int {
	var res int
	for _, char := range str {
		res = res*10 + int(char-48)
	}
	return res
}
func parseLine(line string) []int {
	var res []int
	for _, number := range strings.Split(line, " ")[1:] {
		if number != "" {
			res = append(res, convertStringToInt(number))
		}
	}
	return res
}

func part1(lines []string) int {
	var time []int = parseLine(lines[0])
	var distance []int = parseLine(lines[1])
	fmt.Println(time, distance)
	var product int = 1
	for i := range len(time) {
		product *= getCountGreaterDistance(time[i], distance[i])
	}
	return product
}

func main() {
	var input = strings.TrimSuffix(inputDay, "\n")
	var lines = strings.Split(input, "\n")
	fmt.Println(part1(lines))
}
