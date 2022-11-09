package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(혼자놀기의달인([]int{2, 3, 1}))
}

func 혼자놀기의달인(cards []int) int {

	var openedBoxMap = map[int]bool{}
	var boxPoints []int

	for i := 0; i < len(cards); i++ {
		var boxPoint = 0

		isOpened, ok := openedBoxMap[i]

		if ok && isOpened {
			continue
		}

		boxNum := cards[i]
		openedBoxMap[boxNum] = true
		boxPoint++

		for {
			boxNum = cards[boxNum-1]
			isOpened, ok = openedBoxMap[boxNum]

			if ok && isOpened {
				boxPoints = append(boxPoints, boxPoint)
				break
			}

			openedBoxMap[boxNum] = true
			boxPoint++
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
