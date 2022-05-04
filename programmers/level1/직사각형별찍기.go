package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	직사각형별찍기(a, b)
}

func 직사각형별찍기(a, b int) {

	for i := 0; i < b; i++ {
		for j := 0; j < a; j++ {
			fmt.Print("*")
		}

		if i != b-1 {
			fmt.Println()
		}
	}
}
