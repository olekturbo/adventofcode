package main

import (
	"adventofcode"
	"fmt"
	"strconv"
	"strings"
)

const (
	day = "9"

	left      = "L"
	right     = "R"
	up        = "U"
	down      = "D"
	leftUp    = "LU"
	leftDown  = "LD"
	rightUp   = "RU"
	rightDown = "RD"
)

type point struct {
	x, y int
}

func newPoint() *point {
	return &point{}
}

type move struct {
	count     int
	direction string
}

func main() {
	lines := strings.Split(string(adventofcode.Request(day)), "\n")

	var moves []move

	tailPointsMap := make(map[point]bool, 0)

	for _, l := range lines {
		if len(l) != 0 {
			split := strings.Split(l, " ")
			count, _ := strconv.Atoi(split[1])
			moves = append(moves, move{direction: split[0], count: count})
		}
	}

	head := newPoint()
	tail := newPoint()

	for _, m := range moves {
		for i := 0; i < m.count; i++ {
			switch m.direction {
			case up:
				head.y++
			case down:
				head.y--
			case right:
				head.x++
			case left:
				head.x--
			}
			switch moveTail(head, tail) {
			case up:
				tail.y++
			case down:
				tail.y--
			case right:
				tail.x++
			case left:
				tail.x--
			case rightUp:
				tail.x++
				tail.y++
			case rightDown:
				tail.x++
				tail.y--
			case leftDown:
				tail.x--
				tail.y--
			case leftUp:
				tail.x--
				tail.y++
			}

			tailPoint := point{x: tail.x, y: tail.y}
			tailPointsMap[tailPoint] = true
		}
	}
	fmt.Println(head)
	fmt.Println(tail)
	fmt.Println(len(tailPointsMap))
}

func moveTail(head *point, tail *point) string {
	if head.y == tail.y {
		if head.x == tail.x+2 {
			return right
		}
		if head.x == tail.x-2 {
			return left
		}
	}
	if head.x == tail.x {
		if head.y == tail.y+2 {
			return up
		}
		if head.y == tail.y-2 {
			return down
		}
	}
	if head.x == tail.x+2 && head.y == tail.y+1 || head.x == tail.x+1 && head.y == tail.y+2 ||
		head.x == tail.x+2 && head.y == tail.y+2 {
		return rightUp
	}
	if head.x == tail.x+2 && head.y == tail.y-1 || head.x == tail.x+1 && head.y == tail.y-2 ||
		head.x == tail.x+2 && head.y == tail.y-2 {
		return rightDown
	}
	if head.x == tail.x-2 && head.y == tail.y+1 || head.x == tail.x-1 && head.y == tail.y+2 ||
		head.x == tail.x-2 && head.y == tail.y+2 {
		return leftUp
	}
	if head.x == tail.x-2 && head.y == tail.y-1 || head.x == tail.x-1 && head.y == tail.y-2 ||
		head.x == tail.x-2 && head.y == tail.y-2 {
		return leftDown
	}
	return ""
}
