package main

import "fmt"

func main() {
	brown := 24
	yellow := 24

	fmt.Println(카펫(brown, yellow))
}

func 카펫(brown int, yellow int) []int {

	var yWidth = 1
	var bWidth int
	var bHeight int

	for {

		if yellow%yWidth != 0 {
			yWidth++
			continue
		}

		yHeight := yellow / yWidth

		bWidth = yWidth + 2
		bHeight = yHeight + 2

		bArea := bWidth * bHeight
		yArea := yWidth * yHeight

		if bArea-yArea != brown {
			yWidth++
			continue
		}

		if bHeight > bWidth {
			yWidth++
			continue
		}

		return []int{bWidth, bHeight}
	}
}
