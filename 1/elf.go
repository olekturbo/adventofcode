package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

const (
	maxAge       = 300
	session      = "53616c7465645f5f303f959de1a9733fa60cea3722f02a00fb8404cd9d9fa66c0d8e0e4d56c971bb67d0482d452070e01de490fc97d734af44571de48a07ae3d"
	adventOfCode = "https://adventofcode.com/2022/day/1/input"
)

func main() {
	var client http.Client

	cookie := &http.Cookie{
		Name:   "session",
		Value:  session,
		MaxAge: maxAge,
	}

	req, err := http.NewRequest("GET", adventOfCode, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.AddCookie(cookie)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(body), "\n")

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
