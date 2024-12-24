package graph

import (
	"slices"

	. "github.com/Discobluff/advent-of-code/go/utils/set"
)

type GraphMatrix [][]bool

type Edge[T comparable] struct {
	S1, S2 T
}

type Graph[T comparable] struct {
	Vertices Set[T]
	Edges    Set[Edge[T]]
}

type GraphInt struct {
	Size  int
	Edges Set[Edge[int]]
}

func getCorrespond[T comparable](vertices map[T]int) []T {
	var res []T = make([]T, len(vertices))
	for key, val := range vertices {
		res[val] = key
	}
	return res
}

func GraphToGraphInt[T comparable](graph Graph[T]) (GraphInt, []T) {
	var res GraphInt
	res.Size = len(graph.Vertices)
	var verticesMap map[T]int = make(map[T]int)
	var index int
	for vertice := range graph.Vertices {
		verticesMap[vertice] = index
		index++
	}
	var newEdges = DefSet[Edge[int]]()
	for edge := range graph.Edges {
		Add(newEdges, Edge[int]{S1: verticesMap[edge.S1], S2: verticesMap[edge.S2]})
	}
	res.Edges = newEdges
	return res, getCorrespond(verticesMap)
}

func GraphIntToGraphMatrix(graph GraphInt) GraphMatrix {
	var matrix [][]bool = make([][]bool, graph.Size)
	for i := range graph.Size {
		matrix[i] = make([]bool, graph.Size)
	}
	for edge := range graph.Edges {
		matrix[edge.S1][edge.S2] = true
		matrix[edge.S2][edge.S1] = true
	}
	return matrix
}

func sortBySize(s1 Set[int], s2 Set[int]) int {
	if len(s1) < len(s2) {
		return 1
	}
	if len(s2) < len(s1) {
		return -1
	}
	return 0
}

func removeVertice(neighbors []Set[int], vertice int) {
	for _, neighbor := range neighbors {
		if In(neighbor, vertice) {
			Remove(neighbor, vertice)
		}
	}
}

func funcEqual(s Set[int]) func(Set[int]) bool {
	return func(e Set[int]) bool {
		return Equal(s, e)
	}
}

// Algorithm from Chiba & Nishizeki
func Find3Clique(graph GraphInt) []Set[int] {
	var neighbors []Set[int] = make([]Set[int], graph.Size)
	for i := range neighbors {
		neighbors[i] = DefSet[int]()
	}
	for edge := range graph.Edges {
		Add(neighbors[edge.S1], edge.S2)
		Add(neighbors[edge.S2], edge.S1)
	}
	slices.SortFunc(neighbors, sortBySize)
	var res []Set[int]
	for v, neighborhood := range neighbors {
		for n1 := range neighborhood {
			for n2 := range neighbors[n1] {
				if In(neighborhood, n2) {
					var clique3 = DefSet[int]()
					Add(clique3, v)
					Add(clique3, n1)
					Add(clique3, n2)
					if !slices.ContainsFunc(res, funcEqual(clique3)) {
						res = append(res, clique3)
					}
				}
			}
		}
		removeVertice(neighbors, v)
	}
	return res
}

// Algorithm of Bron-Kerbosch
func bronKerbosch[T comparable](neighbors map[T]Set[T], clique Set[T], candidates Set[T], excluded Set[T], maxClique *Set[T]) {
	if IsEmpty(candidates) && IsEmpty(excluded) {
		if len(clique) > len(*maxClique) {
			*maxClique = clique
		}
		return
	}
	for v := range candidates {
		var setV = DefSet[T]()
		Add(setV, v)
		bronKerbosch(neighbors, Union(clique, setV), Intersect(candidates, neighbors[v]), Intersect(excluded, neighbors[v]), maxClique)
		Remove(candidates, v)
		Add(excluded, v)
	}
}

func CliqueMaximum[T comparable](graph Graph[T]) Set[T] {
	var neighbors map[T]Set[T] = make(map[T]Set[T])
	for vertice := range graph.Vertices {
		neighbors[vertice] = DefSet[T]()
	}
	for edge := range graph.Edges {
		Add(neighbors[edge.S1], edge.S2)
		Add(neighbors[edge.S2], edge.S1)
	}
	var clique = DefSet[T]()
	var excluded = DefSet[T]()
	var candidates = Union(graph.Vertices, DefSet[T]())
	var res = DefSet[T]()
	bronKerbosch(neighbors, clique, candidates, excluded, &res)
	return res
}

func isConnected(g GraphMatrix, v1, v2 int) bool {
	return g[v1][v2]
}
func intersect(g GraphMatrix, vertices []int, v int) []int {
	var result []int
	for _, u := range vertices {
		if isConnected(g, v, u) {
			result = append(result, u)
		}
	}
	return result
}

func remove(slice []int, value int) []int {
	for i, v := range slice {
		if v == value {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}
