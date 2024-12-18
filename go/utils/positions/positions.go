package positions

type Position struct {
	Line, Column int
}

var N = Position{Line: -1, Column: 0}
var S = Position{Line: 1, Column: 0}
var E = Position{Line: 0, Column: 1}
var W = Position{Line: 0, Column: -1}
var Directions = map[byte]Position{'<': W, '>': E, '^': N, 'v': S}
var DirectionsSlice = []Position{N, S, W, E}

func DefPosition(line int, column int) Position {
	return Position{Line: line, Column: column}
}

func OpposedDirection(pos Position) Position {
	return DefPosition(-pos.Line, -pos.Column)
}

func AddPositions(pos1 Position, pos2 Position) Position {
	return DefPosition(pos1.Line+pos2.Line, pos1.Column+pos2.Column)
}

func SearchStartGrid(grid [][]byte, start byte) Position {
	for i, line := range grid {
		for j, char := range line {
			if char == start {
				return DefPosition(i, j)
			}
		}
	}
	return DefPosition(-1, -1)
}

func SearchStartLines(grid []string, start rune) Position {
	for i, line := range grid {
		for j, char := range line {
			if char == start {
				return Position{Line: i, Column: j}
			}
		}
	}
	return Position{Line: -1, Column: -1}
}

func SetScore(pos Position, val int, scores *[][]int) {
	(*scores)[pos.Line][pos.Column] = val
}

func GetScore(pos Position, scores [][]int) int {
	return scores[pos.Line][pos.Column]
}

func DefaultBest(pos1 Position, pos2 Position, scores *[][]int) {
	if GetScore(pos1, *scores) == -1 {
		SetScore(pos1, GetScore(pos2, *scores)+1, scores)
	} else {
		SetScore(pos1, min(GetScore(pos1, *scores), GetScore(pos2, *scores)+1), scores)
	}
}
