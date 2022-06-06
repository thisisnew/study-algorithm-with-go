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

		if !isValidYellowWidth(yellow, yWidth) {
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

		break
	}

	return []int{bWidth, bHeight}

}

func isValidYellowWidth(yellow int, yWidth int) bool {
	if yellow%yWidth != 0 {
		return false
	}

	return true
}
