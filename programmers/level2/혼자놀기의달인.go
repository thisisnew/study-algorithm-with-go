package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(혼자놀기의달인([]int{8, 6, 3, 7, 2, 5, 1, 4}))
}

func 혼자놀기의달인(cards []int) int {

	var results []int

	for {

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

		results = append(results, calculateHighestBoxPoint(boxPoints))

		if len(results) == len(cards) {
			return getHighestBoxPoint(results)
		}

		c := cards[0]
		cards = cards[1:]
		cards = append(cards, c)

	}
}

func getHighestBoxPoint(results []int) int {

	sort.Slice(results, func(i, j int) bool {
		return results[i] > results[j]
	})

	return results[0]
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
