package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(혼자놀기의달인([]int{2, 3, 1}))
}

func 혼자놀기의달인(cards []int) int {

	var boxPoints []int

	for i := 0; i < len(cards); i++ {
		var openedBoxMap = map[int]bool{}
		isOpened, _ := openedBoxMap[i]

		if isOpened {
			continue
		}

		boxNum := cards[i]
		openedBoxMap[boxNum] = true

		for {
			boxNum = cards[boxNum-1]
			isOpened, _ = openedBoxMap[boxNum]

			if isOpened {
				boxPoints = append(boxPoints, len(openedBoxMap))
				break
			}

			openedBoxMap[boxNum] = true
		}

	}

	return getHighestBoxPoint(boxPoints)
}

func getHighestBoxPoint(boxPoints []int) int {

	if len(boxPoints) <= 1 {
		return 0
	}

	sort.Slice(boxPoints, func(i, j int) bool {
		return boxPoints[i] > boxPoints[j]
	})

	return boxPoints[0] * boxPoints[1]
}
