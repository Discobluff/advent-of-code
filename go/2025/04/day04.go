package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
	"time"
	"github.com/Discobluff/advent-of-code/go/utils/positions"
)

func isValid(lines []string, pos positions.Position)bool{
	return pos.Line >=0 && pos.Column >= 0 && pos.Line < len(lines) && pos.Column < len(lines[0])
}

func isAccessible(lines []string, pos positions.Position)bool{
	var count = 0
	for line:=-1; line<=1;line++{
		for column := (-1);column<=1;column++{
			if line != 0 || column != 0{
				posTemp := positions.DefPosition(line, column)
				newPos := positions.AddPositions(posTemp, pos)
				if isValid(lines,newPos){
					if lines[newPos.Line][newPos.Column] == '@'{
						count++
					}
				}
			}
		}
	}
	return count < 4
}

func part1(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var count int

		for l,line:=range lines{
			for c:=range line{
				if lines[l][c] == '@' && isAccessible(lines,positions.DefPosition(l,c)){

					count++
				}
			}
		}
	
	return count
}

func copyStringWithChange(str string, index int)string{
	res := []byte(str)
	res[index] = '.'
	return string(res)
}

func part2(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var count int
	var tempCount int = 1
	for count != tempCount{
		tempCount = count
		for l,line:=range lines{
			for c:=range line{
				if lines[l][c] == '@' && isAccessible(lines,positions.DefPosition(l,c)){
					lines[l] = copyStringWithChange(lines[l], c)
					count++
				}
			}
		}
	}
	return count
}

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("--2025 day 04 solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(string(input)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("part2 : ", part2(string(input)))
	fmt.Println(time.Since(start))
}
