package main

import "fmt"

func main() {
	fmt.Println(count7([]int{7, 17, 77}))
}

func count7(array []int) int {
	var result int

	for i := 0; i < len(array); i++ {

		var n = array[i]
		var share int
		var reminder int

		for {
			share = n / 10
			reminder = n % 10

			if reminder == 7 {
				result++
			}

			if share == 0 {
				break
			}

			n = share
		}
	}

	return result
}
