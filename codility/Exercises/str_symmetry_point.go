package main

import "fmt"

func main() {
	fmt.Println(StrSymmetryPoint("racecar"))
}

func StrSymmetryPoint(S string) int {

	if len([]rune(S)) <= 1 {
		return 0
	}

	return len([]rune(S)) / 2
}
