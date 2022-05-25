package main

import "fmt"

func main() {

	bridge_length := 2
	weight := 10
	truck_weights := []int{7, 4, 5, 6}

	fmt.Println(다리를지나는트럭(bridge_length, weight, truck_weights))

}

func 다리를지나는트럭(bridge_length int, weight int, truck_weights []int) int {

	var result int
	var cTrucks []int

	result++

	totalWeight := getTotalTrucksWeight(cTrucks)

	if totalWeight > weight {

	}

	return result
}

func getTotalTrucksWeight(cTrucks []int) int {

	var result int

	for _, c := range cTrucks {
		result += c
	}

	return result

}

func crossingTheBridge() {

}

func waitingTrucks() {

}
