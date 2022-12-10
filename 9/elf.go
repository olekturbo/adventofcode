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

	ropeLength = 10
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

	var rope []*point

	for i := 0; i < ropeLength; i++ {
		rope = append(rope, newPoint())
	}

	for _, m := range moves {
		for i := 0; i < m.count; i++ {
			switch m.direction {
			case up:
				rope[0].y++
			case down:
				rope[0].y--
			case right:
				rope[0].x++
			case left:
				rope[0].x--
			}

			for j := 1; j < ropeLength; j++ {
				switch moveRope(rope[j-1], rope[j]) {
				case up:
					rope[j].y++
				case down:
					rope[j].y--
				case right:
					rope[j].x++
				case left:
					rope[j].x--
				case rightUp:
					rope[j].x++
					rope[j].y++
				case rightDown:
					rope[j].x++
					rope[j].y--
				case leftDown:
					rope[j].x--
					rope[j].y--
				case leftUp:
					rope[j].x--
					rope[j].y++
				}
				if j == len(rope)-1 {
					tailPoint := point{x: rope[j].x, y: rope[j].y}
					tailPointsMap[tailPoint] = true
				}
			}
		}
	}
	for _, r := range rope {
		fmt.Println(r)
	}
	fmt.Println(len(tailPointsMap))
}

func moveRope(r1 *point, r2 *point) string {
	if r1.y == r2.y {
		if r1.x == r2.x+2 {
			return right
		}
		if r1.x == r2.x-2 {
			return left
		}
	}
	if r1.x == r2.x {
		if r1.y == r2.y+2 {
			return up
		}
		if r1.y == r2.y-2 {
			return down
		}
	}
	if r1.x == r2.x+2 && r1.y == r2.y+1 || r1.x == r2.x+1 && r1.y == r2.y+2 ||
		r1.x == r2.x+2 && r1.y == r2.y+2 {
		return rightUp
	}
	if r1.x == r2.x+2 && r1.y == r2.y-1 || r1.x == r2.x+1 && r1.y == r2.y-2 ||
		r1.x == r2.x+2 && r1.y == r2.y-2 {
		return rightDown
	}
	if r1.x == r2.x-2 && r1.y == r2.y+1 || r1.x == r2.x-1 && r1.y == r2.y+2 ||
		r1.x == r2.x-2 && r1.y == r2.y+2 {
		return leftUp
	}
	if r1.x == r2.x-2 && r1.y == r2.y-1 || r1.x == r2.x-1 && r1.y == r2.y-2 ||
		r1.x == r2.x-2 && r1.y == r2.y-2 {
		return leftDown
	}
	return ""
}
