package main

import "fmt"

func main() {
	fmt.Println(배달(5, [][]int{{1, 2, 1}, {2, 3, 3}, {5, 2, 2}, {1, 4, 2}, {5, 3, 1}, {5, 4, 2}}, 3))
}

var (
	isVisitedDelivery []bool
)

func 배달(N int, road [][]int, k int) int {
	answer := 0

	// [실행] 버튼을 누르면 출력 값을 볼 수 있습니다.
	fmt.Printf("Hello Go")

	isVisitedDelivery = make([]bool, N)

	for i := 0; i < N; i++ {
		dfsDelivery(i, road)
	}

	return answer
}

func dfsDelivery(v int, road [][]int) {
	isVisitedDelivery[v] = true

	for i := 0; i < len(road[v]); i++ {
		if road[v][i] == 1 && !isVisitedDelivery[i] {
			dfsDelivery(i, road)
		}
	}
}
