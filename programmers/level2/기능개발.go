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

	progresses := []int{93, 30, 55}
	speeds := []int{1, 30, 5}

	fmt.Println(기능개발(progresses, speeds))
}

func 기능개발(progresses []int, speeds []int) []int {

	p := &ProgressQueue{}

	for i, pg := range progresses {
		remain := 100 - pg
		remainProgress := math.Ceil(float64(remain / speeds[i]))

		err := p.push(remainProgress)

		if err != nil {
			p.distribute(remainProgress)
		}
	}

	return p.result
}

func (p *ProgressQueue) push(remainProgress float64) error {

	if len(p.items) == 0 {
		p.items = append(p.items, int(remainProgress))
	}

	if int(remainProgress) > p.items[0] {
		return errors.New("must distribute")
	}

	p.items = append(p.items, int(remainProgress))

	return nil
}

func (p *ProgressQueue) distribute(remainProgress float64) {
	p.result = append(p.result, len(p.items))
	p.items = []int{int(remainProgress)}
}
