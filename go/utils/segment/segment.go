package segment

import (
	"fmt"
	"github.com/Discobluff/advent-of-code/go/utils/set"
)

type segment struct {
	a int
	b int
}

func IsValid(s segment)bool{
	return s.a<=s.b
}

func InSegment(s segment, value int)bool{
	return s.a <= value && value <= s.b
}

func InSetSegment(s set.Set[segment], value int)bool{
	for seg := range s{
		if InSegment(seg, value){
			return true
		}
	}
	return false
}

func DefSegment(a int, b int)segment{
	return segment{a, b}
}

func ParseSegmentTiret(s string)segment{
	var res segment
	fmt.Sscanf(s, "%d-%d", &res.a, &res.b)
	return res
}

func ParseLinesSegment(lines []string, parseLine func(string) segment) []segment {
	var res []segment = make([]segment, len(lines))
	for i,line := range lines{
		res[i] = parseLine(line)
	}
	return res
}

func ParseLinesToSetSegment(lines []string, parseLine func(string) segment) set.Set[segment] {
	var res set.Set[segment] = set.DefSet[segment]()
	for _,line := range lines{
		set.Add(res, parseLine(line))
	}
	return res
}

func Size(s segment)int{
	return s.b - s.a + 1
}

//Return if s2 c s1
func Include(s1 segment, s2 segment)bool{
	return s1.a <= s2.a && s2.b <= s1.b
}

func Merge(s1 segment, s2 segment)(segment, bool){
	if s1.a <= s2.a && s2.a <= s1.b && s1.b <= s2.b{
		return DefSegment(s1.a, s2.b),true
	}
	if s2.a <= s1.a && s1.a <= s2.b && s2.b <= s1.b{
		return DefSegment(s2.a, s1.b),true
	}
	return segment{}, false
}

func MergeSetSegment(s set.Set[segment]){
	var change = true
	for change{
		change = false
		var needBreak = false
		for segment1 := range s{
			if !needBreak{
				for segment2 := range s{
					if segment1.a != segment2.a || segment1.b != segment2.b{
						if Include(segment1, segment2){
							set.Remove(s, segment2)
							change = true
						} else {
							var newSegment, ok = Merge(segment1, segment2)
							if ok{
								set.Remove(s, segment1)
							set.Remove(s, segment2)
							set.Add(s, newSegment)
							change = true
							needBreak = true
							break
							}
						}
					}
				}
			}
		}
	}
}
