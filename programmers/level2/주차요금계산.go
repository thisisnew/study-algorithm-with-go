package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

func main() {
	fees := []int{
		180,
		5000,
		10,
		600,
	}

	result := []string{
		"05:34 5961 IN",
		"06:00 0000 IN",
		"06:34 0000 OUT",
		"07:59 5961 OUT",
		"07:59 0148 IN",
		"18:59 0000 IN",
		"19:09 0148 OUT",
		"22:59 5961 IN",
		"23:00 5961 OUT",
	}

	fmt.Println(주차요금계산(fees, result))
}

func 주차요금계산(fees []int, records []string) []float64 {

	vehicles := calculateTimesByVehicles(records)

	return calculateFeesByVehicles(fees, vehicles)

}

func calculateTimesByVehicles(records []string) map[string]float64 {

	var temp = map[string][]string{}
	var result = map[string]float64{}

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

func getDurationTwoTimes(after, before string) float64 {

	layout := "15:04"
	a, _ := time.Parse(layout, after)
	b, _ := time.Parse(layout, before)

	return a.Sub(b).Minutes()
}

func calculateFeesByVehicles(fees []int, vehicles map[string]float64) []float64 {

	baseMinute := float64(fees[0])
	baseFee := float64(fees[1])
	unitMinute := float64(fees[2])
	unitFee := float64(fees[3])

	var result []float64

	for _, vh := range vehicles {

		if vh <= baseMinute {
			result = append(result, float64(baseFee))

			continue
		}

		fee := math.Ceil((vh - baseMinute) / unitMinute)

		sum := baseFee + fee*unitFee
		result = append(result, sum)
	}

	return result
}
