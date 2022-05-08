package main

import "fmt"

func main() {
	fmt.Println(최대공약수와최소공배수(2, 5))
}

func 최대공약수와최소공배수(n int, m int) []int {

	if n > m {
		temp := n
		n = m
		m = temp
	}

	return []int{gcd(n, m), lcm(n, m)}
}

func gcd(n, m int) int {

	var r int

	for {
		if n == 0 {
			break
		}

		r = m % n
		m = n
		n = r
	}

	return m
}

func lcm(n, m int) int {
	return n * m / gcd(n, m)
}
