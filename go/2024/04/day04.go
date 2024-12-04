package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var input string

func searchLine(line string) int {
	var res int
	for j := 0; j < len(line); j++ {
		if j+4 < len(line) {
		}
		if j+3 < len(line) && (line[j:j+4] == "XMAS" || line[j:j+4] == "SAMX") {
			res++
		}
	}
	return res
}

func buildColumn(lines []string, columnIndex int) string {
	var res []byte = make([]byte, len(lines))
	for i := range lines {
		res[i] = lines[i][columnIndex]
	}
	return string(res)
}

func buildDiagonal(lines []string, line int, column int, size int, sens int) string {
	var res []byte = make([]byte, 0)
	for k := 0; k < size && line+k < len(lines) && column+sens*k >= 0; k++ {
		res = append(res, lines[line+k][column+sens*k])
	}
	return string(res)
}

func part1(lines []string) int {
	var res int
	//Search line
	for _, line := range lines {
		res += searchLine(line)
	}
	//Search column
	for j := range lines[0] {
		res += searchLine(buildColumn(lines, j))
	}
	//Search right diagonal
	for i := range lines {
		for j := range lines[i] {
			leftDiag := buildDiagonal(lines, i, j, 4, -1)
			rightDiag := buildDiagonal(lines, i, j, 4, -1)
			if leftDiag == "XMAS" || leftDiag == "SAMX" {
				res++
			}
			if rightDiag == "XMAS" || rightDiag == "SAMX" {
				res++
			}
		}
	}
	return res
}

func part2(lines []string) int {
	var res int
	for i := 1; i < len(lines)-1; i++ {
		for j := 1; j < len(lines[i])-1; j++ {
			if lines[i][j] == 'A' {
				if ((lines[i-1][j-1] == 'M' && lines[i+1][j+1] == 'S') || (lines[i-1][j-1] == 'S' && lines[i+1][j+1] == 'M')) && ((lines[i+1][j-1] == 'M' && lines[i-1][j+1] == 'S') || (lines[i+1][j-1] == 'S' && lines[i-1][j+1] == 'M')) {
					res++
				}
			}
		}
	}
	return res
}

func main() {
	var input = strings.TrimSuffix(input, "\n")
	var lines = strings.Split(input, "\n")
	fmt.Println("--2024 day 03 solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(lines))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("part2 : ", part2(lines))
	fmt.Println(time.Since(start))
}
