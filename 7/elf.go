package main

import (
	"adventofcode"
	"fmt"
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
	files  []*file
	parent *dir
	name   string
	size   int
}

type file struct {
	name string
	size int
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
				newFile := &file{name: split[1], size: size}
				current.files = append(current.files, newFile)
			}
		}
	}

	println(root.countTotalSize())
	println(adventofcode.Min(root.findNeededSizes()))

	root.print(0)
}

func (d *dir) incSize(size int) {
	d.size += size
	if d.parent != nil {
		d.parent.incSize(size)
	}
}

func (d *dir) print(i int) {
	fmt.Print(strings.Repeat("  ", i))
	fmt.Printf("%s DIR %d\n", d.name, d.size)
	for _, f := range d.files {
		fmt.Print(strings.Repeat("  ", i+1))
		fmt.Printf("%s %d\n", f.name, f.size)
	}

	for _, dd := range d.dirs {
		dd.print(i + 1)
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

func (d *dir) findNeededSizes() []int {
	var szs []int

	if d.size >= minimumSize-(systemSize-root.size) {
		szs = append(szs, d.size)
	}

	for _, dd := range d.dirs {
		szs = append(szs, dd.findNeededSizes()...)
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
