package main

import (
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"github.com/Discobluff/advent-of-code/go/utils/segment"
	
)

func part1(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n\n")
	var segments = segment.ParseLinesToSetSegment(strings.Split(lines[0],"\n"), segment.ParseSegmentTiret)
	var ids = strings.Split(lines[1],"\n")
	var res = 0
	for _,id := range ids{
		for s := range segments{
			idInt,_ := strconv.Atoi(id)
			if segment.InSegment(s, idInt){
				res++
				break
			}
		}
	}
	return res
}

func part2(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n\n")
	var segments = segment.ParseLinesToSetSegment(strings.Split(lines[0],"\n"), segment.ParseSegmentTiret)
	segment.MergeSetSegment(segments)
	var res int
	for s := range segments{
		res += segment.Size(s)
	}
	return res
}

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("--2025 day 05 solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(string(input)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("part2 : ", part2(string(input)))
	fmt.Println(time.Since(start))
}
