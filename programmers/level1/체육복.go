package main

import (
	"fmt"
	"sort"
)

func main() {

	n := 3
	lost := []int{1, 2}
	reserve := []int{2, 3}

	fmt.Println(체육복(n, lost, reserve))
}

func 체육복(n int, lost []int, reserve []int) int {

	sort.Slice(lost, func(i, j int) bool {
		return lost[i] < lost[j]
	})

	sort.Slice(reserve, func(i, j int) bool {
		return reserve[i] < reserve[j]
	})

	var count int

	for i := 0; i < len(lost); i++ {
		for j := 0; j < len(reserve); j++ {
			if lost[i] == reserve[j] {
				count++
				lost[i] = -1
				reserve[j] = -1
				break
			}
		}
	}

	for i := 0; i < len(lost); i++ {
		for j := 0; j < len(reserve); j++ {
			if lost[i] == reserve[j]+1 || lost[i] == reserve[j]-1 {
				count++
				reserve[j] = -1
				break
			}
		}
	}

	return n - len(lost) + count

}
