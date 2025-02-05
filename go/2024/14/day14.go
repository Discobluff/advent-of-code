package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
	"time"
)

type Position struct {
	x, y int
}

type Robot struct {
	position, velocity Position
}

func parse(line string) Robot {
	var x, y, vx, vy int
	fmt.Sscanf(line, "p=%d,%d v=%d,%d", &x, &y, &vx, &vy)
	return Robot{Position{x, y}, Position{vx, vy}}
}

func parseSize(line string) (int, int) {
	var length, height int
	fmt.Sscanf(line, "%d,%d", &length, &height)
	return length, height
}

func modulo(a int, b int) int {
	if a < 0 {
		return modulo(a+b, b)
	}
	return a % b
}

func move(robots *[]Robot, height int, length int) {
	for i, robot := range *robots {
		(*robots)[i] = Robot{position: Position{x: modulo(robot.position.x+robot.velocity.x, length), y: modulo(robot.position.y+robot.velocity.y, height)}, velocity: robot.velocity}
	}
}

func countRobots(robots []Robot, height int, length int) int {
	var res []int = make([]int, 4)
	for _, robot := range robots {
		var x, y = robot.position.x, robot.position.y
		if x < length/2 {
			if y < height/2 {
				res[0]++
			}
			if y > height/2 {
				res[1]++
			}
		}
		if x > length/2 {
			if y < height/2 {
				res[2]++
			}
			if y > height/2 {
				res[3]++
			}
		}
	}
	return res[0] * res[1] * res[2] * res[3]
}

func part1(input string) int {
	var lines []string = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var robots []Robot = make([]Robot, len(lines)-1)
	var length, height int = parseSize(lines[0])
	for i, line := range lines[1:] {
		robots[i] = parse(line)
	}
	var loop int = 100
	for range loop {
		move(&robots, height, length)
	}
	return countRobots(robots, height, length)
}

func christmasTree(robots []Robot) bool {
	var robotsMap map[Position]struct{} = make(map[Position]struct{})
	for _, robot := range robots {
		var _, ok = robotsMap[robot.position]
		if ok {
			return false
		}
		robotsMap[robot.position] = struct{}{}
	}
	return true
}

func displayRobots(robots []Robot, length int, height int) {
	var grid [][]uint8 = make([][]uint8, height)
	for i := range height {
		grid[i] = make([]uint8, length)
	}
	for _, robot := range robots {
		grid[robot.position.y][robot.position.x] = 1
	}
	for _, line := range grid {
		fmt.Println(line)
	}
}

func part2(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var robots []Robot = make([]Robot, len(lines)-1)
	var length, height int = parseSize(lines[0])
	for i, line := range lines[1:] {
		robots[i] = parse(line)
	}
	var res int
	for !christmasTree(robots) {
		res++
		move(&robots, height, length)
	}
	// displayRobots(robots, length, height)
	return res
}

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("--2024 day 14 solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(string(input)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("part2 : ", part2(string(input)))
	fmt.Println(time.Since(start))
}
