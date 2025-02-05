package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
	"time"
)

var rock byte = 'A'
var paper byte = 'B'
var scissor byte = 'C'
var moves1 = map[byte]int{rock: 1, paper: 2, scissor: 3}
var moves2 = map[byte]byte{'X': rock, 'Y': paper, 'Z': scissor}
var duel = map[byte]byte{rock: scissor, scissor: paper, paper: rock}

var part2End = map[byte]int{'X': 0, 'Y': 3, 'Z': 6}

func win(m1 byte, m2 byte) int {
	if m1 == m2 {
		return 3
	}
	if m2 == duel[m1] {
		return 0
	}
	return 6
}

func first(s string) byte {
	return s[0]
}

func part1(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var res int
	for _, line := range lines {
		var split = strings.Split(line, " ")
		var m2 = moves2[first(split[1])]
		res += moves1[m2] + win(first(split[0]), m2)
	}
	return res
}

func part2(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var res int
	for _, line := range lines {
		var split = strings.Split(line, " ")
		var m2 = part2End[first(split[1])]
		for move, score := range moves1 {
			if m2 == win(first(split[0]), move) {
				res += score + m2
				continue
			}
		}
	}
	return res
}

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("--2022 day 02 solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(string(input)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("part2 : ", part2(string(input)))
	fmt.Println(time.Since(start))
}
