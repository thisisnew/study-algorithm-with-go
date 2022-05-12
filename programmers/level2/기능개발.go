package main

import (
	"fmt"
	"math"
)

type ProgressQueue struct {
	items  []int
	result []int
}

func main() {

	progresses := []int{95, 90, 99, 99, 80, 99}
	speeds := []int{1, 1, 1, 1, 1, 1}

	fmt.Println(기능개발(progresses, speeds))
}

func 기능개발(progresses []int, speeds []int) []int {

	p := &ProgressQueue{}

	for i, pg := range progresses {
		remain := 100 - pg
		remainProgress := math.Ceil(float64(remain) / float64(speeds[i]))

		peek := p.peek()

		if peek == nil || int(remainProgress) <= *peek {
			p.push(remainProgress)
			continue
		}

		p.result = append(p.result, len(p.items))
		p.items = []int{int(remainProgress)}
	}

	if len(p.items) > 0 {
		p.result = append(p.result, len(p.items))
	}

	return p.result
}

func (p *ProgressQueue) peek() *int {
	if len(p.items) == 0 {
		return nil
	}

	return &p.items[0]
}

func (p *ProgressQueue) push(remainProgress float64) {
	p.items = append(p.items, int(remainProgress))
}
