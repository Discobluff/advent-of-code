package positions

type Position struct {
	line, column int
}

var N = Position{line: -1, column: 0}
var S = Position{line: 1, column: 0}
var E = Position{line: 0, column: 1}
var W = Position{line: 0, column: -1}
var directions = map[byte]Position{'<': W, '>': E, '^': N, 'v': S}

func addPositions(pos1 Position, pos2 Position) Position {
	return Position{line: pos1.line + pos2.line, column: pos1.column + pos2.column}
}
