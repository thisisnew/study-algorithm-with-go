package main

import "fmt"

func main() {
	fmt.Println(조이스틱("JEROEN"))
}

func 조이스틱(name string) int {
	var result int
	var idx int

	for _, rn := range name {

		upCnt, upIdx := controlUp(idx, rn-65)
		downCnt, downIdx := controlDown(idx, rn-65)

		if upCnt > downCnt {
			result += downCnt
			idx = downIdx
		} else {
			result += upCnt
			idx = upIdx
		}

	}

	return result
}

func controlUp(idx int, dist int32) (int, int) {

	var result int

	for {
		if int32(idx) == dist {
			break
		}

		result++
		idx++

		if idx > 26 {
			idx = 0
		}
	}

	return result, idx
}

func controlDown(idx int, dist int32) (int, int) {
	var result int

	for {
		if int32(idx) == dist {
			break
		}

		result++
		idx--

		if idx < 0 {
			idx = 26
		}
	}

	return result, idx
}

func cursorLeft() {

}

func cursorRight() {

}
