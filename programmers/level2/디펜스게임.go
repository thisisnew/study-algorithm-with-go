package main

import (
	"errors"
	"fmt"
)

type DefenseGame struct {
	games []int
}

func (d *DefenseGame) push(g int) {
	d.games = append(d.games, g)
}

func (d *DefenseGame) pop() (int, error) {

	if len(d.games) == 0 {
		return -1, errors.New("empty games")
	}

	g := d.games[len(d.games)-1]
	d.games = d.games[0 : len(d.games)-1]

	return g, nil
}

func (d *DefenseGame) empty() bool {
	return len(d.games) == 0
}

func main() {
	fmt.Println(디펜스게임(7, 3, []int{4, 2, 4, 5, 3, 3, 1}))
}

func 디펜스게임(n int, k int, enemy []int) int {

	if k == len(enemy) {
		return k
	}

	var defense = DefenseGame{}

	for _, e := range enemy {
		defense.push(e)
	}

	var result int

	return result
}
