package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
	"time"

	. "github.com/Discobluff/advent-of-code/go/utils/graph"
	. "github.com/Discobluff/advent-of-code/go/utils/set"
)

//go:embed input.txt
var input string

func parseLine(line string) Set[string] {
	var res = DefSet[string]()
	var computers = strings.Split(line, "-")
	Add(res, computers[0])
	Add(res, computers[1])
	return res
}

func parse(lines []string) []Set[string] {
	var res []Set[string] = make([]Set[string], len(lines))
	for i, line := range lines {
		res[i] = parseLine(line)
	}
	return res
}

func checkSet(set Set[string]) bool {
	for s := range set {
		if s[0] == 't' {
			return true
		}
	}
	return false
}

func part1(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var connections = parse(lines)
	var vertices = DefSet[string]()
	var edges = DefSet[Edge[string]]()
	for _, connection := range connections {
		vertices = Union(vertices, connection)
		var edge = SetToSlice(connection)
		Add(edges, Edge[string]{S1: edge[0], S2: edge[1]})
	}
	var graph = Graph[string]{Vertices: vertices, Edges: edges}
	var newGraph, bijection = GraphToGraphInt(graph)
	var clique3 = Find3Clique(newGraph)
	var clique3String []Set[string] = buildSetFromBijection(clique3, bijection)
	var res int
	for _, set := range clique3String {
		if checkSet(set) {
			res++
		}
	}
	return res
}

func buildSliceFromBijection(slice []int, bijection []string) []string {
	var res []string = make([]string, len(slice))
	for j, i := range slice {
		res[j] = bijection[i]
	}
	return res
}

func buildSetFromBijection(slice []Set[int], bijection []string) []Set[string] {
	var res []Set[string] = make([]Set[string], len(slice))
	for i, set := range slice {
		res[i] = make(Set[string])
		for j := range set {
			Add(res[i], bijection[j])
		}
	}
	return res
}

func initTab(n int) Set[int] {
	var res = DefSet[int]()
	for i := range n {
		Add(res, i)
	}
	return res
}

func part2(input string) string {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var connections = parse(lines)
	var vertices = DefSet[string]()
	var edges = DefSet[Edge[string]]()
	for _, connection := range connections {
		vertices = Union(vertices, connection)
		var edge = SetToSlice(connection)
		Add(edges, Edge[string]{S1: edge[0], S2: edge[1]})
	}
	var graph = Graph[string]{Vertices: vertices, Edges: edges}
	var maxCliqueString = SetToSlice(CliqueMaximum(graph))
	slices.Sort(maxCliqueString)
	var res []byte
	for _, s := range maxCliqueString {
		res = append(res, []byte(s)...)
		res = append(res, ',')
	}
	return string(res[:len(res)-1])
}

func main() {
	fmt.Println("--2024 day 23 solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(input))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("part2 : ", part2(input))
	fmt.Println(time.Since(start))
}
