package main

import (
	_ "embed"
	"fmt"
	"math"
	"math/rand"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	. "github.com/Discobluff/advent-of-code/go/utils/set"
)

type Connection struct {
	wire1, wire2, output, gate string
}

func and(a int, b int) int {
	if a == 1 && b == 1 {
		return 1
	}
	return 0
}

func or(a int, b int) int {
	if a == 1 || b == 1 {
		return 1
	}
	return 0
}

func xor(a int, b int) int {
	if a != b {
		return 1
	}
	return 0
}

func parseWire(lines []string) map[string]int {
	var res map[string]int = make(map[string]int)
	for _, line := range lines {
		var split = strings.Split(line, ": ")
		var number, _ = strconv.Atoi(split[1])
		res[split[0]] = number
	}
	return res
}

func parseConnection(lines []string) Set[Connection] {
	var res = DefSet[Connection]()
	for _, line := range lines {
		var split = strings.Split(line, " ")
		Add(res, Connection{wire1: split[0], wire2: split[2], output: split[4], gate: split[1]})
	}
	return res
}

func binaryToDecimal(binary []int) int {
	decimal := 0
	for i, bit := range binary {
		if bit == 1 {
			decimal += int(math.Pow(2, float64(i)))
		}
	}
	return decimal
}
func decimalToBinary(decimal int) []int {
	if decimal == 0 {
		return []int{0}
	}
	var binary []int
	for decimal > 0 {
		bit := decimal % 2
		// binary = append([]int{bit}, binary...)
		binary = append(binary, bit)
		decimal = decimal / 2
	}
	return binary
}

func selectByPrefix(m map[string]int, prefix byte) []int {
	var count int
	for wire := range m {
		if wire[0] == prefix {
			count++
		}
	}
	var res []int = make([]int, count)
	for wire, val := range m {
		if wire[0] == prefix {
			var index, _ = strconv.Atoi(wire[1:])
			res[index] = val
		}
	}
	return res
}

func solve(wires map[string]int, connections Set[Connection]) []int {
	for !IsEmpty(connections) {
		for connection := range connections {
			var _, ok1 = wires[connection.wire1]
			var _, ok2 = wires[connection.wire2]
			if ok1 && ok2 {
				if connection.gate == "OR" {
					wires[connection.output] = or(wires[connection.wire1], wires[connection.wire2])
				}
				if connection.gate == "AND" {
					wires[connection.output] = and(wires[connection.wire1], wires[connection.wire2])
				}
				if connection.gate == "XOR" {
					wires[connection.output] = xor(wires[connection.wire1], wires[connection.wire2])
				}
				Remove(connections, connection)
			}
		}
	}
	return selectByPrefix(wires, 'z')
}

func part1(input string) int {
	var lines = strings.TrimSuffix(input, "\n")
	var split = strings.Split(lines, "\n\n")
	var wires = parseWire(strings.Split(split[0], "\n"))
	var connections = parseConnection(strings.Split(split[1], "\n"))
	return binaryToDecimal(solve(wires, connections))
}

func differenceBits(x []int, y []int) []int {
	var res []int
	for i := range x {
		if i >= len(y) {
			y = append(y, 0)
		}
		if x[i] != y[i] {
			res = append(res, i)
		}
	}
	return res
}

func appendstr(a []byte, s string) []byte {
	for _, char := range s {
		a = append(a, byte(char))
	}
	return a
}

// [8 9 15 16 17 18 22 23 35 36 37]
func swap(s Set[Connection], o1 string, o2 string) Set[Connection] {
	var res = DefSet[Connection]()
	for connec := range s {
		var c Connection = connec
		if c.output == o1 {
			c.output = o2
		} else {
			if c.output == o2 {
				c.output = o1
			}
		}
		Add(res, c)
	}
	return res
}

