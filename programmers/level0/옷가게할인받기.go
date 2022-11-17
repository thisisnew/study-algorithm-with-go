package main

import "fmt"

func main() {
	fmt.Println(옷가게할인받기(150000))
}

func 옷가게할인받기(price int) int {

	switch {
	case price > 500000:
		return int(float32(price) * 0.8)
	case price > 300000:
		return int(float32(price) * 0.9)
	case price > 100000:
		return int(float32(price) * 0.95)
	}

	return price
}
