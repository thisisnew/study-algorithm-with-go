package main

import (
	"fmt"
	"time"
)

const parkingTimeLayout = "15:04"

func main() {
	fmt.Println(ParkingBill("10:00", "13:21"))
}

func ParkingBill(E string, L string) int {
	var result = 2

	eTime, _ := time.Parse(parkingTimeLayout, E)
	lTime, _ := time.Parse(parkingTimeLayout, L)
	difMin := lTime.Sub(eTime).Minutes()

	var isFirstTime = false
	var total int

	for difMin > 0 {
		if !isFirstTime {
			result += 3
			difMin -= 60
			isFirstTime = true
			continue
		}

		total++
		difMin -= 60
	}

	result += total * 4

	return result
}
