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

func parse(lines []string) Set[Input] {
	var res = DefSet[Input]()
	for _, line := range lines {
		var new Input
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &new.sensor.Column, &new.sensor.Line, &new.beacon.Column, &new.beacon.Line)
		Add(res, new)
	}
	return res
}

func size(intervals Set[Interval]) int {
	var res int
	for i := range intervals {
		res += i.b - i.a + 1
	}
	return res
}

func mergeIntervals(i1, i2 Interval, swap bool) (Interval, bool) {
	if i1.a <= i2.a && i2.b <= i1.b {
		return i1, true
	}
	// i1 <= i2
	if i2.a <= i1.b && i1.a <= i2.a {
		return Interval{a: i1.a, b: i2.b}, true
	}
	//i2 c i1
	if swap {
		return mergeIntervals(i2, i1, false)
	}
	return Interval{a: 0, b: 0}, false
}

func merge(i Set[Interval]) Set[Interval] {
	for i1 := range i {
		for i2 := range i {
			if i1 != i2 {
				var inter, ok = mergeIntervals(i1, i2, true)
				if ok {
					Remove(i, i1)
					Remove(i, i2)
					Add(i, inter)
					return merge(i)
				}
			}
		}
	}
	return i
}

func solveLine(input string, line int) Set[Interval] {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var set = parse(lines)
	var intervals = DefSet[Interval]()
	for s := range set {
		var distance = Distance(s.sensor, s.beacon)
		var a = s.sensor.Column - distance + Abs(-s.sensor.Line+line)
		var b = s.sensor.Column + distance - Abs(-s.sensor.Line+line)
		if s.beacon == (Position{Line: line, Column: a}) {
			a++
		}
		if s.beacon == (Position{Line: line, Column: b}) {
			b--
		}
		if a <= b {
			var interval Interval = Interval{a: a, b: b}
			Add(intervals, interval)
		}
	}
	return merge(intervals)
}

func part1(input string) int {
	if len(strings.Split(strings.TrimSuffix(input, "\n"), "\n")) == 14 {
		return size(solveLine(input, 10))
	}
	return size(solveLine(input, 2_000_000))
}

func convertInterval(intervals Set[Interval], min, max int) Set[int] {
	var res = DefSet[int]()
	for interval := range intervals {
		for i := interval.a; i <= interval.b; i++ {
			if min <= i && i <= max {
				Add(res, i)
			}
		}
	}
	return res
}

func size2(intervals Set[Interval], minBound, maxBound int) int {
	var res int
	for interval := range intervals {
		res += min(interval.b, maxBound) - max(interval.a, minBound) + 1
	}
	return res
}

func tuningFrequency(pos Position) int {
	return pos.Line + pos.Column*4000000
}

func solve2(input string, min, max int) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var set = parse(lines)
	for line := min; line <= max; line++ {
		var intervalI = solveLine(input, line)
		var beacons = DefSet[int]()
		for inp := range set {
			if inp.beacon.Line == line && inp.beacon.Column >= min && inp.beacon.Column <= max {
				Add(beacons, inp.beacon.Column)
			}
		}
		if size2(intervalI, min, max)+len(beacons) != max-min+1 {
			var allInts = convertInterval(intervalI, min, max)
			allInts = Union(allInts, beacons)
			var setAll = DefSet[Interval]()
			Add(setAll, Interval{a: min, b: max})
			var res = Deprived(convertInterval(setAll, min, max), allInts)
			var column int
			for obj := range res {
				column = obj
			}
			return tuningFrequency(Position{Line: line, Column: column})

		}
	}
	return -1
}

func part2(input string) int {
	if len(strings.Split(strings.TrimSuffix(input, "\n"), "\n")) == 14 {
		return solve2(input, 0, 20)
	}
	return solve2(input, 0, 4000000)
}

func main() {
	fmt.Println("--2024 day 15 solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(input))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("part2 : ", part2(input))
	fmt.Println(time.Since(start))
}
