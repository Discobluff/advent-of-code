package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var inputDay string

func isSymbol(char byte) bool {
	return (char < 48 || char > 58) && char != 46
}

func isStar(char byte) bool {
	return char == 42
}

func isDigit(char byte) bool {
	return char >= 48 && char <= 58
}

func isPoint(char byte) bool {
	return char == 46
}

func checkRightNumber1(lines []string, deb int, j int, i int) bool {
	//Check on the left
	if deb > 0 && isSymbol(lines[i][deb-1]) {
		return true
	}
	//Check on the right
	if j < len(lines[i]) && isSymbol(lines[i][j]) {
		return true
	}
	//Check the line above
	if i > 0 {
		for k := max(0, deb-1); k <= min(j, len(lines[i])-1); k++ {
			if isSymbol(lines[i-1][k]) {
				return true
			}
		}
	}
	//Check the line below
	if i < len(lines)-1 {
		for k := max(0, deb-1); k <= min(j, len(lines[i])-1); k++ {
			if isSymbol(lines[i+1][k]) {
				return true
			}
		}
	}
	return false
}

func part1(lines []string) int {
	var sum int
	for i, line := range lines {
		//Find the numbers in the line
		var number int
		var deb int
		for j, char := range line {
			if isDigit(line[j]) {
				if number == 0 {
					deb = j
				}
				number = 10*number + int(char-48)
				if j == len(line)-1 {
					if checkRightNumber1(lines, deb, j+1, i) {
						sum += number
					}
				}
			} else {
				if number != 0 {
					//We found a number, we must check it is adjacent to symbol
					if checkRightNumber1(lines, deb, j, i) {
						sum += number
					}
					number = 0
					deb = 0
				}
			}
		}
	}
	return sum
}

func getNumberByStart(line string, index int) int {
	var number int
	for i := index; i <= len(line)-1 && isDigit(line[i]); i++ {
		number = 10*number + int(line[i]-48)
	}
	return number
}

func getNumberNotByStart(line string, index int) int {
	var start int
	for i := index; i >= 0 && isDigit(line[i]); i-- {
		start = i
	}
	return getNumberByStart(line, start)
}

func part2(lines []string) int {
	var sum int
	for i, line := range lines {
		for j := range line {
			if isStar(line[j]) {
				var number1 int
				var number2 int
				// Nombre à gauche
				if j-1 >= 0 {
					if isDigit(line[j-1]) {
						number1 = getNumberNotByStart(line, j-1)
					}
				}
				//Nombre à droite
				if j+1 <= len(line)-1 {
					if isDigit(line[j+1]) {
						number2 = getNumberByStart(line, j+1)
					}
				}
				// Cas X.X
				//		.*.
				if i-1 >= 0 && j-1 >= 0 && j+1 <= len(line)-1 {
					if isDigit(lines[i-1][j-1]) && isPoint(lines[i-1][j]) && isDigit(lines[i-1][j+1]) {
						number1 = getNumberNotByStart(lines[i-1], j-1)
						number2 = getNumberByStart(lines[i-1], j+1)
					}
				}
				// Cas .*.
				//	   X.X
				if i+1 <= len(lines)-1 && j-1 >= 0 && j+1 <= len(line)-1 {
					if isDigit(lines[i+1][j-1]) && isPoint(lines[i+1][j]) && isDigit(lines[i+1][j+1]) {
						number1 = getNumberNotByStart(lines[i+1], j-1)
						number2 = getNumberByStart(lines[i+1], j+1)
					}
				}
				// Nombre au dessus
				if i-1 >= 0 {
					var index int = -1
					if j-1 >= 0 && isDigit(lines[i-1][j-1]) {
						index = j - 1
					}
					if isDigit(lines[i-1][j]) {
						index = j
					}
					if j+1 <= len(line)-1 && isDigit(lines[i-1][j+1]) {
						index = j + 1
					}
					if index != -1 {
						if number1 == 0 {
							number1 = getNumberNotByStart(lines[i-1], index)
						} else {
							if number2 == 0 {
								number2 = getNumberNotByStart(lines[i-1], index)
							}
						}
					}
				}
				//Nombre en dessous
				if i+1 <= len(lines)-1 {
					var index int = -1
					if j-1 >= 0 && isDigit(lines[i+1][j-1]) {
						index = j - 1
					}
					if isDigit(lines[i+1][j]) {
						index = j
					}
					if j+1 <= len(line)-1 && isDigit(lines[i+1][j+1]) {
						index = j + 1
					}
					if index != -1 {
						if number1 == 0 {
							number1 = getNumberNotByStart(lines[i+1], index)
						} else {
							if number2 == 0 {
								number2 = getNumberNotByStart(lines[i+1], index)
							}
						}
					}
				}
				if number1 != 0 && number2 != 0 {
					sum += number1 * number2
				}
			}
		}
	}
	return sum
}

func main() {
	var input = strings.TrimSuffix(inputDay, "\n")
	var lines = strings.Split(input, "\n")
	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}
