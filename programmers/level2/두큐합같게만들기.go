package main

import "fmt"

func main() {
	fmt.Println(두큐합같게만들기([]int{3, 2, 7, 2}, []int{4, 6, 5, 1}))
}

func 두큐합같게만들기(queue1 []int, queue2 []int) int {

	middleValue := getMiddleValue(queue1, queue2)
	var result int

	for {
		if getSum(queue1) == middleValue && getSum(queue2) == middleValue {
			return result
		}
	}

	return -1
}

func getSum(sl []int) int {

	var result int

	for _, v := range sl {
		result += v
	}

	return result

}

func getMiddleValue(queue1 []int, queue2 []int) int {

	var sumSl = append(queue1, queue2...)

	return getSum(sumSl) / 2

}
