package main

import (
	"errors"
	"fmt"
	"sort"
)

type Jobs struct {
	Jobs     [][]int
	IsSorted bool
}

func (j *Jobs) push(job []int) {
	j.Jobs = append(j.Jobs, job)
	j.IsSorted = false
}

func (j *Jobs) pop() ([]int, error) {
	if len(j.Jobs) == 0 {
		return []int{}, errors.New("empty jobs")
	}

	job := j.Jobs[0]
	j.Jobs = j.Jobs[1:]

	return job, nil
}

func (j *Jobs) top() ([]int, error) {
	if len(j.Jobs) == 0 {
		return []int{}, errors.New("empty jobs")
	}

	return j.Jobs[0], nil
}

func (j *Jobs) empty() bool {
	return len(j.Jobs) == 0
}

func (j *Jobs) sort() {
	if !j.IsSorted {
		sort.Slice(j.Jobs, func(a, b int) bool {
			return j.Jobs[a][1] < j.Jobs[b][1]
		})
	}

	j.IsSorted = true
}

func main() {
	fmt.Println(디스크컨트롤러([][]int{{7, 8}, {3, 5}, {9, 6}}))
	//fmt.Println(디스크컨트롤러([][]int{{0, 3}, {1, 9}, {2, 6}}))
}

func 디스크컨트롤러(jobs [][]int) int {

	var result int
	var time int
	var waitingJobs = Jobs{jobs, false}
	var progressJobs = Jobs{}

	for {
		if waitingJobs.empty() {
			processJobs(&progressJobs, &result, &time)
			break
		}

		t, _ := waitingJobs.top()

		if time >= t[0] {
			p, _ := waitingJobs.pop()
			progressJobs.push(p)
			continue
		}

		if !progressJobs.empty() {
			processJobs(&progressJobs, &result, &time)
			continue
		}

		time++
	}

	return result / 3
}

func processJobs(progressJobs *Jobs, result *int, time *int) {

	progressJobs.sort()

	for {
		if progressJobs.empty() {
			return
		}

		job, err := progressJobs.pop()

		if err != nil {
			return
		}

		//이 부분이 의심됨
		*result += (*time - job[0]) + job[1]

		//job[0] + job[1]

		if job[0]+job[1] > *time {
			*time = job[0] + job[1]
		}

		*time += job[1]
	}
}
