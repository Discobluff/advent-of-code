package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var inputDay string

func addToDict(line string, dict *map[string][2]string) {
	var key string = strings.Split(line, " = ")[0]
	var value [2]string
	value[0] = strings.Split(strings.Split(line, "(")[1], ", ")[0]
	value[1] = strings.Split(strings.Split(line, ", ")[1], ")")[0]
	(*dict)[key] = value
}

func endByChar(str string, char byte) bool {
	return str[len(str)-1] == char
}

func part1(lines []string, start string, end string) int {
	var rule string = lines[0]
	var dict map[string][2]string = make(map[string][2]string)
	for _, line := range lines[2:] {
		addToDict(line, &dict)
	}
	var count int = 0
	for start != end {
		if rule[count%len(rule)] == 'L' {
			start = dict[start][0]
		}
		if rule[count%len(rule)] == 'R' {
			start = dict[start][1]
		}
		count += 1
	}
	return count

}

func countTrue(tab [][2]int) int {
	var res int
	for _, b := range tab {
		if b[0] != -1 {
			res += 1
		}
	}
	return res
}

func explore(start string, dict map[string][2]string, rule string, index int, end byte) (int, string) {
	var count int
	var stop bool
	for !stop {
		if rule[index%len(rule)] == 'L' {
			start = dict[start][0]
		} else {
			start = dict[start][1]
		}
		count += 1
		index += 1
		stop = endByChar(start, end)
	}
	return count, start
}

func part2(lines []string) int {
	var rule string = lines[0]
	var dict map[string][2]string = make(map[string][2]string)
	for _, line := range lines[2:] {
		addToDict(line, &dict)
	}
	var starts []string
	for key := range dict {
		if endByChar(key, 'A') {
			starts = append(starts, key)
		}
	}
	var res int = 1
	for _, start := range starts {
		firstZ, end := explore(start, dict, rule, 0, 'Z')
		secondZ, _ := explore(end, dict, rule, firstZ, 'Z')
		res = lcm(res, secondZ)
	}
	return res
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func main() {
	var input = strings.TrimSuffix(inputDay, "\n")
	var lines = strings.Split(input, "\n")
	fmt.Println(part1(lines, "AAA", "ZZZ"))
	fmt.Println(part2(lines))
}
