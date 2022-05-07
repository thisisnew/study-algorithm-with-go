package main

import "fmt"

func main() {

	idList := []string{"muzi", "frodo", "apeach", "neo"}
	report := []string{
		"muzi frodo", "apeach frodo", "frodo neo", "muzi neo", "apeach muzi",
	}

	fmt.Println(신고결과받기(idList, report, 2))
}

func 신고결과받기(id_list []string, report []string, k int) []int {

	var result []int

	for _, id := range id_list {

		var cnt int

		for _, rp := range report {

		}

	}

	return result
}
