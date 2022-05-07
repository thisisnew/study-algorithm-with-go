package main

import (
	"fmt"
)

func main() {

	seoul := []string{"Jane", "Kim"}

	fmt.Println(서울에서김서방찾기(seoul))
}

func 서울에서김서방찾기(seoul []string) string {

	var idx int

	for i, item := range seoul {

		if item != "Kim" {
			continue
		}

		idx = i
		break
	}

	return fmt.Sprintf("김서방은 %v에 있다", idx)
}
