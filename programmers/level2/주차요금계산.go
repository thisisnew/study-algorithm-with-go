package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	fees := []int{180, 5000, 10, 600}
	result := []string{"05:34 5961 IN", "06:00 0000 IN", "06:34 0000 OUT", "07:59 5961 OUT", "07:59 0148 IN", "18:59 0000 IN", "19:09 0148 OUT", "22:59 5961 IN", "23:00 5961 OUT"}

	fmt.Println(주차요금계산(fees, result))
}

func 주차요금계산(fees []int, records []string) []int {

	//baseMinute := fees[0]
	//baseFee := fees[1]
	//unitMinute := fees[2]
	//unitFee := fees[3]

	vehicles := calculateTimesByVehicles(records)
	fmt.Println(vehicles)
	return []int{}
}

func calculateTimesByVehicles(records []string) map[string]int {

	var temp = map[string][]string{}
	var result = map[string]int{}

	for _, record := range records {

		recordSlice := strings.Split(record, " ")

		tm := recordSlice[0]
		num := recordSlice[1]

		temp[num] = append(temp[num], tm)

	}

	for k, v := range temp {

		if len(v)%2 == 0 {
			continue
		}

		temp[k] = append(temp[k], "23:59")
	}

	for k, v := range temp {

		for i := len(v) - 1; i >= 0; i-- {

			if i%2 == 0 {
				continue
			}

			result[k] += getDurationTwoTimes(v[i], v[i-1])

		}
	}

	return result
}

func getDurationTwoTimes(after, before string) int {

	layout := "15:04"
	a, _ := time.Parse(layout, after)
	b, _ := time.Parse(layout, before)

	return int(a.Sub(b).Minutes())
}
