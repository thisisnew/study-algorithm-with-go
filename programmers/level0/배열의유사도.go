package main

import "fmt"

func main() {
	fmt.Println(배열의유사도([]string{"a", "b", "c"}, []string{"com", "b", "d", "p", "c"}))
}

func 배열의유사도(s1 []string, s2 []string) int {

	var result int

	for _, s := range s1 {
		for _, p := range s2 {
			if s == p {
				result++
				break
			}
		}
	}

	return result
}
