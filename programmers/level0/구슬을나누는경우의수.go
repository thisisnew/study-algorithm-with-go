package main

import "fmt"

func main() {
	fmt.Println(구슬을나누는경우의수(5, 3))
}

func 구슬을나누는경우의수(n int, m int) int {
	return int(fac(n) / (fac(n-m) * fac(m)))
}

func fac(ball int) float64 {

	if ball <= 1 {
		return float64(ball)
	}

	return fac(ball-1) * float64(ball)
}
