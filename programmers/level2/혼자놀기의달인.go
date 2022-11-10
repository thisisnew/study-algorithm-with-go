package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(혼자놀기의달인([]int{2, 1}))
}

func 혼자놀기의달인(cards []int) int {

	var results []int
	var startIdx = 0
	var idx = 0

	for {

		var openedBoxMap = map[int]bool{}
		var boxPoints []int

		idx = startIdx
		var boxPoint = 0

		isOpened, ok := openedBoxMap[idx]

		if ok && isOpened {
			continue
		}

		boxNum := cards[idx]
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

		results = append(results, calculateHighestBoxPoint(boxPoints))

		if len(results) == len(cards) {
			return getHighestBoxPoint(results)
		}

		startIdx++
		if startIdx > len(cards)-1 {
			startIdx = 0
		}
	}
}

func getHighestBoxPoint(results []int) int {

	sort.Slice(results, func(i, j int) bool {
		return results[i] > results[j]
	})

	return results[0] * results[1]
}

func calculateHighestBoxPoint(boxPoints []int) int {

	if len(boxPoints) <= 1 {
		return 0
	}

	sort.Slice(boxPoints, func(i, j int) bool {
		return boxPoints[i] > boxPoints[j]
	})

	return boxPoints[0] * boxPoints[1]
}
