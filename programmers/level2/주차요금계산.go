package main

import "strings"

func main() {

}

func 주차요금계산(fees []int, records []string) []int {

	baseMinute := fees[0]
	baseFee := fees[1]
	unitMinute := fees[2]
	unitFee := fees[3]

	vehicles := calculateTimesByVehicles(records)

	return []int{}
}

func calculateTimesByVehicles(records []string) map[string]int {

	var result = map[string]int{}

	for _, record := range records {

		recordSlice := strings.Split(record, " ")

	}

	return result
}
