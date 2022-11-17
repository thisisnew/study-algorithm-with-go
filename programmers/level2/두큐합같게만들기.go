package main

import "fmt"

func main() {
	fmt.Println(두큐합같게만들기([]int{1, 2, 1, 2}, []int{1, 10, 1, 2}))
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

	middleValue, remain := getMiddleValue(q1, q2)

	if remain == 1 {
		return -1
	}

	var result int

	for {
		switch {
		case q2.sum() > q1.sum():
			p := q2.pop()
			q1.push(p)
		case q2.sum() < q1.sum():
			p := q1.pop()
			q2.push(p)
		default:
			if q1.sum() == middleValue && q2.sum() == middleValue {
				return result
			}

			return -1
		}

		result++
	}
}

func getMiddleValue(q1, q2 Queue) (int, int) {

	var q = Queue{}
	q.items = append(q1.items, q2.items...)

	sum := q.sum()

	return sum / 2, sum % 2
}
