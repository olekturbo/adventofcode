package main

import (
	"adventofcode"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const (
	day = "1"
)

func main() {

	lines := strings.Split(string(adventofcode.Request(day)), "\n")

	elves := make([]int, 0)
	count := 0

	for _, l := range lines {
		if len(l) == 0 {
			elves = append(elves, count)
			count = 0
		} else {
			ii, _ := strconv.Atoi(l)
			count += ii
		}
	}

	sort.Slice(elves, func(i, j int) bool {
		return elves[i] > elves[j]
	})

	topThree := 0

	for i := 0; i < 3; i++ {
		topThree += elves[i]
	}

	fmt.Println(elves[0])
	fmt.Println(topThree)
}
