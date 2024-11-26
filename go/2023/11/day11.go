package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"
)

//go:embed input.txt
var inputDay string

type Point struct {
	line, column int
}

func initPoint(line int, column int) Point {
	var res Point
	res.line = line
	res.column = column
	return res
}

func searchGalaxies(lines []string) []Point {
	var res []Point
	for i := range lines {
		for j, char := range lines[i] {
			if char == '#' {
				res = append(res, initPoint(i, j))
			}
		}
	}
	return res
}

func linesFromPoints(points []Point) []int {
	var res []int = make([]int, len(points))
	for i, point := range points {
		res[i] = point.line
	}
	return res
}

func columnsFromPoints(points []Point) []int {
	var res []int = make([]int, len(points))
	for i, point := range points {
		res[i] = point.column
	}
	return res
}

func isPresent(sliceInt []int, elem int) bool {
	for _, val := range sliceInt {
		if val == elem {
			return true
		}
	}
	return false
}

func linesColumnsWithoutGalaxies(linesInput []string) ([]int, []int) {
	var linesW []int
	var columnsW []int
	var galaxies []Point = searchGalaxies(linesInput)
	var lines, columns = linesFromPoints(galaxies), columnsFromPoints(galaxies)
	for l := range len(linesInput) {
		if !isPresent(lines, l) {
			linesW = append(linesW, l)
		}
	}
	for c := range len(linesInput[0]) {
		if !isPresent(columns, c) {
			columnsW = append(columnsW, c)
		}
	}
	return linesW, columnsW
}

func shortestDistance(point1 Point, point2 Point, linesW []int, columnsW []int, add int) int {
	var distance int = int(math.Abs(float64(point1.line-point2.line)) + math.Abs(float64(point1.column-point2.column)))
	for line := min(point1.line, point2.line) + 1; line <= max(point1.line, point2.line)-1; line++ {
		if isPresent(linesW, line) {
			distance += add
		}
	}
	for column := min(point1.column, point2.column) + 1; column <= max(point1.column, point2.column)-1; column++ {
		if isPresent(columnsW, column) {
			distance += add
		}
	}
	return distance
}

func part1(lines []string) int {
	var galaxies []Point = searchGalaxies(lines)
	var linesWithoutGalaxies, columnsWithoutGalaxies = linesColumnsWithoutGalaxies(lines)
	var res int
	for i := range len(galaxies) - 1 {
		for j := i + 1; j < len(galaxies); j++ {
			res += shortestDistance(galaxies[i], galaxies[j], linesWithoutGalaxies, columnsWithoutGalaxies, 1)
		}
	}
	return res
}

func part2(lines []string) int {
	var galaxies []Point = searchGalaxies(lines)
	var linesWithoutGalaxies, columnsWithoutGalaxies = linesColumnsWithoutGalaxies(lines)
	var res int
	for i := range len(galaxies) - 1 {
		for j := i + 1; j < len(galaxies); j++ {
			res += shortestDistance(galaxies[i], galaxies[j], linesWithoutGalaxies, columnsWithoutGalaxies, 1000000-1)
		}
	}
	return res
}

func main() {
	var input = strings.TrimSuffix(inputDay, "\n")
	var lines = strings.Split(input, "\n")
	fmt.Println(part1(lines))
	fmt.Println(part2(lines))

}
