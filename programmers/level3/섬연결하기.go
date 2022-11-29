package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(섬연결하기(4, [][]int{{0, 1, 1}, {0, 2, 2}, {1, 2, 5}, {1, 3, 1}, {2, 3, 8}}))
}

func 섬연결하기(n int, costs [][]int) int {

	sort.Slice(costs, func(i, j int) bool {
		return costs[i][0] < costs[j][0]
	})

	var routeMap = make(map[int][]int)
	//var costMap = make(map[int]int)

	for _, cost := range costs {
		st := cost[0]
		ed := cost[1]
		//v := cost[2]

		routeMap[st] = append(routeMap[st], ed)

	}

	fmt.Println(routeMap)
	//fmt.Println(costMap)
	return 0
}
