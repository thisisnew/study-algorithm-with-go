package main

import "fmt"

func main() {

	bridgeLength := 2                 // 다리에 올라갈 수 있는 트럭 수
	weight := 10                      // 다리가 견딜 수 있는 무게
	truckWeights := []int{7, 4, 5, 6} //트럭 별 무게

	fmt.Println(다리를지나는트럭(bridgeLength, weight, truckWeights))

}

func 다리를지나는트럭(bridgeLength int, weight int, truckWeights []int) int {

	var result int
	var truckCount = len(truckWeights)
	var passingTrucks = make([]int, bridgeLength)
	var passedTrucks []int

	for {

		if sumTruckWeight(passingTrucks) == 0 {
			passingTrucks = climbTheBridge(passingTrucks, truckWeights[0])
			truckWeights = passTruck(truckWeights)
			continue
		}

		//한칸진행
		passingTrucks, passedTrucks = crossingTheBridge(passingTrucks, passedTrucks, bridgeLength)

		if passingTrucks[0] == 0 && len(truckWeights) > 0 {
			if sumTruckWeight(passingTrucks)+truckWeights[0] <= weight {
				passingTrucks = climbTheBridge(passedTrucks, truckWeights[0])
				truckWeights = passTruck(truckWeights)
			}
		}

		result++

		if len(passedTrucks) == truckCount {
			return result
		}
	}

}

func climbTheBridge(passingTrucks []int, truck int) []int {

	var idx = 0

	for i, p := range passingTrucks {
		if p == 0 {
			idx = i
			break
		}
	}

	passingTrucks[idx] = truck

	return passingTrucks
}

func crossingTheBridge(passingTrucks, passedTrucks []int, bridgeLength int) ([]int, []int) {

	var result = make([]int, bridgeLength)

	for i, p := range passingTrucks {

		//끝에 도달하면 넘어가기
		if i == len(passingTrucks)-1 {
			passedTrucks = append(passedTrucks, p)
			passingTrucks[i] = 0
			continue
		}

		result[i+1] = p
	}

	return result, passedTrucks

}

func passTruck(truck []int) []int {

	if len(truck) == 0 {
		return []int{}
	}

	return truck[1:]
}

func sumTruckWeight(passingTrucks []int) int {

	var result int

	for _, p := range passingTrucks {
		result += p
	}

	return result

}
