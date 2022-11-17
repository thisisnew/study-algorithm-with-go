package main

import "fmt"

func main() {
	fmt.Println(배열원소의길이([]string{"We", "are", "the", "world!"}))
}

func 배열원소의길이(strlist []string) []int {

	var result = make([]int, len(strlist))

	for i, str := range strlist {
		result[i] = len([]rune(str))
	}

	return result
}
