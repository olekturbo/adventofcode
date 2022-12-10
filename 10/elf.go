package main

import (
	"adventofcode"
	"fmt"
	"strconv"
	"strings"
)

const (
	day  = "10"
	addx = "addx"
	noop = "noop"
)

type instruction struct {
	command string
	value   int
}

func main() {

	lines := strings.Split(string(adventofcode.Request(day)), "\n")

	var instructions []instruction

	for _, l := range lines {
		if len(l) != 0 {
			split := strings.Split(l, " ")
			if len(split) == 1 {
				instructions = append(instructions, instruction{command: noop})
			} else {
				value, _ := strconv.Atoi(split[1])
				instructions = append(instructions, instruction{command: addx, value: value})
			}
		}
	}

	cycles := 0
	x := 1
	var strengths []int
	var crt [6][40]string

	for _, in := range instructions {
		if in.command == noop {
			drawPixel(x, cycles, &crt)
			cycles++
			strengths = tryAddStrength(strengths, cycles, x)
		} else {
			for i := 0; i < 2; i++ {
				drawPixel(x, cycles, &crt)
				cycles++
				strengths = tryAddStrength(strengths, cycles, x)
			}
			x += in.value
		}
	}

	fmt.Println(strengths)
	fmt.Println(adventofcode.Sum(strengths))

	for i := range crt {
		for j := range crt[i] {
			print(crt[i][j])
		}
		println()
	}
}

func tryAddStrength(strengths []int, cycle, strength int) []int {
	var mustCycles = []int{20, 60, 100, 140, 180, 220}

	if adventofcode.Contains(mustCycles, cycle) {
		strengths = append(strengths, strength*cycle)
	}

	return strengths
}

func drawPixel(x int, cycle int, crt *[6][40]string) {
	i, j := cycle/40, cycle%40
	if x == j || x == j+1 || x == j-1 {
		crt[i][j] = "#"
	} else {
		crt[i][j] = "."
	}
}
