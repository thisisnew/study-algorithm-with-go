package main

import (
	"errors"
	"fmt"
	"sort"
)

type Jobs struct {
	jobs [][]int
}

func (j *Jobs) push(job []int) {
	j.jobs = append(j.jobs, job)
}

func (j *Jobs) pop() ([]int, error) {
	if len(j.jobs) == 0 {
		return []int{}, errors.New("empty jobs")
	}

	job := j.jobs[0]
	j.jobs = j.jobs[1:]

	return job, nil
}

func main() {
	fmt.Println(디스크컨트롤러([][]int{{0, 3}, {1, 9}, {2, 6}}))
}

func 디스크컨트롤러(jobs [][]int) int {

	jobTimes := getJobTimesAsc(jobs)
	var jobsAsc [][]int

	for _, jobTime := range jobTimes {
		jobsAsc = append(jobsAsc, getJobsAsc(jobTime, jobs)...)
	}

	var progressJobs = Jobs{jobsAsc}
	var result int
	var time int

	for {
		job, err := progressJobs.pop()

		if err != nil {
			return result / 3
		}

		if job[0] <= time {
			result = result + (time - job[0]) + (job[1] - job[0])
		}

		time = job[1]
	}
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
