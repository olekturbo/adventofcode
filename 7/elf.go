package main

import (
	"adventofcode"
	"strconv"
	"strings"
	"unicode"
)

const (
	day         = "7"
	totalSize   = 100000
	systemSize  = 70000000
	minimumSize = 30000000
)

type dir struct {
	dirs   []*dir
	parent *dir
	name   string
	size   int
}

var root *dir
var current *dir

func main() {
	lines := strings.Split(string(adventofcode.Request(day)), "\n")
	root = &dir{name: "root"}

	for _, l := range lines {
		if len(l) != 0 {
			switch {
			case strings.HasPrefix(l, "$ cd"):
				d := strings.Replace(l, "$ cd ", "", 1)
				switch d {
				case "/":
					current = root
				case "..":
					current = current.parent
				default:
					current = current.findDir(d)
				}
			case l == "$ ls":
				continue
			case strings.HasPrefix(l, "dir"):
				name := strings.Replace(l, "dir ", "", 1)
				newDir := &dir{parent: current, name: name}
				current.dirs = append(current.dirs, newDir)
			case unicode.IsDigit(rune(l[0])):
				split := strings.Split(l, " ")
				size, _ := strconv.Atoi(split[0])
				current.incSize(size)
			}
		}
	}

	println(root.countTotalSize())

	needSize := minimumSize - (systemSize - root.size)
	println(adventofcode.Min(root.findNeededSizes(needSize)))
}

func (d *dir) incSize(size int) {
	d.size += size
	if d.parent != nil {
		d.parent.incSize(size)
	}
}

func (d *dir) countTotalSize() int {
	var total int

	if d.size <= totalSize {
		total += d.size
	}

	for _, dd := range d.dirs {
		total += dd.countTotalSize()
	}

	return total
}

func (d *dir) findNeededSizes(needSize int) []int {
	var szs []int

	if d.size >= needSize {
		szs = append(szs, d.size)
	}

	for _, dd := range d.dirs {
		szs = append(szs, dd.findNeededSizes(needSize)...)
	}

	return szs
}

func (d *dir) findDir(name string) *dir {
	for _, dd := range d.dirs {
		if dd.name == name {
			return dd
		}
	}
	return nil
}
