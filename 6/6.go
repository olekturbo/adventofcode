package main

import (
	"adventofcode"
	"strings"
)

const (
	day      = "6"
	distinct = 14
)

func main() {
	req := adventofcode.Request(day)

	content := strings.TrimSpace(string(req))

	for i := range content {
		m := make(map[rune]int, 0)
		x := content[i : distinct+i]
		for _, xx := range x {
			m[xx]++
		}
		if len(m) == distinct {
			print(i + distinct)
			break
		}
	}
}
