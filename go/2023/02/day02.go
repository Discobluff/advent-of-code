package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

type triplet [3]int

//go:embed input.txt
var inputDay string

func part1(informations [][]triplet) int {
	var sum int
	for i, row := range informations {
		add := true
		for _, info := range row {
			if info[0] > 12 || info[1] > 13 || info[2] > 14 {
				add = false
			}
		}
		if add {
			sum += i + 1
		}
	}
	return sum
}

func part2(informations [][]triplet) int {
	var sum int
	for _, row := range informations {
		var minColor triplet
		for _, info := range row {
			for i := range 3 {
				if info[i] > minColor[i] {
					minColor[i] = info[i]
				}
			}
		}
		sum += minColor[0] * minColor[1] * minColor[2]
	}
	return sum
}

func main() {
	var input = strings.TrimSuffix(inputDay, "\n")
	var lines = strings.Split(input, "\n")
	var nombreManches int = len(lines)
	var informations = make([][]triplet, 0, nombreManches)
	for _, line := range lines {
		line = strings.Split(line, ":")[1]
		var infoMancheI = make([]triplet, 0, 1)
		for _, row := range strings.Split(line, ";") {
			var couleurs triplet
			for _, info := range strings.Split(row, ",") {
				temp := strings.Split(info, " ")
				if temp[2] == "red" {
					couleurs[0], _ = strconv.Atoi(temp[1])
				}
				if temp[2] == "green" {
					couleurs[1], _ = strconv.Atoi(temp[1])
				}
				if temp[2] == "blue" {
					couleurs[2], _ = strconv.Atoi(temp[1])
				}
			}
			infoMancheI = append(infoMancheI, couleurs)
		}
		informations = append(informations, infoMancheI)
	}
	fmt.Println(informations)

	fmt.Println(part1(informations))
	fmt.Println(part2(informations))
}
