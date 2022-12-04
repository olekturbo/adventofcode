package main

import (
	"adventofcode"
	"strings"
)

const (
	day = "3"
)

func main() {
	lines := strings.Split(string(adventofcode.Request(day)), "\n")

	sum := 0

	for i := 0; i < len(lines)-1; i += 3 {
		inter := adventofcode.Intersection([]rune(lines[i]), []rune(lines[i+1]), true)
		inter = adventofcode.Intersection(inter, []rune(lines[i+2]), true)
		for _, v := range inter {
			vString := string(v)
			vInt := int(v)
			if vInt != 0 {
				if strings.ToUpper(vString) == vString {
					sum += vInt - 38
				} else {
					sum += vInt - 96
				}
			}
		}
	}
	println(sum)
}
