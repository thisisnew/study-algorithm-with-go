package main

import (
	"errors"
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

		peek, err := p.peek()

		if err != nil {
			p.push(remainProgress)
			continue
		}

		if int(remainProgress) <= peek {
			p.push(remainProgress)
			continue
		}

		p.push(remainProgress)
		p.result = append(p.result, p.len())
	}

	if len(p.items) > 0 {
		p.result = append(p.result, len(p.items))
	}

	return p.result
}

func (p *ProgressQueue) peek() (int, error) {
	if len(p.items) == 0 {
		return 0, errors.New("empty slice")
	}

	return p.items[0], nil
}

func (p *ProgressQueue) push(remainProgress float64) {
	p.items = append(p.items, int(remainProgress))
}

func (p *ProgressQueue) pop() (int, error) {
	if len(p.items) == 0 {
		return 0, errors.New("empty slice")
	}

	return p.items[0], nil
}

func (p *ProgressQueue) len() int {
	return len(p.items)
}
