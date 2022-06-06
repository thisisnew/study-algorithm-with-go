package main

import "fmt"

func main() {
	brown := 24
	yellow := 24

	fmt.Println(카펫(brown, yellow))
}

func 카펫(brown int, yellow int) []int {

	var yWidth = 1
	var yHeight int
	var bWidth int
	var bHeight int

	for {

		if !isValidYellowWidth(yellow, yWidth) {
			yWidth++
			continue
		}

		yHeight = calculateYellowHeight(yellow, yWidth)

		bWidth, bHeight = calculateBrownWidthAndHeight(yWidth, yHeight)

		if countBrown(calculateArea(bWidth, bHeight), calculateArea(yWidth, yHeight)) != brown {
			yWidth++
			continue
		}

		if isBrownHeightLongerThenBrownWidth(bHeight, bWidth) {
			yWidth++
			continue
		}

		return []int{bWidth, bHeight}
	}

}

func isValidYellowWidth(yellow int, yWidth int) bool {
	if yellow%yWidth != 0 {
		return false
	}

	return true
}

func calculateYellowHeight(yellow int, yWidth int) int {
	return yellow / yWidth
}

func calculateBrownWidthAndHeight(yWidth, yHeight int) (int, int) {
	bWidth := yWidth + 2
	bHeight := yHeight + 2

	return bWidth, bHeight
}

func calculateArea(width, height int) int {
	return width * height
}

func countBrown(bArea int, yArea int) int {
	return bArea - yArea
}

func isBrownHeightLongerThenBrownWidth(bHeight, bWidth int) bool {
	return bHeight > bWidth
}
