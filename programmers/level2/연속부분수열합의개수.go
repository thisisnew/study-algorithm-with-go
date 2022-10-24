package main

import "fmt"

var dupElementsMap = map[int]bool{}

func main() {
	fmt.Println(연속부분수열합의개수([]int{7, 9, 1, 1, 4}))
}

func 연속부분수열합의개수(elements []int) int {

	var result = 0
	var dif = 1

	for {
		result += len(countElementsSumByDif(elements, dif))

		if dif == len(elements) {
			break
		}

		dif++
	}

	return result
}

func countElementsSumByDif(elements []int, dif int) []int {
	var result []int

	for start := 0; start < len(elements); start++ {
		sum := sumElementsBetweenDifByStartIndex(elements, start, dif)

		if ok, _ := dupElementsMap[sum]; ok {
			continue
		}

		dupElementsMap[sum] = true
		result = append(result, sum)
	}

	return result
}

func sumElementsBetweenDifByStartIndex(elements []int, start, dif int) int {
	var result = 0
	var idx = start
	var cnt = 0

	for {
		result += elements[idx]
		idx++
		cnt++

		if cnt == dif {
			return result
		}

		if idx > len(elements)-1 {
			idx = 0
		}
	}
}
