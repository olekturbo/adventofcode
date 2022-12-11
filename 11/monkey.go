package main

import (
	"adventofcode"
	"fmt"
	"math"
	"sort"
	"strings"
)

const (
	day = "11"

	rounds = 10000

	startingItems = "Starting items: "
	operation     = "Operation: new = old "
	test          = "Test: divisible by "
	passTest      = "If true: throw to monkey "
	failTest      = "If false: throw to monkey "
)

type monkey struct {
	items     []int
	operation string
	test      int
	passTest  int
	failTest  int
}

func main() {

	var monkeys []monkey
	nww := 1

	parts := strings.Split(string(adventofcode.Request(day)), "\n\n")
	for _, part := range parts {
		m := monkey{}
		for _, line := range strings.Split(part, "\n") {
			if len(line) != 0 {
				line = strings.TrimSpace(line)
				switch {
				case strings.HasPrefix(line, startingItems):
					line = strings.ReplaceAll(line, startingItems, "")
					items := adventofcode.MustAtoiArr(strings.Split(line, ", "))
					m.items = items
				case strings.HasPrefix(line, operation):
					line = strings.ReplaceAll(line, operation, "")
					m.operation = line
				case strings.HasPrefix(line, test):
					line = strings.ReplaceAll(line, test, "")
					m.test = adventofcode.MustAtoi(line)
				case strings.HasPrefix(line, passTest):
					line = strings.ReplaceAll(line, passTest, "")
					m.passTest = adventofcode.MustAtoi(line)
				case strings.HasPrefix(line, failTest):
					line = strings.ReplaceAll(line, failTest, "")
					m.failTest = adventofcode.MustAtoi(line)
				}
			}
		}
		monkeys = append(monkeys, m)
		nww *= m.test
	}

	inspections := make([]int, len(monkeys))

	for round := 0; round < rounds; round++ {
		for i, m := range monkeys {
			op := strings.Split(m.operation, " ")
			for range m.items {
				inspections[i]++

				var opValue int
				var value int

				if op[1] == "old" {
					opValue = m.items[0]
				} else {
					opValue = adventofcode.MustAtoi(op[1])
				}

				if op[0] == "+" {
					value = m.items[0] + opValue
				} else {
					value = m.items[0] * opValue
				}

				if rounds == 20 {
					value = int(math.Floor(float64(value / 3)))
				} else {
					value %= nww
				}

				testPassed := value%m.test == 0

				if testPassed {
					monkeys[m.passTest].items = append(monkeys[m.passTest].items, value)
				} else {
					monkeys[m.failTest].items = append(monkeys[m.failTest].items, value)
				}

				monkeys[i].items = append(monkeys[i].items[:0], monkeys[i].items[1:]...)
			}
		}
	}

	fmt.Println(monkeys)
	sort.Ints(inspections)
	fmt.Println(inspections[len(inspections)-1] * inspections[len(inspections)-2])
	println(nww)
}
