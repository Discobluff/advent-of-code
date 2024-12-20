package search

import (
	. "github.com/Discobluff/advent-of-code/go/utils/priorityQueue"
	. "github.com/Discobluff/advent-of-code/go/utils/set"
)

type Score struct {
	Score, Heuristic int
}

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
		if !In(visited, next) {
			Add(visited, next)
			for neighbor, cost := range neighbors(next) {
				scores[neighbor] = bestDijkstra(scores, neighbor, scores[next]+cost)
				AddQueue(&nexts, neighbor, cmpDijkstra(scores))
			}
		}
	}
	return scores
}

func evalHeuristic(s Score, a int, b int) int {
	return s.Score*a + s.Heuristic*b
}

func cmpAStar[T comparable](scores map[T]Score, score int, heuristic int) func(T, T) bool {
	return func(p1 T, p2 T) bool {
		return evalHeuristic(scores[p1], score, heuristic) <= evalHeuristic(scores[p2], score, heuristic)
	}
}

func AStar[T comparable](start T, end T, funcHeuristic func(T) int, neighbors func(T) map[T]int, score int, heuristic int) map[T]Score {
	var scores map[T]Score = make(map[T]Score)
	scores[start] = Score{Score: 0, Heuristic: 0}
	var nexts = DefQueue[T]()
	AddQueue(&nexts, start, cmpAStar(scores, score, heuristic))
	var visited Set[T] = DefSet[T]()
	for !IsEmpty(nexts) && !In(visited, end) {
		var next = PopQueue(&nexts)
		if !In(visited, next) {
			Add(visited, next)
			for neighbor, cost := range neighbors(next) {
				var _, ok = scores[neighbor]
				if !ok || cmpAStar(scores, score, heuristic)(next, neighbor) {
					scores[neighbor] = Score{Score: scores[next].Score + cost, Heuristic: score*scores[next].Score + heuristic*funcHeuristic(next)}
					AddQueue(&nexts, neighbor, cmpAStar(scores, score, heuristic))
				}
			}
		}
	}
	return scores
}
