package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
	"time"

	. "github.com/Discobluff/advent-of-code/go/utils/search"
	. "github.com/Discobluff/advent-of-code/go/utils/positions"
)

type State struct {
	pos, direction Position
	streak int
}

func isValidPosition(grid []string, pos Position) int{
	if pos.Line >= len(grid) || pos.Column >= len(grid[0]) || pos.Line < 0 || pos.Column < 0{
		return -1
	}
	return (int)(grid[pos.Line][pos.Column]) - '0'
}

func part1(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var start = State{pos:DefPosition(0, 0), direction: DefPosition(0,0), streak: 0}
	var fun = func(state State) map[State]int {
		var res = make(map[State]int)
		for _, direction := range Directions {
			if state.direction != direction && state.direction != OpposedDirection(direction){
				var newPos = AddPositions(state.pos, direction)
				var val = isValidPosition(lines, newPos)
				if val != -1{
					var newState = State{pos: newPos, direction: direction, streak: 1}
					res[newState] = val
				}
			}
			if state.direction == direction && state.streak < 3{
				var newPos = AddPositions(state.pos, direction)
				var val = isValidPosition(lines, newPos)
				if val != -1{
					var newState = State{pos: newPos, direction: direction, streak: state.streak+1}
					res[newState] = val
				}
			}
		}
		return res
	}
	var resDij = Dijkstra(start, fun)
	var res = -1
	var end = DefPosition(len(lines)-1, len(lines[0])-1)
	for state,cost := range resDij{
		if state.pos == end{
			if res == -1 || cost < res{
				res = cost
			}
		}
	}
	return res
}

func part2(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var start = State{pos:DefPosition(0, 0), direction: DefPosition(0,0), streak: 0}
	var fun = func(state State) map[State]int {
		var res = make(map[State]int)
		for _, direction := range Directions {
			if state.direction != direction && state.direction != OpposedDirection(direction) && (state.streak >= 4 || state.streak == 0){
				var newPos = AddPositions(state.pos, MultPosition(direction,1))
				var val = isValidPosition(lines, newPos)
				if val != -1{
					var newState = State{pos: newPos, direction: direction, streak: 1}
					res[newState] = val
				}
			}
			if state.direction == direction && state.streak < 10{
				var newPos = AddPositions(state.pos, direction)
				var val = isValidPosition(lines, newPos)
				if val != -1{
					var newState = State{pos: newPos, direction: direction, streak: state.streak+1}
					res[newState] = val
				}
			}
		}
		return res
	}
	var resDij = Dijkstra(start, fun)
	var res = -1
	var end = DefPosition(len(lines)-1, len(lines[0])-1)
	for state,cost := range resDij{
		if state.pos == end{
			if res == -1 || cost < res{
				res = cost
			}
		}
	}
	return res
}

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("--2023 day 17 solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(string(input)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("part2 : ", part2(string(input)))
	fmt.Println(time.Since(start))
}
