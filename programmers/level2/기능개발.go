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

		ln, err := p.len()

		if err == nil {
			p.result = append(p.result, ln)
		}

		p.push(remainProgress)
	}

	ln, err := p.len()

	if err == nil {
		p.result = append(p.result, ln)
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

	pop := p.items[0]

	p.items = p.items[1:]

	return pop, nil
}

func (p *ProgressQueue) len() (int, error) {

	if len(p.items) == 0 {
		return 0, errors.New("empty slice")
	}

	var ln int

	for {

		_, err := p.pop()

		if err != nil {
			break
		}

		ln++
	}

	return ln, nil
}
