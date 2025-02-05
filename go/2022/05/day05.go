package main

import (
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Move struct {
	count, start, end int
}

type Stack []byte

func add(s *Stack, e byte) {
	(*s) = append((*s), e)
}

func pop(s *Stack) byte {
	var res = (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return res
}

func parse(input string) ([]Stack, []Move) {
	var split = strings.Split(strings.TrimSuffix(input, "\n"), "\n\n")
	var split2 = strings.Split(split[0], "\n")
	var length, _ = strconv.Atoi(string(split2[len(split2)-1][len(split2[len(split2)-1])-2]))
	var stacks []Stack = make([]Stack, length)
	for i := range stacks {
		for j := length - 1; j >= 0 && split2[j][4*i+1] != ' '; j-- {
			add(&stacks[i], split2[j][4*i+1])
		}
	}
	var moves []Move
	for _, line := range strings.Split(split[1], "\n") {
		var count, start, end int
		fmt.Sscanf(line, "move %d from %d to %d", &count, &start, &end)
		moves = append(moves, Move{count, start, end})
	}
	return stacks, moves
}

func moveCrate1(stacks []Stack, move Move) {
	for range move.count {
		add(&stacks[move.end-1], pop(&stacks[move.start-1]))
	}
}

func moveCrate2(stacks []Stack, move Move) {
	var tempStack Stack
	for range move.count {
		add(&tempStack, pop(&stacks[move.start-1]))
	}
	for range move.count {
		add(&stacks[move.end-1], pop(&tempStack))
	}
}

func result(stacks []Stack) string {
	var res []byte = make([]byte, len(stacks))
	for i, stack := range stacks {
		res[i] = stack[len(stack)-1]
	}
	return string(res)
}

func solve(input string, funcMove func([]Stack, Move)) string {
	var stacks, moves = parse(input)
	for _, move := range moves {
		funcMove(stacks, move)
	}
	return result(stacks)
}

func part1(input string) string {
	return solve(input, moveCrate1)
}

func part2(input string) string {
	return solve(input, moveCrate2)
}

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("--2022 day 05 solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(string(input)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("part2 : ", part2(string(input)))
	fmt.Println(time.Since(start))
}
