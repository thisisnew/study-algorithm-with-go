package main

import "fmt"

func main() {
	fmt.Println(두큐합같게만들기([]int{3, 3, 3, 3}, []int{3, 3, 21, 3}))
}

type Queue struct {
	items []int
}

func (q *Queue) pop() int {

	if len(q.items) == 0 {
		return -1
	}

	p := q.items[0]

	q.items = q.items[1:len(q.items)]

	return p
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

func (q *Queue) sum() int {

	var result float64

	for _, v := range q.items {
		result += float64(v)
	}

	return int(result)

}

func 두큐합같게만들기(queue1 []int, queue2 []int) int {

	q1 := Queue{queue1}
	q2 := Queue{queue2}
	q1Sum := q1.sum()
	q2Sum := q2.sum()

	middleValue, remain := getMiddleValue(q1Sum, q2Sum)
	limitCount := 2 * (len(q1.items) + len(q2.items))

	if remain == 1 {
		return -1
	}

	var result int

	for {

		if limitCount < 0 {
			return -1
		}

		switch {
		case q2Sum > q1Sum:
			p := q2.pop()
			q1.push(p)
			q1Sum += p
			q2Sum -= p
		case q2Sum < q1Sum:
			p := q1.pop()
			q2.push(p)
			q2Sum += p
			q1Sum -= p
		default:
			if q1Sum == middleValue && q2Sum == middleValue {
				return result
			}

			return -1
		}

		result++
		limitCount--
	}
}

func getMiddleValue(q1Sum, q2Sum int) (int, int) {
	sum := q1Sum + q2Sum
	return sum / 2, sum % 2
}
