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

	var gcf int

	for i := 1; i <= m; i++ {

		if m%i != 0 {
			continue
		}

		if i > n {
			break
		}

		gcf = i

	}

	return []int{gcf, gcf * (n / gcf) * (m / gcf)}
}
