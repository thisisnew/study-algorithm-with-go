package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

		text, _, _ := read.ReadLine()

		command := strings.Split(string(text), " ")

		s.execute(command)
	}
}

func (s *Stack) execute(command []string) {

	var cmd = command[0]
	var value string

	if len(command) > 1 {
		value = command[1]
	}

	switch cmd {
	case "push":
		v, _ := strconv.Atoi(value)
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
