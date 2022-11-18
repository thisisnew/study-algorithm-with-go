package main

import (
	"fmt"
	"sort"
)

//[[1, 2], [1, 1], [1, 1]] -> [1, 2, 2]

func main() {
	fmt.Println(등수매기기([][]int{{1, 2}, {1, 1}, {1, 1}}))
}

func 등수매기기(score [][]int) []int {

	var avgs = make([]float64, len(score))
	var cp = make([]float64, len(score))

	for i, sc := range score {
		avs := getAverageScore(sc)
		avgs[i] = avs
	}

	copy(cp, avgs)

	sort.Slice(cp, func(i, j int) bool {
		return cp[i] > cp[j]
	})

	return rankAverageScores(avgs, cp)
}

func getAverageScore(sc []int) float64 {

	ln := len(sc)
	var result int

	for _, s := range sc {
		result += s
	}

	return float64(result) / float64(ln)
}

func rankAverageScores(avgs, cp []float64) []int {

	var result []int
	var prevScore float64
out:
	for _, avg := range avgs {
		for j, sc := range cp {
			if avg == sc {
				if len(result) > 0 && prevScore == sc {
					result = append(result, result[len(result)-1])
				} else {
					result = append(result, j+1)
				}

				prevScore = avg
				continue out
			}
		}

	}

	return result
}
