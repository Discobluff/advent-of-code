package main

import (
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	id, start, length int
}

func initTab(line string) []int {
	var res []int
	for i, char := range line {
		var val int
		if i%2 == 1 {
			val = -1
		} else {
			val = i / 2
		}
		var count, _ = strconv.Atoi(string(char))
		for range count {
			res = append(res, val)
		}
	}
	return res
}

func createBlock(id int, start int, length int) Block {
	var res Block
	res.id = id
	res.start = start
	res.length = length
	return res
}

func parse(line string) []Block {
	var res []Block
	var shift int
	var isPoint bool
	var i int
	for _, val := range line {
		var intVal, _ = strconv.Atoi(string(val))
		if !isPoint {
			res = append(res, createBlock(i, shift, intVal))
			i++
		}
		isPoint = !isPoint
		shift += intVal
	}
	return res
}

func solve1(line string) int {
	var tab []int = initTab(line)
	var end int = len(tab) - 1
	var start int
	var res int
	for start < end {
		for tab[start] != -1 {
			res += start * tab[start]
			start++
		}
		for tab[end] == -1 {
			end--
		}
		if start < end {
			tab[start] = tab[end]
			tab[end] = -1
		}
	}
	return res
}

// Insert value which is at index i at index j
func insert(tab *[]Block, i int, j int) {
	var temp Block = (*tab)[i]
	for k := i; k > j; k-- {
		(*tab)[k] = (*tab)[k-1]

	}
	(*tab)[j] = temp
}

func search(tab []Block, id int) int {
	for j, block := range tab {
		if block.id == id {
			return j
		}
	}
	return -1
}

func solve2(line string) int {
	var tab []Block = parse(line)
	var id = len(tab) - 1
	var index = search(tab, id)
	var res int
	for tab[index].id > 0 {
		for i := range index {
			if tab[i+1].start-(tab[i].start+tab[i].length) >= tab[index].length {
				tab[index].start = tab[i].start + tab[i].length
				insert(&tab, index, i+1)
				index = i + 1
				break
			}
		}
		res += tab[index].length*tab[index].id*tab[index].start + tab[index].id*tab[index].length*(tab[index].length-1)/2
		id--
		index = search(tab, id)
	}
	return res
}

func part1(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	return solve1(lines[0])
}

func part2(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	return solve2(lines[0])
}
func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("--2024 day 09 solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(string(input)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("part2 : ", part2(string(input)))
	fmt.Println(time.Since(start))
}