//Algo de résolution "à la main"
//Lancer le programme sur 10 entrées aléatoires (les x et les y)
//Identifer les endroits qui posent problèmes
//Tenter des échanges qui résolvent les problèmes
//On essaie pas les échanges : qui échange deux bits de la même valeur, si les deux noeuds dépendent l'un de l'autre
//Problèmes : 8,9 - 15,16,17 - 22,23,24,25,26,27 - 35,36,37,38,39,40,41
// gvw <-> qjb règle 8 et 9
// z15 <-> jgc (pas rgt, fbv, rfj, le z15 doit obligatoirement être modifié car il dépend directement d'un x et d'un y) règle 15,16,17
// Essais : z22 <-> drg (premier coup) règle 22 à 27
// Essais : z35 <-> jbp (dtj <-> vcs, dtj <-> ppf, vcs <-> ppf, ppf <-> qrg, qrg <-> jbp, qrg <-> grd, jbp <-> grd, rbm <-> sck, dtj <-> jms, vcs <-> jms, rbm <-> jms  marche pas) règle 35 à 41

// Résultat : drg,gvw,jbp,jgc,qjb,z15,z22,z35
func part2(input string) {
	var lines = strings.TrimSuffix(input, "\n")
	var split = strings.Split(lines, "\n\n")
	var wires = parseWire(strings.Split(split[0], "\n"))
	for wire := range wires {
		wires[wire] = rand.Intn(2)
	}
	var connections = parseConnection(strings.Split(split[1], "\n"))
	var x = selectByPrefix(wires, 'x')
	var y = selectByPrefix(wires, 'y')
	var z = solve(wires, Union(connections, DefSet[Connection]()))
	var expected = decimalToBinary(binaryToDecimal(x) + binaryToDecimal(y))
	var wrongBits []int = differenceBits(expected, z)
	fmt.Println(wrongBits)
}

func createDot(input string) {
	var lines = strings.TrimSuffix(input, "\n")
	var split = strings.Split(lines, "\n\n")
	var wires = parseWire(strings.Split(split[0], "\n"))
	var connections = parseConnection(strings.Split(split[1], "\n"))
	var x = selectByPrefix(wires, 'x')
	var y = selectByPrefix(wires, 'y')
	var z = solve(wires, connections)
	var expected = decimalToBinary(binaryToDecimal(x) + binaryToDecimal(y))
	var wrongBits []int = differenceBits(expected, z)
	for _, bit := range wrongBits {
		fmt.Printf("%d", z[bit])
	}
	fmt.Printf("\n")
	var dot []byte = appendstr([]byte{}, "digraph LogicCircuit {\n")
	for wire, val := range wires {
		dot = appendstr(dot, wire)
		dot = appendstr(dot, " [label=")
		dot = append(dot, '"')
		dot = appendstr(dot, wire)
		dot = appendstr(dot, "=")
		if val == 0 {
			dot = append(dot, '0')
		} else {
			dot = append(dot, '1')
		}
		dot = append(dot, '"')
		if wire[0] == 'z' {
			var color = "green"
			dot = appendstr(dot, ", style=filled, fillcolor=")
			var index, _ = strconv.Atoi(wire[1:])
			if slices.Contains(wrongBits, index) {
				color = "red"
			}
			dot = appendstr(dot, color)
		}
		dot = appendstr(dot, "];\n")
	}
	connections = parseConnection(strings.Split(split[1], "\n"))
	for connection := range connections {
		dot = appendstr(dot, connection.wire1)
		dot = appendstr(dot, " -> ")
		dot = appendstr(dot, connection.gate+connection.output)
		dot = appendstr(dot, ";\n")
		dot = appendstr(dot, connection.wire2)
		dot = appendstr(dot, " -> ")
		dot = appendstr(dot, connection.gate+connection.output)
		dot = appendstr(dot, ";\n")
		dot = appendstr(dot, connection.gate+connection.output)
		dot = appendstr(dot, " [label=")
		dot = append(dot, '"')
		dot = appendstr(dot, connection.gate)
		dot = append(dot, '"')
		dot = appendstr(dot, ", style=filled, fillcolor=gold];\n")
		dot = appendstr(dot, connection.gate+connection.output)
		dot = appendstr(dot, " -> ")
		dot = appendstr(dot, connection.output)
		dot = appendstr(dot, ";\n")

	}
	dot = append(dot, '}')
	err := os.WriteFile("output.dot", dot, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("--2024 day 24 solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(string(input)))
	fmt.Println(time.Since(start))
	for range 10 {
		part2(string(input))
	}
	// createDot(input)
}
