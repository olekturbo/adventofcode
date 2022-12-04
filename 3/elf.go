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

	lbp := make([][]rune, len(lines)-1)
	rbp := make([][]rune, len(lines)-1)

	for i, l := range lines {
		if len(l) != 0 {
			lbp[i] = make([]rune, 52)
			rbp[i] = make([]rune, 52)
			for j, ll := range l[:len(l)/2] {
				lbp[i][j] = ll
			}
			for j, ll := range l[len(l)/2:] {
				rbp[i][j] = ll
			}
		}
	}

	sum := 0

	for i := 0; i < len(lines)-1; i++ {
		for _, v := range adventofcode.Intersection(lbp[i], rbp[i], true) {
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
