package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	maxAge       = 300
	session      = "53616c7465645f5f303f959de1a9733fa60cea3722f02a00fb8404cd9d9fa66c0d8e0e4d56c971bb67d0482d452070e01de490fc97d734af44571de48a07ae3d"
	adventOfCode = "https://adventofcode.com/2022/day/2/input"
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
