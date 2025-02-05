package main

import (
	_ "embed"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	. "github.com/Discobluff/advent-of-code/go/utils/positions"
	// . "github.com/Discobluff/advent-of-code/go/utils/slice"
)

type Path struct {
	Position
	path      []byte
	changeDir int
}

var keypad map[byte]Position = map[byte]Position{'7': DefPosition(0, 0), '8': DefPosition(0, 1), '9': DefPosition(0, 2), '4': DefPosition(1, 0), '5': DefPosition(1, 1), '6': DefPosition(1, 2), '1': DefPosition(2, 0), '2': DefPosition(2, 1), '3': DefPosition(2, 2), ' ': DefPosition(3, 0), '0': DefPosition(3, 1), 'A': DefPosition(3, 2)}
var directionalpad map[byte]Position = map[byte]Position{' ': DefPosition(0, 0), '^': DefPosition(0, 1), 'A': DefPosition(0, 2), '<': DefPosition(1, 0), 'v': DefPosition(1, 1), '>': DefPosition(1, 2)}

func copyPath(path []byte) []byte {
	var res []byte = make([]byte, len(path))
	for i, b := range path {
		res[i] = b
	}
	return res
}

func copyPaths(tab []Path) []Path {
	var res []Path = make([]Path, len(tab))
	for i, t := range tab {
		res[i] = Path{Position: t.Position, changeDir: t.changeDir, path: copyPath(t.path)}
	}
	return res
}

func nextPath(pad map[byte]Position, startB byte, endB byte) []byte {
	var paths []Path
	var start Position = pad[startB]
	var end Position = pad[endB]
	paths = append(paths, Path{Position: start, changeDir: 0, path: []byte{}})
	var res [][]byte
	for len(paths) > 0 {
		var newPaths []Path = make([]Path, 0)
		for _, path := range paths {
			if path.Position == end {
				res = append(res, append(path.path, 'A'))
			} else {
				for dir, posDir := range Directions {
					var newPos = AddPositions(path.Position, posDir)
					if (len(path.path) == 0 || dir == path.path[len(path.path)-1] || path.changeDir < 2) && newPos != pad[' '] && Distance(newPos, end) < Distance(path.Position, end) {
						if len(path.path) > 0 && dir == path.path[len(path.path)-1] {
							newPaths = append(newPaths, Path{Position: newPos, changeDir: path.changeDir, path: append(path.path, dir)})
						} else {
							newPaths = append(newPaths, Path{Position: newPos, changeDir: path.changeDir + 1, path: append(path.path, dir)})
						}
					}
				}

			}
		}
		paths = copyPaths(newPaths)
	}
	if len(res) == 1 {
		return res[0]
	}
	var ordre = []byte{'<', '^', 'v', '>'}
	if slices.Index(ordre, res[0][0]) < slices.Index(ordre, res[0][1]) {
		return res[0]
	}
	return res[1]

	// var ordre = []byte{'^', '>', 'v', '<'}
	// if res[0][0] == res[1][0] {
	// 	fmt.Println(res)
	// }
	// if Distance(directionalpad[res[0][0]], directionalpad['A']) < Distance(directionalpad[res[1][0]], directionalpad['A']) {
	// 	return res[0]
	// }
	// return res[1]

}

func complexity(s string) int {
	var res, _ = strconv.Atoi(s[:len(s)-1])
	return res
}

func solve(code map[string]int, robots int, pad map[byte]Position) map[string]int {
	if robots == 0 {
		return code
	}
	var codeMap map[string]int = make(map[string]int)
	for char, count := range code {
		var start = byte('A')
		for _, c := range char {
			var path = nextPath(pad, byte(start), byte(c))
			start = byte(c)
			codeMap[string(path)] += count
		}
	}
	return solve(codeMap, robots-1, directionalpad)
}

func length(code map[string]int) int {
	var res int
	for c, count := range code {
		res += len(c) * count
	}
	return res
}

func build(s string) map[string]int {
	var res map[string]int = make(map[string]int)
	var split = strings.Split(s, "A")
	for _, spl := range split {
		res[spl+"A"] += 1
	}
	return res

}

func part1(input string) int {
	var codes = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var res int
	var robots int = 3
	for _, code := range codes {
		var codeMap map[string]int = make(map[string]int)
		codeMap[code] = 1
		var path = solve(codeMap, robots, keypad)
		res += length(path) * complexity(code)
	}
	return res
}

func part2(input string) int {
	var codes = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var res int
	var robots int = 26
	for _, code := range codes {
		var codeMap map[string]int = make(map[string]int)
		codeMap[code] = 1
		var path = solve(codeMap, robots, keypad)
		res += length(path) * complexity(code)
	}
	return res
}

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("--2024 day 21 solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(string(input)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("part2 : ", part2(string(input)))
	fmt.Println(time.Since(start))
}
