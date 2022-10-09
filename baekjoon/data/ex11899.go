package main

import (
	"bufio"
	"fmt"
	"os"
)

type braceStack []string

func (b *braceStack) IsEmpty() bool {
	return len(*b) == 0
}

func (b *braceStack) Push(s string) {
	*b = append(*b, s)
}

func (b *braceStack) Pop() string {

	if b.IsEmpty() {
		return ""
	}

	el := (*b)[len(*b)-1]
	*b = (*b)[:len(*b)-1]
	return el
}

func main() {

	var read = bufio.NewReader(os.Stdin)
	var s string
	fmt.Fscanln(read, &s)

	var result int
	var bs braceStack

	for _, b := range s {

		s := string(b)

		if s == "(" {
			bs.Push(s)
		}

		if s == ")" {
			if bs.IsEmpty() {
				result++
			} else {
				bs.Pop()
			}
		}
	}

	fmt.Println(result + len(bs))
}
