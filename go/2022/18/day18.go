package main

import (
	_ "embed"
	"fmt"
	"math"
	"os"
	"strings"
	"time"

	. "github.com/Discobluff/advent-of-code/go/utils/search"
	. "github.com/Discobluff/advent-of-code/go/utils/set"
)

type Point struct {
	x, y, z int
}

func parse(lines []string) Set[Point] {
	var res Set[Point] = DefSet[Point]()
	for _, line := range lines {
		var point Point
		fmt.Sscanf(line, "%d,%d,%d", &point.x, &point.y, &point.z)
		Add(res, point)
	}
	return res
}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}

func solve(points Set[Point]) int {
	var res = len(points) * 6
	for p1 := range points {
		for p2 := range points {
			if p1 != p2 && ((p1.x == p2.x && p1.y == p2.y && abs(p1.z-p2.z) == 1) || (p1.x == p2.x && p1.z == p2.z && abs(p1.y-p2.y) == 1) || (p1.z == p2.z && p1.y == p2.y && abs(p1.x-p2.x) == 1)) {
				res--
			}
		}
	}
	return res
}

func part1(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var points = parse(lines)
	return solve(points)
}

func getSize(points Set[Point]) (int, int, int, int, int, int) {
	var minX, maxX, minY, maxY, minZ, maxZ int = -1, -1, -1, -1, -1, -1
	for p := range points {
		if minX == -1 || minX > p.x {
			minX = p.x
		}
		if maxX == -1 || maxX < p.x {
			maxX = p.x
		}
		if minY == -1 || minY > p.y {
			minY = p.y
		}
		if maxY == -1 || maxY < p.y {
			maxY = p.y
		}
		if minZ == -1 || minZ > p.z {
			minZ = p.z
		}
		if maxZ == -1 || maxZ < p.z {
			maxZ = p.z
		}
	}
	return minX, maxX, minY, maxY, minZ, maxZ
}

func touchCorner(p Point, minX, maxX, minY, maxY, minZ, maxZ int) bool {
	return p.x == minX || p.x == maxX || p.y == minY || p.y == maxY || p.z == minZ || p.z == maxZ
}

func touchCornerPoints(points Set[Point], minX, maxX, minY, maxY, minZ, maxZ int) bool {
	for p := range points {
		if touchCorner(p, minX, maxX, minY, maxY, minZ, maxZ) {
			return true
		}
	}
	return false
}

func isValid(p Point, minX, maxX, minY, maxY, minZ, maxZ int) bool {
	return p.x >= minX && p.x <= maxX && p.y >= minY && p.y <= maxY && p.z >= minZ && p.z <= maxZ
}

func neighborsFunc(points Set[Point], minX, maxX, minY, maxY, minZ, maxZ int) func(Point) Set[Point] {
	return func(p Point) Set[Point] {
		var res = DefSet[Point]()
		var candidates = DefSet[Point]()
		Add(candidates, Point{x: p.x + 1, y: p.y, z: p.z})
		Add(candidates, Point{x: p.x - 1, y: p.y, z: p.z})
		Add(candidates, Point{x: p.x, y: p.y + 1, z: p.z})
		Add(candidates, Point{x: p.x, y: p.y - 1, z: p.z})
		Add(candidates, Point{x: p.x, y: p.y, z: p.z + 1})
		Add(candidates, Point{x: p.x, y: p.y, z: p.z - 1})
		for c := range candidates {
			if !In(points, c) && isValid(c, minX, maxX, minY, maxY, minZ, maxZ) {
				Add(res, c)
			}
		}
		return res
	}
}

func part2(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var points = parse(lines)
	var res = solve(points)
	var minX, maxX, minY, maxY, minZ, maxZ = getSize(points)
	var interiorPoints = DefSet[Point]()
	var exteriorPoints = BFS(Point{x: minX - 1, y: minY - 1, z: minZ - 1}, neighborsFunc(points, minX-1, maxX+1, minY-1, maxY+1, minZ-1, maxZ+1))
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			for z := minZ; z <= maxZ; z++ {
				var p = Point{x: x, y: y, z: z}
				if !In(points, p) && !In(interiorPoints, p) && !In(exteriorPoints, p) {
					var neighbors = BFS(p, neighborsFunc(points, minX, maxX, minY, maxY, minZ, maxZ))
					interiorPoints = Union(interiorPoints, neighbors)
				}
			}
		}
	}
	return res - solve(interiorPoints)
}

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("--2022 day 18 solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(string(input)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("part2 : ", part2(string(input)))
	fmt.Println(time.Since(start))
}
