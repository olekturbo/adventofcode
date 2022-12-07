package adventofcode

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	maxAge       = 300
	session      = "53616c7465645f5f303f959de1a9733fa60cea3722f02a00fb8404cd9d9fa66c0d8e0e4d56c971bb67d0482d452070e01de490fc97d734af44571de48a07ae3d"
	adventOfCode = "https://adventofcode.com/2022/day/{DAY}/input"
)

func Request(day string) []byte {
	var client http.Client

	cookie := &http.Cookie{
		Name:   "session",
		Value:  session,
		MaxAge: maxAge,
	}

	url := strings.Replace(adventOfCode, "{DAY}", day, 1)

	req, err := http.NewRequest("GET", url, nil)
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

	resp.Body.Close()

	return body
}

func Intersection[v int | rune](s1, s2 []v, dups bool) []v {
	var inter []v
	h := make(map[v]bool)
	for _, e := range s1 {
		h[e] = true
	}
	for _, e := range s2 {
		if h[e] {
			inter = append(inter, e)
		}
	}
	if dups {
		inter = removeDups(inter)
	}
	return inter
}

func removeDups[v int | rune](elements []v) []v {
	var ret []v
	m := make(map[v]bool)
	for _, e := range elements {
		if !m[e] {
			ret = append(ret, e)
			m[e] = true
		}
	}
	return ret
}

func Min(arr []int) int {
	var min = arr[0]

	for _, v := range arr {
		if min > v {
			min = v
		}
	}

	return min
}
