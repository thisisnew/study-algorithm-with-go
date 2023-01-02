package main

import (
	"fmt"
	"time"
)

const parkingTimeLayout = "15:04:05"

func main() {
	fmt.Println(ParkingBill("10:00", "13:21"))
}

func ParkingBill(E string, L string) int {
	var result = 2

	eTime, _ := time.Parse(parkingTimeLayout, E)
	lTime, _ := time.Parse(parkingTimeLayout, L)

	diff := lTime.Sub(eTime)

	difMin := diff.Minutes()

	fmt.Println(difMin)

	return result
}
