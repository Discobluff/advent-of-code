package search

import (
	. "github.com/Discobluff/advent-of-code/go/utils/priorityQueue"
	. "github.com/Discobluff/advent-of-code/go/utils/set"
)

func cmpDijkstra[T comparable](scores map[T]int) func(T, T) bool {
	return func(p1 T, p2 T) bool {
		return scores[p1] <= scores[p2]
	}
}

func bestDijkstra[T comparable](scores map[T]int, t1 T, score2 int) int {
	var score1, ok = scores[t1]
	if !ok {
		return score2
	}
	return min(score1, score2)
}

func Dijkstra[T comparable](start T, neighbors func(T) map[T]int) map[T]int {
	var scores map[T]int = make(map[T]int)
	scores[start] = 0
	var nexts = DefQueue[T]()
	AddQueue(&nexts, start, cmpDijkstra(scores))
	var visited Set[T] = DefSet[T]()
	for !IsEmpty(nexts) {
		var next = PopQueue(&nexts)
		var _, ok = visited[next]
		if !ok {
			Add(visited, next)
			for neighbor, cost := range neighbors(next) {
				scores[neighbor] = bestDijkstra(scores, neighbor, scores[next]+cost)
				AddQueue(&nexts, neighbor, cmpDijkstra(scores))
			}
		}
	}
	return scores
}
