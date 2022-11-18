package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(최빈값구하기([]int{1, 2, 3, 3, 3, 4}))
}

func 최빈값구하기(array []int) int {

	var cntMap = map[int]int{}

	for _, num := range array {
		cntMap[num]++
	}

	var cnt []int

	for _, v := range cntMap {
		cnt = append(cnt, v)
	}

	sort.Slice(cnt, func(i, j int) bool {
		return cnt[i] > cnt[j]
	})

	var isKeyFind = false
	var key = 0
	for k, v := range cntMap {

		if v == cnt[0] {
			if isKeyFind {
				return -1
			}

			isKeyFind = true
			key = k
		}

	}

	if isKeyFind {
		return key
	}

	return -1
}
