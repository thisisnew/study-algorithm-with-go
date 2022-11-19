package main

import (
	"fmt"
	"sort"
)

type Printer struct {
	items []int
}

func (p *Printer) pop() int {
	var result = p.items[0]

	p.items = p.items[1:]

	return result
}

func (p *Printer) top() int {
	var result = p.items[0]
	return result
}

func (p *Printer) push(x int) {
	p.items = append(p.items, x)
}

func (p *Printer) copy(x []int) {
	var cp = make([]int, len(x))
	copy(cp, x)
	p.items = cp
}

func (p *Printer) reverseSort() {
	sort.Slice(p.items, func(i, j int) bool {
		return p.items[i] > p.items[j]
	})
}

func main() {
	priorities := []int{2, 1, 3, 2}
	location := 2
	fmt.Println(print(priorities, location))
}

func print(priorities []int, location int) int {

	var count int
	var result int
	var pr = Printer{priorities}
	var cp = Printer{}
	cp.copy(pr.items)
	cp.reverseSort()

	for {
		p := pr.pop()

		if p == cp.items[0] {
			count++
			cp.pop()

			if location == 0 {
				result = count
				break
			}

			location--
			continue
		}

		pr.push(p)

		if location == 0 {
			location = len(pr.items) - 1
		} else {
			location--
		}
	}

	return result
}
