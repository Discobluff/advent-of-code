package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
	"time"
)

func part1(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	return len(lines)
}

func part2(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	return len(lines)
}

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("--%annee day %jour solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(string(input)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("part2 : ", part2(string(input)))
	fmt.Println(time.Since(start))
}
