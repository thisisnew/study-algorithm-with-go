package main

import "fmt"

func main() {
	fmt.Println(주사위의개수([]int{10, 8, 6}, 3))
}

func 주사위의개수(box []int, n int) int {

	boxVolume := box[0] * box[1] * box[2]
	cube := n * n * n

	return boxVolume / cube
}
