package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"

	. "github.com/Discobluff/advent-of-code/go/utils/positions"
	. "github.com/Discobluff/advent-of-code/go/utils/set"
)

//go:embed input.txt
var input string

type Input struct {
	sensor, beacon Position
}

type Interval struct {
	a, b int
}

func addInterval(intervals []Interval, interval Interval) []Interval {
	for index, i := range intervals {
		if i.b <= interval.a {
			intervals[index] = Interval{a: i.a, b: interval.b}
			return intervals
		}
		if interval.b <= i.a {
			intervals[index] = Interval{a: interval.a, b: i.b}
			return intervals
		}
	}
	return append(intervals, interval)
}

func parse(lines []string) Set[Input] {
	var res = DefSet[Input]()
	for _, line := range lines {
		var new Input
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &new.sensor.Column, &new.sensor.Line, &new.beacon.Column, &new.beacon.Line)
		Add(res, new)
	}
	return res
}

func solve1(input string, line int) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var set = parse(lines)
	var parcourus = DefSet[Position]()
	for s := range set {
		var distance = Distance(s.sensor, s.beacon)
		for column := s.sensor.Column - distance + Abs(-s.sensor.Line+line); column <= s.sensor.Column+distance-Abs(-s.sensor.Line+line); column++ {
			var pos = Position{Line: line, Column: column}
			if pos == s.beacon {
				continue
			}
			if Distance(pos, s.sensor) <= distance {
				Add(parcourus, pos)
			} else {
				fmt.Println("ioio")
			}
		}
	}
	return len(parcourus)
	// return res
}

func part1(input string) int {
	if len(strings.Split(strings.TrimSuffix(input, "\n"), "\n")) == 14 {
		return solve1(input, 10)
	}
	return solve1(input, 2_000_000)

}

// func part2(input string) int {
// 	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
// 	var res int
// 	return res
// }

func main() {
	fmt.Println("--2024 day 15 solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(input))
	fmt.Println(time.Since(start))
	// start = time.Now()
	// fmt.Println("part2 : ", part2(input))
	// fmt.Println(time.Since(start))
}
