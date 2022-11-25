package main

import (
	"fmt"
	"sort"
)

type HallOfFame struct {
	Items []int
}

func (h *HallOfFame) len() int {
	return len(h.Items)
}

func (h *HallOfFame) push(x int) {
	h.Items = append(h.Items, x)
}

func (h *HallOfFame) pop() int {
	if h.len() == 0 {
		return -1
	}

	p := h.Items[0]
	h.Items = h.Items[1:]
	return p
}

func (h *HallOfFame) top() int {
	if h.len() == 0 {
		return -1
	}

	return h.Items[0]
}

func (h *HallOfFame) sort() {
	if h.len() == 0 {
		return
	}

	sort.Slice(h.Items, func(i, j int) bool {
		return h.Items[i] < h.Items[j]
	})
}

type Minimums struct {
	Items   []int
	Minimum int
}

func (m *Minimums) push(x int) {
	m.Items = append(m.Items, x)
}

func (m *Minimums) min(x int) int {
	if x < m.Minimum {
		m.Minimum = x
	}

	return m.Minimum
}

func main() {
	fmt.Println(명예의전당(4, []int{0, 300, 40, 300, 20, 70, 150, 50, 500, 1000}))
}

func 명예의전당(k int, score []int) []int {

	var hall = HallOfFame{}
	var mins = Minimums{Minimum: score[0]}

	for _, s := range score {

		if hall.len() < k {
			mins.push(mins.min(s))
			hall.push(s)
			continue
		}

		hall.push(s)
		hall.sort()
		hall.pop()
		mins.Minimum = hall.top()
		mins.push(mins.Minimum)
	}

	return mins.Items
}
