package main

import (
	"bufio"
	"fmt"
	"os"
)

type Stack struct {
	items []int
}

func main() {

	var n int
	var read = bufio.NewReader(os.Stdin)
	fmt.Fscanln(read, &n)

	s := Stack{}

	for i := 0; i < n; i++ {
		var cmd string
		var v int
		fmt.Fscanln(read, &cmd, &v)

		s.execute(cmd, v)
	}
}

func (s *Stack) execute(cmd string, v int) {

	switch cmd {
	case "push":
		s.push(v)
	case "pop":
		fmt.Println(s.pop())
	case "size":
		fmt.Println(s.size())
	case "empty":
		fmt.Println(s.empty())
	case "top":
		fmt.Println(s.top())
	}

}

func (s *Stack) push(x int) {
	s.items = append(s.items, x)
}

func (s *Stack) pop() int {
	if len(s.items) == 0 {
		return -1
	}

	prop := s.items[len(s.items)-1]

	s.items = s.items[0 : len(s.items)-1]

	return prop
}

func (s *Stack) size() int {
	return len(s.items)
}

func (s *Stack) empty() int {
	if len(s.items) == 0 {
		return 1
	}

	return 0
}

func (s *Stack) top() int {
	if len(s.items) == 0 {
		return -1
	}

	return s.items[len(s.items)-1]
}
