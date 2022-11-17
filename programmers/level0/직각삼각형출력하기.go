package main

import "fmt"

func main() {
	직각삼각형출력하기(3)
}

func 직각삼각형출력하기(n int) {

	var stars = 1

	for i := 0; i < n; i++ {

		for j := 0; j < stars; j++ {
			fmt.Print("*")
		}

		stars++

		if i < n-1 {
			fmt.Println()
		}
	}

}
