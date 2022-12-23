package main

import "fmt"

func main() {
	fmt.Println(크기가작은부분문자열("3141592", "271"))
}

func 크기가작은부분문자열(t string, p string) int {

	var result int
	var ln = len([]rune(t))
	for i := 0; i < ln; i++ {

		if i > ln-3 {
			break
		}

		if t[i:i+1] == "0" {
			continue
		}

		s := t[i : i+3]

		if s < p {
			result++
		}
	}

	return result
}
