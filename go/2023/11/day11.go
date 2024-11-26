package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed test.txt
var inputDay string

func part1(lines []string) int {
	return 0
}

func main() {
	var input = strings.TrimSuffix(inputDay, "\n")
	var lines = strings.Split(input, "\n")
	fmt.Println(part1(lines))
	// fmt.Println(part2(lines))

}
