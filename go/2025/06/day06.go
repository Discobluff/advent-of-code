package main

import (
	_ "embed"
	"fmt"
	"os"
	// "slices"
	"strconv"
	"strings"
	"time"
)

func printTab(tab []int){
	for _,val := range tab{
		println(val)
	}
}

func printTabString(tab []string){
	for _,val := range tab{
		println(val)
	}
}

func part1(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var calculs []int = make([]int, 0)
	var firstLine = strings.Split(lines[0], " ")
	for _,number := range firstLine{
		if number != ""{
			var numberInt, _ = strconv.Atoi(number)
			calculs = append(calculs, numberInt)
		}
	}
	// printTab(calculs)
	var ops []byte = make([]byte, len(calculs))
	var opsLine = strings.Split(lines[len(lines)-1], " ")
	var index int
	for _, op := range opsLine{
		if op != ""{
			ops[index] = op[0]
			index++
		}
	}
	var secondLine = strings.Split(lines[1], " ")
	var thirdLine = strings.Split(lines[2], " ")
	// var fourthLine = strings.Split(lines[3], " ")
	index = 0
	for _,number := range secondLine{
		if number != ""{
			var numberInt, _ = strconv.Atoi(number)
			if ops[index] == '+'{
				calculs[index] += numberInt
			} else{
				calculs[index] *= numberInt
			}
			index++
		}
	}
	// printTab(calculs)
	index = 0
	for _,number := range thirdLine{
		if number != ""{
			var numberInt, _ = strconv.Atoi(number)
			if ops[index] == '+'{
				calculs[index] += numberInt
			} else{
				calculs[index] *= numberInt
			}
			index++
		}
	}
	// printTab(calculs)
	// index = 0
	// for _,number := range fourthLine{
	// 	if number != ""{
	// 		var numberInt, _ = strconv.Atoi(number)
	// 		if ops[index] == '+'{
	// 			calculs[index] += numberInt
	// 		} else{
	// 			calculs[index] *= numberInt
	// 		}
	// 		index++
	// 	}
	// }
	var res int
	for _,val := range calculs{
		res+=val
	}
	return res
}

func deleteSlice(s []string)[]string{
	var index int
	for _,val := range s{
		if val != ""{
			index++
		}
	}
	var res []string = make([]string, index)
	index = 0
	for _,val := range s{
		if val != ""{
			res[index] = val
			index++
		}
	}
	return res
}

func split(line string, indexes []int)[]string{
	var res []string = make([]string, len(indexes)+1)
	start := 0
	for i, idx := range indexes {
		res[i] = line[start:idx]
		start = idx+1
	}
	res[len(indexes)] = line[start:]
	return res
}

func part2(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var parseLines [][]string = make([][]string, len(lines)-1)
	// var indexes []int = []int{3, 7, 11}
	var indexes []int = make([]int, 0)
	for i := range len(lines[0]){
		var add = true
		for j := range len(lines)-1{
			if lines[j][i] != ' '{
				add = false
			}
		}
		if add{
			indexes = append(indexes, i)
		}
	}
	for i := range len(lines)-1{
		parseLines[i] = split(lines[i], indexes)
	}
	// printTabString(parseLines[1])
	var ops []byte = make([]byte, 0)
	var opsLine = strings.Split(lines[len(lines)-1], " ")
	// var index int
	for _, op := range opsLine{
		if op != ""{
			ops = append(ops, op[0])
			// index++
		}
	}
	// printTabString(parseLines[0])
	// var index = 0
	var res int
	for i := range parseLines[0]{
		if ops[i] == '+'{
			var val =0
			for nbNumb := range parseLines[0][i]{
				var nb int
				for j := range parseLines{
					if parseLines[j][i][nbNumb] == ' '{
						continue
					}else{
						nb = 10*nb + int(parseLines[j][i][nbNumb]-'0')
					}
				}
				val += nb
			}
			res+= val
		}else{
			var val =1
			for nbNumb := range parseLines[0][i]{
				var nb int
				for j := range parseLines{
					if parseLines[j][i][nbNumb] == ' '{
						continue
					}else{
						nb = 10*nb + int(parseLines[j][i][nbNumb]-'0')
					}
				}
				val *= nb
			}
			res+= val
		}
	}
	
	return res
}

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("--2025 day 06 solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(string(input)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("part2 : ", part2(string(input)))
	fmt.Println(time.Since(start))
}
