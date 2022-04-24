package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Queue struct {
	props []int
}

func main() {
	var n int
	var read = bufio.NewReader(os.Stdin)
	fmt.Fscanln(read, &n)

	q := Queue{}

	for i := 0; i < n; i++ {
		text, _, _ := read.ReadLine()

		cm, x, _ := getCommand(string(text))
		q.execute(cm, x)
	}
}

func (q *Queue) execute(cm string, x int) {
	switch cm {
	case "front":
		fmt.Println(q.front())
	case "back":
		fmt.Println(q.back())
	case "empty":
		fmt.Println(q.empty())
	case "size":
		fmt.Println(q.size())
	case "pop":
		v, _ := q.pop()
		fmt.Println(v)
	case "push":
		q.push(x)
	}
}

func getCommand(command string) (string, int, error) {

	c := strings.Split(command, " ")

	var fn = c[0]
	var x int

	if len(c) == 1 {
		return fn, x, errors.New("not exist value")
	}

	vl, _ := strconv.Atoi(c[1])
	x = vl

	return fn, x, nil
}

func (q *Queue) push(x int) {
	q.props = append(q.props, x)
}

func (q *Queue) pop() (int, error) {
	if len(q.props) == 0 {
		return -1, errors.New("empty")
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
