package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"

	. "github.com/Discobluff/advent-of-code/go/utils/positions"
)

//go:embed input.txt
var input string

type Path struct {
	Position
	path []byte
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
		res[i] = Path{Position: t.Position, path: copyPath(t.path)}
	}
	return res
}

func pathsKeyPad(start Position, end Position, gap Position) [][]byte {
	var paths []Path
	paths = append(paths, Path{Position: start, path: []byte{}})
	var res [][]byte
	for len(paths) > 0 {
		var newPaths []Path = make([]Path, 0)
		for _, path := range paths {
			if path.Position == end {
				res = append(res, append(path.path, 'A'))
			} else {
				for dir, posDir := range Directions {
					var newPos = AddPositions(path.Position, posDir)
					if newPos != gap && Distance(newPos, end) < Distance(path.Position, end) {
						newPaths = append(newPaths, Path{Position: newPos, path: append(path.path, dir)})
					}
				}
			}
		}
		paths = copyPaths(newPaths)
	}
	return res
}

func cartesianProduct(t1 [][]byte, t2 [][]byte) [][]byte {
	if len(t1) == 0 {
		return t2
	}
	var res [][]byte = make([][]byte, len(t1)*len(t2))
	var index int
	for _, b1 := range t1 {
		for _, b2 := range t2 {
			res[index] = append(b1, b2...)
			index++
		}
	}
	return res
}

func displayPath(path [][]byte) {
	for _, p := range path {
		fmt.Printf("%s\n", p)
	}
}

func deepCopy(t [][]byte) [][]byte {
	var res = make([][]byte, len(t))
	for i, b := range t {
		res[i] = copyPath(b)
	}
	return res
}

func lenMini(t [][]byte) (int, int) {
	var res int = -1
	var count int = 0
	for _, path := range t {
		if len(path) == res {
			count++
		}
		if res == -1 || len(path) < res {
			res = len(path)
			count = 1
		}
	}
	return res, count
}

func complexity(s string) int {
	var res, _ = strconv.Atoi(s[:len(s)-1])
	return res
}

func part1(input string) int {
	var res int
	var codes = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var pads []map[byte]Position = []map[byte]Position{keypad, directionalpad, directionalpad}
	for _, code := range codes {
		var paths [][]byte
		paths = append(paths, []byte(code))
		for i := range 3 {
			var pathKeyPad [][]byte
			for _, path := range paths {
				var start Position = pads[i]['A']
				var path2 [][]byte
				for _, digit := range path {
					path2 = deepCopy(cartesianProduct(path2, pathsKeyPad(start, pads[i][byte(digit)], pads[i][' '])))
					start = pads[i][byte(digit)]
				}
				pathKeyPad = append(pathKeyPad, path2...)
			}
			paths = deepCopy(pathKeyPad)
			paths = pathKeyPad
		}
		var m, _ = lenMini(paths)
		res += m * complexity(code)
	}
	return res
}

// func part2(input string) int {
// 	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
// 	var res int
// 	return res
// }

func main() {
	fmt.Println("--2024 day 21 solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(input))
	fmt.Println(time.Since(start))
	// start = time.Now()
	// fmt.Println("part2 : ", part2(input))
	// fmt.Println(time.Since(start))
}
