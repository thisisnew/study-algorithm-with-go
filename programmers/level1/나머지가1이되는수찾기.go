package main

import "fmt"

func main() {
	fmt.Println(나머지가1이되는수찾기(10))
}

func 나머지가1이되는수찾기(n int) int {

	var result int

	for i := 2; i < n; i++ {

		if n%i != 1 {
			continue
		}

		result = i
		break
	}

	return result
}
