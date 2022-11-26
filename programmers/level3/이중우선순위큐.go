package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(이중우선순위큐([]string{"I -45", "I 653", "D 1", "I -642", "I 45", "I 97", "D 1", "D -1", "I 333"}))
}

type Operation struct {
	Items    []int
	IsSorted bool
}

func (o *Operation) push(x int) {
	o.Items = append(o.Items, x)
}

func (o *Operation) Max() int {
	if len(o.Items) == 0 {
		return 0
	}

	if !o.IsSorted {
		o.sort()
	}

	return o.Items[len(o.Items)-1]
}

func (o *Operation) Min() int {
	if len(o.Items) == 0 {
		return 0
	}

	if !o.IsSorted {
		o.sort()
	}

	return o.Items[0]
}

func (o *Operation) removeMax() error {
	if len(o.Items) == 0 {
		return errors.New("empty items")
	}

	if !o.IsSorted {
		o.sort()
	}

	o.Items = o.Items[:len(o.Items)-1]

	return nil
}

func (o *Operation) removeMin() error {
	if len(o.Items) == 0 {
		return errors.New("empty items")
	}

	if !o.IsSorted {
		o.sort()
	}

	o.Items = o.Items[1:]

	return nil
}

func (o *Operation) sort() error {
	if len(o.Items) == 0 {
		return errors.New("empty items")
	}

	sort.Slice(o.Items, func(i, j int) bool {
		return o.Items[i] < o.Items[j]
	})

	o.IsSorted = true

	return nil
}

func 이중우선순위큐(operations []string) []int {

	var result = Operation{}

	for _, operation := range operations {
		commands := strings.Fields(operation)

		switch commands[0] {
		case "I":
			v, _ := strconv.Atoi(commands[1])
			result.push(v)
			result.IsSorted = false
		default:
			if commands[1] == "1" {
				result.removeMax()
			} else {
				result.removeMin()
			}
		}
	}

	return []int{result.Max(), result.Min()}
}
