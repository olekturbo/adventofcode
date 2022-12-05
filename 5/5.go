package main

import (
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	arr := make([][]rune, 9)
	for i := range arr {
		arr[i] = make([]rune, 0)
	}
	arr[0] = []rune{'R', 'P', 'C', 'D', 'B', 'G'}
	arr[1] = []rune{'H', 'V', 'G'}
	arr[2] = []rune{'N', 'S', 'Q', 'D', 'J', 'P', 'M'}
	arr[3] = []rune{'P', 'S', 'L', 'G', 'D', 'C', 'N', 'M'}
	arr[4] = []rune{'J', 'B', 'N', 'C', 'P', 'F', 'L', 'S'}
	arr[5] = []rune{'Q', 'B', 'D', 'Z', 'V', 'G', 'T', 'S'}
	arr[6] = []rune{'B', 'Z', 'M', 'H', 'F', 'T', 'Q'}
	arr[7] = []rune{'C', 'M', 'D', 'B', 'F'}
	arr[8] = []rune{'F', 'C', 'Q', 'G'}

	re := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)

	lines := strings.Split(string(content), "\n")

	reverse := false

	for _, l := range lines {
		match := re.FindStringSubmatch(l)
		count, _ := strconv.Atoi(match[1])
		from, _ := strconv.Atoi(match[2])
		to, _ := strconv.Atoi(match[3])
		temp := arr[from-1][len(arr[from-1])-count:]
		if reverse {
			for k, j := 0, len(temp)-1; k < j; k, j = k+1, j-1 {
				temp[k], temp[j] = temp[j], temp[k]
			}
		}
		arr[from-1] = arr[from-1][:len(arr[from-1])-count]
		arr[to-1] = append(arr[to-1], temp...)
	}

	for _, a := range arr {
		print(string(a[len(a)-1]))
	}
}
