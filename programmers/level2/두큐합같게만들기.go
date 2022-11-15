package main

import "fmt"

func main() {
	fmt.Println(두큐합같게만들기([]int{3, 2, 7, 2}, []int{4, 6, 5, 1}))
}

type Queue struct {
	items []int
}

func (q *Queue) pop() int {

	if len(q.items) == 0 {
		return -1
	}

	pack := q.items[len(q.items)-1]

	q.items = q.items[0 : len(q.items)-1]

	return pack
}

func (q *Queue) top() int {

	if len(q.items) == 0 {
		return -1
	}

	return q.items[len(q.items)-1]
}

func (q Queue) empty() bool {
	return len(q.items) == 0
}

func (q *Queue) push(x int) {
	q.items = append(q.items, x)
}

func 두큐합같게만들기(queue1 []int, queue2 []int) int {

	middleValue, remain := getMiddleValue(queue1, queue2)

	if remain == 1 {
		return -1
	}

	var result int
	var ln = len(queue1) * 2

	for {

		result++

		if getSum(queue1) < getSum(queue2) {
			queue, p := popProp(queue2)
			queue2 = queue
			queue1 = pushProp(queue1, p)
		} else {
			queue, p := popProp(queue1)
			queue1 = queue
			queue2 = pushProp(queue2, p)
		}

		if isSameQueueValue(queue1, queue2, middleValue) {
			return result
		}

		if result > ln*2 {
			return -1
		}

	}
}

func isSameQueueValue(queue1, queue2 []int, middleValue int) bool {
	return getSum(queue1) == middleValue && getSum(queue2) == middleValue
}

func getSum(sl []int) int {

	var result float64

	for _, v := range sl {
		result += float64(v)
	}

	return int(result)

}

func getMiddleValue(queue1 []int, queue2 []int) (int, int) {

	var sumSl = append(queue1, queue2...)

	sum := getSum(sumSl)

	return sum / 2, sum % 2

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
