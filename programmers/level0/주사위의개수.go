package main

import "fmt"

func main() {
	fmt.Println(주사위의개수([]int{10, 8, 6}, 3))
}

func 주사위의개수(box []int, n int) int {

	w := box[0]
	l := box[1]
	h := box[2]

	var wCnt int
	var lCnt int
	var hCnt int

	for {
		if w-n < 0 {
			break
		}

		w -= n
		wCnt++
	}

	for {
		if l-n < 0 {
			break
		}

		l -= n
		lCnt++
	}

	for {
		if h-n < 0 {
			break
		}

		h -= n
		hCnt++
	}

	return wCnt * lCnt * hCnt
}
