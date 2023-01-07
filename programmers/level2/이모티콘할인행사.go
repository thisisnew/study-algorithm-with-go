package main

import "fmt"

func main() {
	fmt.Println([][]int{{40, 10000}, {25, 10000}}, []int{7000, 9000})
}

func 이모티콘할인행사(users [][]int, emoticons []int) []int {

	var saleRatio = []int{40, 30, 20, 10}

	for _, user := range users {

		acceptableRange := user[0]
		limitPrice := user[1]

		for _, sr := range saleRatio {

		}

	}

	return []int{}
}
