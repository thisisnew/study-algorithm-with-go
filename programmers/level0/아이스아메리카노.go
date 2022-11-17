package main

import "fmt"

func main() {
	fmt.Println(아이스아메리카노(5500))
}

func 아이스아메리카노(money int) []int {
	return []int{money / 5500, money % 5500}
}
