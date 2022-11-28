package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(디스크컨트롤러([][]int{{0, 3}, {1, 9}, {2, 6}}))
}

func 디스크컨트롤러(jobs [][]int) int {

	jobTimes := getJobTimesAsc(jobs)
	var jobsAsc [][]int

	for _, jobTime := range jobTimes {
		jobsAsc = append(jobsAsc, getJobsAsc(jobTime, jobs)...)
	}

	return 0
}

func getJobTimesAsc(jobs [][]int) []int {
	var result = make([]int, len(jobs))

	for i, job := range jobs {
		result[i] = job[1] - job[0]
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})

	return result
}

func getJobsAsc(jobTime int, jobs [][]int) [][]int {

	var result [][]int

	for _, job := range jobs {
		if job[1]-job[0] != jobTime {
			continue
		}

		result = append(result, job)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})

	return result
}
