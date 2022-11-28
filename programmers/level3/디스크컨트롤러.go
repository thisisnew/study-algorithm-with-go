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
	fmt.Println(디스크컨트롤러([][]int{{0, 3}, {1, 9}, {2, 6}}))
}

func 디스크컨트롤러(jobs [][]int) int {

	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i][0] < jobs[j][0]
	})

	var waitingJobs = Jobs{jobs, false}
	var progressJobs = Jobs{}
	var result int
	var time int

	for {

		switch {
		case progressJobs.empty() && waitingJobs.empty():
			break
		case !progressJobs.empty():
			progressJob(&progressJobs, &time, &result)
		default:
			pushProgress(&waitingJobs, &progressJobs, &time, &result)
		}

	}

	return result / 3
}

func progressJob(progressJobs *Jobs, time *int, result *int) {
	progressJobs.sort()

	pj, _ := progressJobs.pop()

	if pj[0] > *time {
		*result += pj[0] - *time
		time = &(pj)[0]
	}

	*result += pj[1]
	*time = *time + pj[1]
}

func pushProgress(waitingJobs, progressJobs *Jobs, time, result *int) {
	for {
		t, err := waitingJobs.top()

		if err != nil {
			return
		}

		if t[0] > *time {
			if !progressJobs.empty() {
				return
			}

			*result += t[0] - *time
			time = &(t)[0]

		} else {

			wj, _ := waitingJobs.pop()
			progressJobs.push(wj)
		}
	}
}
