package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(게임369(3))
}

func 게임369(order int) int {

	orderStr := strconv.Itoa(order)
	var result int

	for _, os := range orderStr {
		switch string(os) {
		case "3", "6", "9":
			result++
		}
	}

	return result
}
