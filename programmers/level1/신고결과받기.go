package main

import (
	"fmt"
	"strings"
)

func main() {

	idList := []string{"muzi", "frodo", "apeach", "neo"}
	report := []string{
		"muzi frodo", "apeach frodo", "frodo neo", "muzi neo", "apeach muzi",
	}

	fmt.Println(신고결과받기(idList, report, 2))
}

func 신고결과받기(id_list []string, report []string, k int) []int {

	reporters := make(map[string]int, len(report))
	userMap := generateUserMap(id_list)

	for _, rp := range report {
		reportedPerson := strings.Split(rp, " ")[1]
		userMap[reportedPerson]++
	}

	stopUserMap := generateStopUserMap(userMap, k)

	for _, rp := range report {

		rpDetail := strings.Split(rp, " ")
		reporter := rpDetail[0]
		reportedPerson := rpDetail[1]

		isReportedPersonStopped := isReportedPersonStopped(reportedPerson, stopUserMap)

		if isReportedPersonStopped {
			reporters[reporter]++
		}
	}

	var result []int

	for _, id := range id_list {
		result = append(result, reporters[id])
	}

	return result
}

func generateUserMap(id_list []string) map[string]int {

	result := make(map[string]int, len(id_list))

	for _, id := range id_list {
		result[id] = 0
	}

	return result
}

func generateStopUserMap(userMap map[string]int, k int) map[string]bool {

	result := make(map[string]bool, len(userMap))

	for key, val := range userMap {

		if val == k {
			result[key] = true
			continue
		}

		result[key] = false
	}

	return result
}

func isReportedPersonStopped(reportedPerson string, stopUserMap map[string]bool) bool {

	for k, v := range stopUserMap {

		if k != reportedPerson {
			continue
		}

		return v
	}

	return false
}
