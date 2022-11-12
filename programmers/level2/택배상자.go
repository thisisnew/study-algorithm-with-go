package main

import "fmt"

func main() {
	fmt.Println(택배상자([]int{5, 4, 3, 2, 1}))
}

type Packs struct {
	items []int
}

func (p *Packs) pop() int {

	if len(p.items) == 0 {
		return -1
	}

	pack := p.items[len(p.items)-1]

	p.items = p.items[0 : len(p.items)-1]

	return pack
}

func (p *Packs) top() int {

	if len(p.items) == 0 {
		return -1
	}

	return p.items[len(p.items)-1]
}

func (p Packs) empty() bool {
	return len(p.items) == 0
}

func (p *Packs) push(x int) {
	p.items = append(p.items, x)
}

func (p Packs) contains(x int) bool {

	for _, pack := range p.items {
		if x == pack {
			return true
		}
	}

	return false
}

func 택배상자(order []int) int {

	var result int
	packsInContainer := getPacksInContainer(order)
	subContainer := Packs{}

out:
	for {
		if len(order) == 0 {
			return result
		}

		o := order[0]

		if packsInContainer.contains(o) {

			for {
				p := packsInContainer.pop()

				if p == o {
					result++
					order = order[1:]
					continue out
				}

				subContainer.push(p)
			}
		}

		if order[0] == subContainer.top() {
			result++
			order = order[1:]
			subContainer.pop()
			continue out
		}

		return result
	}
}

func getPacksInContainer(order []int) Packs {

	var packs = make([]int, len(order))

	for i := 0; i < len(packs); i++ {
		packs[i] = len(packs) - i
	}

	return Packs{packs}
}
