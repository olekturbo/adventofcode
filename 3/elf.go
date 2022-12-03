package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	maxAge       = 300
	session      = "53616c7465645f5f303f959de1a9733fa60cea3722f02a00fb8404cd9d9fa66c0d8e0e4d56c971bb67d0482d452070e01de490fc97d734af44571de48a07ae3d"
	adventOfCode = "https://adventofcode.com/2022/day/3/input"
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

	lbp := make([][]rune, len(lines)-1)
	rbp := make([][]rune, len(lines)-1)

	for i, l := range lines {
		if len(l) != 0 {
			lbp[i] = make([]rune, 52)
			rbp[i] = make([]rune, 52)
			for j, ll := range l[:len(l)/2] {
				lbp[i][j] = ll
			}
			for j, ll := range l[len(l)/2:] {
				rbp[i][j] = ll
			}
		}
	}

	sum := 0

	for i := 0; i < len(lines)-1; i++ {
		for _, v := range intersection(lbp[i], rbp[i]) {
			vString := string(v)
			vInt := int(v)
			if vInt != 0 {
				if strings.ToUpper(vString) == vString {
					sum += vInt - 38
				} else {
					sum += vInt - 96
				}
			}
		}
	}

	println(sum)

}

func intersection(s1, s2 []rune) []rune {
	var inter []rune
	h := make(map[rune]bool)
	for _, e := range s1 {
		h[e] = true
	}
	for _, e := range s2 {
		if h[e] {
			inter = append(inter, e)
		}
	}
	return removeDups(inter)
}

func removeDups(elements []rune) []rune {
	var ret []rune
	m := make(map[rune]bool)
	for _, e := range elements {
		if !m[e] {
			ret = append(ret, e)
			m[e] = true
		}
	}
	return ret
}
