package main

import (
	"fmt"
	"strings"
)

func main() {

	idList := []string{"con", "ryan"}
	report := []string{
		"ryan con", "ryan con", "ryan con", "ryan con",
	}

	fmt.Println(신고결과받기(idList, report, 3))
}

func 신고결과받기(id_list []string, report []string, k int) []int {

	reporters := generateReporterList(report)
	reportedPeopleCount := generateReportedPersonCount(reporters)

	var result []int
	for _, id := range id_list {
		var cnt int
		reportedPeople := reporters[id]

		for _, reportedPerson := range reportedPeople {
			if reportedPeopleCount[reportedPerson] >= k {
				cnt++
			}
		}

		result = append(result, cnt)
	}

	return result
}

func generateReporterList(report []string) map[string][]string {

	var reporters = make(map[string][]string, len(report))

	for _, rp := range report {
		rpSlice := strings.Split(rp, " ")

		reporters[rpSlice[0]] = addReportedPerson(reporters[rpSlice[0]], rpSlice[1])
	}

	return reporters
}

func addReportedPerson(reportedPeople []string, reportedPerson string) []string {

	if reportedPeople == nil {
		return []string{reportedPerson}
	}

	for _, rp := range reportedPeople {
		if rp == reportedPerson {
			return reportedPeople
		}
	}

	reportedPeople = append(reportedPeople, reportedPerson)
	return reportedPeople
}

func generateReportedPersonCount(reporters map[string][]string) map[string]int {
	var result = make(map[string]int, len(reporters))

	for _, reportedPeople := range reporters {
		for _, reportedPerson := range reportedPeople {
			result[reportedPerson]++
		}
	}

	return result
}
