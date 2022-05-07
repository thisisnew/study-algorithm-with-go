package main

import "fmt"

func main() {
	n := []int{
		1, 2, 3, 4,
	}

	fmt.Println(평균구하기(n))
}

func 평균구하기(arr []int) float64 {

	var result float64

	for _, n := range arr {
		result += float64(n)
	}

	return result / float64(len(arr))

}
