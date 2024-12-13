package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input string

type Comb struct {
	a, b int
}

type Machine struct {
	x, y int
}

func tokens(comb Comb) int {
	return 3*comb.a + comb.b
}

func findComb(machineA Machine, machineB Machine, result Machine, part2 bool) int {
	if part2 {
		result.x += 10000000000000
		result.y += 10000000000000
	}
	var a = result.y*machineB.x - result.x*machineB.y
	var b = result.x*machineA.y - result.y*machineA.x
	var quotient = machineA.y*machineB.x - machineB.y*machineA.x
	if a%quotient != 0 || b%quotient != 0 {
		return 0
	}
	return tokens(Comb{a: a / quotient, b: b / quotient})
}

//(L1): a*xA+b*xB = xR
//(L2): a*yA+b*yB = yR
//(L3)=(L1)-xA/yA*(L2): b(xB-xA*yB/yA) = xR-yR*xA/yA
//yA*(L3): b(xB*yA-xA*yB) = xR*yA-yR*xA
//On peut ainsi déduire a et b

func parseMachine(lineButton string) Machine {
	var x, _ = strconv.Atoi(strings.Split(strings.Split(lineButton, "+")[1], ",")[0])
	var y, _ = strconv.Atoi(strings.Split(lineButton, "+")[2])
	return Machine{x: x, y: y}
}

func parseResult(lineResult string) Machine {
	var x, _ = strconv.Atoi(strings.Split(strings.Split(lineResult, "=")[1], ",")[0])
	var y, _ = strconv.Atoi(strings.Split(lineResult, "=")[2])
	return Machine{x: x, y: y}
}

func solve(input string, part2 bool) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var res int
	for i := 0; i < len(lines); i += 4 {
		var a = parseMachine(lines[i])
		var b = parseMachine(lines[i+1])
		var result = parseResult(lines[i+2])
		var tok = findComb(a, b, result, part2)
		res += tok
	}
	return res
}

func part1(input string) int {
	return solve(input, false)
}

func part2(input string) int {
	return solve(input, true)
}

func main() {
	fmt.Println("--2024 day 13 solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(input))
	fmt.Println(time.Since(start))
	fmt.Println()
	start = time.Now()
	fmt.Println("part2 : ", part2(input))
	fmt.Println(time.Since(start))
}
