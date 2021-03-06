package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"time"
)

const (
	Layout   = "15:04"
	LastTime = "23:59"
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

func 주차요금계산(fees []int, records []string) []int {

	useTimes, vehicles := calculateTimesByVehicles(records)

	return calculateFeesByVehicles(fees, useTimes, vehicles)
}

func calculateTimesByVehicles(records []string) (map[string]float64, []string) {

	var temp = map[string][]string{}
	var result = map[string]float64{}
	var tempVehicles = map[string][]string{}
	var vehicles []string

	for _, record := range records {

		recordSlice := strings.Split(record, " ")

		tm := recordSlice[0]
		num := recordSlice[1]

		temp[num] = append(temp[num], tm)

		if !isDuplicate(vehicles, num) {
			vehicles = append(vehicles, num)
		}
	}

	sort.Strings(vehicles)

	for _, vh := range vehicles {
		tempVehicles[vh] = temp[vh]
	}

	for k, vh := range tempVehicles {

		if len(vh)%2 == 0 {
			continue
		}

		vh = append(vh, LastTime)
		tempVehicles[k] = vh
	}

	for _, vh := range vehicles {

		v := tempVehicles[vh]

		for i := len(v) - 1; i >= 0; i-- {

			if i%2 == 0 {
				continue
			}

			result[vh] += getDurationTwoTimes(v[i], v[i-1])
		}
	}

	return result, vehicles
}

func isDuplicate(vehicles []string, vehicle string) bool {

	for _, vh := range vehicles {
		if vh == vehicle {
			return true
		}
	}

	return false
}

func getDurationTwoTimes(after, before string) float64 {

	a, _ := time.Parse(Layout, after)
	b, _ := time.Parse(Layout, before)

	return a.Sub(b).Minutes()
}

func calculateFeesByVehicles(fees []int, useTimes map[string]float64, vehicles []string) []int {

	baseMinute := float64(fees[0])
	baseFee := float64(fees[1])
	unitMinute := float64(fees[2])
	unitFee := float64(fees[3])

	var result []int

	for _, vh := range vehicles {

		useTime := useTimes[vh]

		if useTime <= baseMinute {
			result = append(result, int(baseFee))

			continue
		}

		sum := baseFee + math.Ceil((useTime-baseMinute)/unitMinute)*unitFee
		result = append(result, int(sum))
	}

	return result
}
