package queue

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type Queue struct {
	props []int
}

func main() {
	var n int
	var read = bufio.NewReader(os.Stdin)

	for i := 0; i < n; i++ {
		var input string
		fmt.Fscanln(read, &input)

	}
}

func (q *Queue) push(x int) []int {
	q.props = append(q.props, x)
}

func (q *Queue) pop() (int, error) {
	if len(q.props) == 0 {
		return -1, errors.New("")
	}

	x, props := q.props[0], q.props[1:]

	q.props = props

	return x, nil
}

func (q *Queue) size() int {
	return len(q.props)
}

func (q *Queue) empty() int {
	if len(q.props) == 0 {
		return 1
	}

	return 0
}

func (q *Queue) front() int {
	if len(q.props) == 0 {
		return -1
	}

	return q.props[0]
}

func (q *Queue) back() int {
	if len(q.props) == 0 {
		return -1
	}

	return q.props[len(q.props)-1]
}
