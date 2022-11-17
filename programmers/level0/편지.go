package main

import "fmt"

func main() {
	fmt.Println(편지("happy birthday!"))
}

func 편지(message string) int {
	return len([]rune(message)) * 2
}
