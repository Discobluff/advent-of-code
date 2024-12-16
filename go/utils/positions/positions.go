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

func OpposedDirection(pos Position) Position {
	return Position{Line: -pos.Line, Column: -pos.Column}
}

func AddPositions(pos1 Position, pos2 Position) Position {
	return Position{Line: pos1.Line + pos2.Line, Column: pos1.Column + pos2.Column}
}

func SearchStartGrid(grid [][]byte, start byte) Position {
	for i, line := range grid {
		for j, char := range line {
			if char == start {
				return Position{Line: i, Column: j}
			}
		}
	}
	return Position{Line: -1, Column: -1}
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
