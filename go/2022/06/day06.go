package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
	"time"
)

type Queue []byte

func add(q *Queue, e byte) {
	(*q) = append((*q), e)
}

func pop(q *Queue) {
	(*q) = (*q)[1:]
}

func allDifferents(q Queue) bool {
	for i := range q {
		for j := range q {
			if i != j && q[i] == q[j] {
				return false
			}
		}
	}
	return true
}

func solve(s string, size int) int {
	var queue Queue
	var count int
	for _, char := range s {
		if len(queue) == size {
			pop(&queue)
		}
		add(&queue, byte(char))
		count++
		if len(queue) == size && allDifferents(queue) {
			return count
		}
	}
	return -1
}

func part1(input string) int {
	var line = strings.TrimSuffix(input, "\n")
	return solve(line, 4)
}

func part2(input string) int {
	var line = strings.TrimSuffix(input, "\n")
	return solve(line, 14)
}

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("--2022 day 06 solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(string(input)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("part2 : ", part2(string(input)))
	fmt.Println(time.Since(start))
}
