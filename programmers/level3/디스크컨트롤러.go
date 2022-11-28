package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(디스크컨트롤러([][]int{{0, 3}, {1, 9}, {2, 6}}))
}

func 디스크컨트롤러(jobs [][]int) int {

	var jobTimes = make([]int, len(jobs))

	for i, job := range jobs {
		jobTimes[i] = job[1] - job[0]
	}

	sort.Slice(jobTimes, func(i, j int) bool {
		return jobTimes[i] < jobTimes[j]
	})

	return 0
}
