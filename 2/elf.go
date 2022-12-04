package main

import (
	"adventofcode"
	"fmt"
	"strings"
)

const (
	day = "2"
)

func main() {
	lines := strings.Split(string(adventofcode.Request(day)), "\n")

	sum := 0

	for _, l := range lines {
		if len(l) != 0 {
			op := string(l[0])
			my := string(l[2])

			if my == "X" {
				sum += 1
				if op == "A" {
					sum += 3
				} else if op == "C" {
					sum += 6
				}
			} else if my == "Y" {
				sum += 2
				if op == "B" {
					sum += 3
				} else if op == "A" {
					sum += 6
				}
			} else if my == "Z" {
				sum += 3
				if op == "C" {
					sum += 3
				} else if op == "B" {
					sum += 6
				}
			}
		}
	}

	fmt.Println(sum)

	ssum := 0

	for _, l := range lines {
		if len(l) != 0 {
			op := string(l[0])
			my := string(l[2])

			if my == "X" {
				if op == "A" {
					ssum += 3
				} else if op == "B" {
					ssum += 1
				} else if op == "C" {
					ssum += 2
				}
			} else if my == "Y" {
				ssum += 3
				if op == "A" {
					ssum += 1
				} else if op == "B" {
					ssum += 2
				} else if op == "C" {
					ssum += 3
				}
			} else if my == "Z" {
				ssum += 6
				if op == "A" {
					ssum += 2
				} else if op == "B" {
					ssum += 3
				} else if op == "C" {
					ssum += 1
				}
			}
		}
	}

	fmt.Println(ssum)
}
