package main

import "fmt"

func main() {
	fmt.Println(양꼬치(10, 3))
}

func 양꼬치(n int, k int) int {
	return n*12000 + 2000*(k-(n/10))
}
