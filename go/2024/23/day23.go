package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
	"time"

	. "github.com/Discobluff/advent-of-code/go/utils/set"
)

//go:embed input.txt
var input string

func parseLine(line string) Set[string] {
	var res = DefSet[string]()
	var computers = strings.Split(line, "-")
	Add(res, computers[0])
	Add(res, computers[1])
	return res
}

func parse(lines []string) []Set[string] {
	var res []Set[string] = make([]Set[string], len(lines))
	for i, line := range lines {
		res[i] = parseLine(line)
	}
	return res
}

func funcEqual(s Set[string]) func(Set[string]) bool {
	return func(e Set[string]) bool {
		return Equal(s, e)
	}
}

func checkSet(set Set[string]) bool {
	for s := range set {
		if s[0] == 't' {
			return true
		}
	}
	return false
}

func buildConnections(connections []Set[string]) map[string]Set[string] {
	var res map[string]Set[string] = make(map[string]Set[string])
	for _, connection := range connections {
		var computers = SetToSlice(connection)

		var _, ok1 = res[computers[0]]
		if !ok1 {
			res[computers[0]] = DefSet[string]()
		}
		var _, ok2 = res[computers[1]]
		if !ok2 {
			res[computers[1]] = DefSet[string]()
		}
		Add(res[computers[0]], computers[1])
		Add(res[computers[1]], computers[0])
	}
	return res
}

func get3Network(input string) []Set[string] {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var connections = buildConnections(parse(lines))
	var connected3 []Set[string]
	for computer1, connected := range connections {
		for computer2 := range connected {
			for computer3 := range connected {
				if In(connections[computer2], computer3) {
					var inter = DefSet[string]()
					Add(inter, computer1)
					Add(inter, computer2)
					Add(inter, computer3)
					if !slices.ContainsFunc(connected3, funcEqual(inter)) {
						connected3 = append(connected3, inter)
					}
				}
			}
		}
	}
	return connected3
}

func part1(input string) int {
	var connected3 = get3Network(input)
	var res int
	for _, set := range connected3 {
		if checkSet(set) {
			res++
		}
	}
	return res
}

func part22(input string) string {
	// var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var connected3 = get3Network(input)
	// var connections = buildConnections(parse(lines))
	var nextConnected []Set[string]
	fmt.Println("ready!!", len(connected3))
	for i, c1 := range connected3 {
		fmt.Println(i)
		for _, c2 := range connected3 {
			for _, c3 := range connected3 {
				if len(Intersect(c1, Intersect(c2, c3))) == 1 && len(Union(c1, Union(c2, c3))) == 4 {
					var add = Union(c1, Union(c2, c3))
					if !slices.ContainsFunc(nextConnected, funcEqual(add)) {
						nextConnected = append(nextConnected, add)
					}

				}
			}
		}
	}
	fmt.Println(nextConnected, len(nextConnected))
	return "ui"
}

func part2(input string) []string {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var connections = parse(lines)
	var computers = DefSet[string]()
	for _, connection := range connections {
		computers = Union(computers, connection)
	}
	var connectedAll []string
	fmt.Println("ui")
	for computer := range computers {
		var connected = DefSet[string]()
		for _, connection := range connections {
			if In(connection, computer) {
				Add(connected, SetToSlice(Without(connection, computer))[0])
			}
		}
		if Equal(connected, Without(computers, computer)) {
			connectedAll = append(connectedAll, computer)
		}
		fmt.Println(computer, connected)
	}
	slices.Sort(connectedAll)
	return connectedAll
}

func main() {
	fmt.Println("--2024 day 23 solution--")
	start := time.Now()
	// fmt.Println("part1 : ", part1(input))
	// fmt.Println(time.Since(start))
	// start = time.Now()
	fmt.Println("part2 : ", part22(input))
	fmt.Println(time.Since(start))
}
