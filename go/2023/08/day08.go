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

func main() {
	var input = strings.TrimSuffix(inputDay, "\n")
	var lines = strings.Split(input, "\n")
	fmt.Println(part1(lines, "AAA", "ZZZ"))
}
