package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(옷가게할인받기(150000))
}

func 옷가게할인받기(price int) int {

	dc := 1.0

	switch {
	case price >= 500000:
		dc = 0.8
	case price >= 300000:
		dc = 0.9
	case price >= 100000:
		dc = 0.95
	}

	return int(math.Floor(float64(price) * dc))
}
