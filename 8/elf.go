package main

import (
	"adventofcode"
	"strconv"
	"strings"
)

const (
	day = "8"
)

func main() {
	lines := strings.Split(string(adventofcode.Request(day)), "\n")

	trees := make([][]int, len(lines)-1)

	for i, l := range lines {
		if len(l) != 0 {
			trees[i] = make([]int, len(lines)-1)
			for j, c := range l {
				a, _ := strconv.Atoi(string(c))
				trees[i][j] = a
			}
		}
	}

	var sum int

	for i, v := range trees {
		for j := range v {
			upper := make([]int, 0)
			lower := make([]int, 0)

			if i == 0 {
				sum++
				continue
			}
			if i == len(trees)-1 {
				sum++
				continue
			}
			if j == 0 {
				sum++
				continue
			}
			if j == len(v)-1 {
				sum++
				continue
			}
			rightMax := adventofcode.Max(v[j+1:])
			if trees[i][j] > rightMax {
				sum++
				continue
			}
			leftMax := adventofcode.Max(v[:j])
			if trees[i][j] > leftMax {
				sum++
				continue
			}
			for k := 0; k < i; k++ {
				upper = append(upper, trees[k][j])
			}
			for k := i + 1; k < len(trees); k++ {
				lower = append(lower, trees[k][j])
			}
			upperMax := adventofcode.Max(upper)
			if trees[i][j] > upperMax {
				sum++
				continue
			}
			lowerMax := adventofcode.Max(lower)
			if trees[i][j] > lowerMax {
				sum++
				continue
			}
		}
	}

	println(sum)
}
