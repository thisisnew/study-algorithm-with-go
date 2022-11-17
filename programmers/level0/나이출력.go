package main

import "fmt"

func main() {
	fmt.Println(나이출력(40))
}

func 나이출력(age int) int {
	return 2022 - age + 1
}
