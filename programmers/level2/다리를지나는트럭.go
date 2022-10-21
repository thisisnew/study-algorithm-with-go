package main

import "fmt"

func main() {

	bridgeLength := 2
	weight := 10
	truckWeights := []int{7, 4, 5, 6}

	fmt.Println(다리를지나는트럭(bridgeLength, weight, truckWeights))

}

func 다리를지나는트럭(bridgeLength int, weight int, truckWeights []int) int {

	var result int
	var cTrucks []int

	result++

	totalWeight := getTotalTrucksWeight(cTrucks)

	if totalWeight < weight {

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
