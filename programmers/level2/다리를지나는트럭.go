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
	var passingTrucks []int
	var passedTrucks []int

	for {
		if len(passingTrucks) < bridgeLength && len(truckWeights) > 0 {
			//하나더 진행
			result++
			passingTrucks = append(passingTrucks, truckWeights[0])
			truckWeights = passTruck(truckWeights)

			if sumTruckWeight(passingTrucks) <= weight {
				continue
			} else {
				//되돌려놓기
				truckWeights = append([]int{passingTrucks[len(passingTrucks)-1]}, truckWeights...)
				passingTrucks = passingTrucks[:len(passingTrucks)-1]
				//result++

				//각각 한칸앞으로
				passedTrucks = append(passedTrucks, passingTrucks[0])
				passingTrucks = passTruck(passingTrucks)
				passingTrucks = append(passingTrucks, truckWeights[0])
				truckWeights = passTruck(truckWeights)
				result++
			}

		} else {
			//한칸앞으로
			passedTrucks = append(passedTrucks, passingTrucks[0])
			passingTrucks = passTruck(passingTrucks)
			result++
		}

		if len(passedTrucks) == truckCount {
			return result
		}
	}

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
