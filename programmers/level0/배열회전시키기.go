package main

import "fmt"

type Rotation struct {
	items []int
}

func (r *Rotation) moveLeft() {
	top := r.items[0]
	r.items = r.items[1:]
	r.items = append(r.items, top)
}

func (r *Rotation) moveRight() {
	bottom := r.items[len(r.items)-1]
	r.items = append([]int{bottom}, r.items[0:len(r.items)-1]...)
}

func (r *Rotation) execute(direction string) {
	switch direction {
	case "left":
		r.moveLeft()
	case "right":
		r.moveRight()
	}
}

func 배열회전시키기(numbers []int, direction string) []int {

	var r = Rotation{numbers}
	r.execute(direction)

	return r.items
}

func main() {
	fmt.Println(배열회전시키기([]int{1, 2, 3}, "right"))
}
