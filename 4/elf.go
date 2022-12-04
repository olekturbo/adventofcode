package main

import (
	"adventofcode"
	"strconv"
	"strings"
)

const day = "4"

func main() {
	lines := strings.Split(string(adventofcode.Request(day)), "\n")

	sum := 0
	sumB := 0
	for _, l := range lines {
		if len(l) != 0 {
			ll := strings.Split(l, ",")
			leftRange := parse(strings.Split(ll[0], "-"))
			rightRange := parse(strings.Split(ll[1], "-"))
			inter := len(adventofcode.Intersection(leftRange, rightRange, false))
			if inter == len(leftRange) || inter == len(rightRange) {
				sum++
			}
			if inter > 0 {
				sumB++
			}
		}
	}

	println(sum)
	println(sumB)
}

func parse(arr []string) []int {
	l, _ := strconv.Atoi(arr[0])
	r, _ := strconv.Atoi(arr[1])

	var ret []int
	for i := l; i <= r; i++ {
		ret = append(ret, i)
	}

	return ret
}
