package main

import (
	_ "embed"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

type Monkey struct {
	items     []int
	funcNew   func(int) int
	funcTest  func(int) bool
	testTrue  int
	testFalse int
}

func parse(line string) (Monkey, int) {
	var split = strings.Split(line, "\n")
	var res Monkey
	var numbersString = strings.Split(strings.Split(split[1], ": ")[1], ", ")
	for _, n := range numbersString {
		var conv, _ = strconv.Atoi(n)
		res.items = append(res.items, conv)
	}
	var op, count string
	fmt.Sscanf(split[2], "  Operation: new = old %s %s", &op, &count)
	var conv, _ = strconv.Atoi(count)
	res.funcNew = func(worry int) int {
		var second int = worry
		if count != "old" {
			second = conv
		}
		if op == "+" {
			return worry + second
		}
		return worry * second
	}
	var test int
	fmt.Sscanf(split[3], "  Test: divisible by %d", &test)
	res.funcTest = func(old int) bool {
		return old%test == 0
	}
	fmt.Sscanf(split[4], "    If true: throw to monkey %d", &res.testTrue)
	fmt.Sscanf(split[5], "    If false: throw to monkey %d", &res.testFalse)
	return res, test
}

func solve(input string, loop int, part2 bool) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n\n")
	var monkeys []Monkey
	var modulo int = 1
	for _, line := range lines {
		var parsed, m = parse(line)
		modulo *= m
		monkeys = append(monkeys, parsed)
	}
	var inspections []int = make([]int, len(monkeys))
	for range loop {
		for i := 0; i < len(monkeys); i++ {
			for _, worry := range monkeys[i].items {
				monkeys[i].items = monkeys[i].items[1:]
				inspections[i]++
				var new = monkeys[i].funcNew(worry)
				if !part2 {
					new = new / 3
				} else {
					new = new % modulo
				}
				if monkeys[i].funcTest(new) {
					monkeys[monkeys[i].testTrue].items = append(monkeys[monkeys[i].testTrue].items, new)
				} else {
					monkeys[monkeys[i].testFalse].items = append(monkeys[monkeys[i].testFalse].items, new)
				}
			}
		}
	}
	slices.Sort(inspections)
	return inspections[len(inspections)-1] * inspections[len(inspections)-2]
}

func part1(input string) int {
	return solve(input, 20, false)
}

func part2(input string) int {
	return solve(input, 10000, true)
}

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("--2024 day 11 solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(string(input)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("part2 : ", part2(string(input)))
	fmt.Println(time.Since(start))
}
