package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
	"time"
)

type Range [2]int

func createRange(a int, b int)Range{
	return Range{a,b}
}

func isValidRange(r Range)bool{
	return r[0] <= r[1]
}

func printRange(r Range){
		println(r[0],r[1])
}

func convertStringToInt(str string) int {
	var res int
	for _, char := range str {
		res = res*10 + int(char-'0')
	}
	return res
}

func convertSliceStringToSliceInt(strSlice []string) []int {
	var res []int
	for _, number := range strSlice {
		res = append(res, convertStringToInt(number))
	}
	return res
}

func getSeeds(str string) []int {
	return convertSliceStringToSliceInt(strings.Split(strings.Split(str, ": ")[1], " "))
}

func getSeeds2(str string) []Range {
	var oldSeeds = getSeeds(str)
	var len = len(oldSeeds)/2
	var res []Range = make([]Range, len)
	for i := range len {
		res[i][0] = oldSeeds[i*2]
		res[i][1] = oldSeeds[i*2] + oldSeeds[i*2+1] -1
	}
	return res
}

func getValues(str string) []int {
	return convertSliceStringToSliceInt(strings.Split(str, " "))
}

func minSlice(slice []int) int {
	var res int = -1
	for _, elem := range slice {
		if res == -1 || elem < res {
			res = elem
		}
	}
	return res
}

func isPresent(char byte, str string) bool {
	for i := range str {
		if str[i] == char {
			return true
		}
	}
	return false
}

func parse(lines []string) [][][3]int {
	var index int = -1
	var res [][][3]int = make([][][3]int, 0)
	for i := 1; i < len(lines); i++ {
		if lines[i] == "" {
			index += 1
			var newTab [][3]int
			res = append(res, newTab)
		} else {
			if !isPresent(':', lines[i]) {
				res[index] = append(res[index], [3]int(getValues(lines[i])))
			}
		}
	}
	return res
}

func calcul(mapRange [3]int, val int) (int, bool) {
	if mapRange[1] <= val && val < mapRange[1]+mapRange[2] {
		return mapRange[0] + val - mapRange[1], true
	}
	return val, false
}

func calculSlice(tabMap [][3]int, val int) int {
	for _, mapRange := range tabMap {
		var newVal, ok = calcul(mapRange, val)
		if ok {
			return newVal
		}
	}
	return val
}

func part1(input string) int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var seeds []int = getSeeds(lines[0])
	var tabRanges [][][3]int = parse(lines)
	for _, tabRange := range tabRanges {
		for j, seed := range seeds {
			seeds[j] = calculSlice(tabRange, seed)
		}
	}
	return minSlice(seeds)
}

func intersectRanges(r1 Range, r2 Range) (Range,bool){
	var res Range = createRange(max(r1[0],r2[0]),min(r1[1],r2[1]))
	return res, isValidRange(res)
}

func minSliceRange(tab []Range, defaut int)int{
	var mini = defaut
	for _,r := range tab{
		mini = min(r[0],mini)
	}
	return mini
}

func maxSliceRange(tab []Range, defaut int)int{
	var maxi = defaut
	for _,r := range tab{
		maxi = max(r[1],maxi)
	}
	return maxi
}


func part2(input string)int {
	var lines = strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var seeds []Range = getSeeds2(lines[0])
	var tabRanges [][][3]int = parse(lines)
	for _, tabRange := range tabRanges {
		var nextSeeds []Range = make([]Range, 0)
		for _, seed := range seeds {
			var newRanges = make([]Range,0)
			var notFixed = make([]Range,0)
			for _,rang := range tabRange{
				var inter,ok = intersectRanges(seed,Range{rang[1],rang[2]+rang[1]-1})
				if (ok){
					notFixed = append(notFixed, inter)
					var newInter = Range{inter[0]-rang[1]+rang[0],inter[1]-rang[1]+rang[0]}
					newRanges = append(newRanges, newInter)
				}
			}
			//Fix point
			nextSeeds = append(nextSeeds, newRanges...)
			var fix1 = createRange(seed[0], minSliceRange(notFixed, seed[1])-1)
			var fix2 = createRange(maxSliceRange(notFixed, seed[0])+1, seed[1])
			if isValidRange(fix1){
				nextSeeds = append(nextSeeds, fix1)
			}
			if isValidRange(fix2){
				nextSeeds = append(nextSeeds, fix2)
			}
		}

		seeds = nextSeeds
	}
	var mini int= -1
	for _,seed := range seeds{
		if mini == -1 {mini = seed[0]}
		mini = min(mini,seed[0])
	}
	return mini
}

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("--2023 day 05 solution--")
	start := time.Now()
	fmt.Println("part1 : ", part1(string(input)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("part2 : ", part2(string(input)))
	fmt.Println(time.Since(start))
}
