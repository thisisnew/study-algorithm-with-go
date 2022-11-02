package main

import "fmt"

func main() {
	fmt.Println(두큐합같게만들기([]int{1, 1}, []int{1, 5}))
}

func 두큐합같게만들기(queue1 []int, queue2 []int) int {

	middleValue := getMiddleValue(queue1, queue2)
	var result int
	var isQueue1Finished = false

	for {

		if !isQueue1Finished {
			queue, p := popProp(queue1)
			queue1 = queue
			queue2 = pushProp(queue2, p)
		} else {
			queue, p := popProp(queue2)
			queue2 = queue
			queue1 = pushProp(queue1, p)
		}

		if len(queue1) == 0 {
			isQueue1Finished = true
		}

		result++

		if isSameQueueValue(queue1, queue2, middleValue) {
			return result
		}

		if len(queue2) == 0 {
			return -1
		}
	}
}

func isSameQueueValue(queue1, queue2 []int, middleValue int) bool {
	return getSum(queue1) == middleValue && getSum(queue2) == middleValue
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

func pushProp(q []int, prop int) []int {
	q = append(q, prop)

	return q
}

func popProp(q []int) ([]int, int) {

	if len(q) == 0 {
		return q, 0
	}

	var prop = q[0]

	q = q[1:]

	return q, prop
}
