package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var n int
	var read = bufio.NewReader(os.Stdin)

	fmt.Fscan(read, &n)

	for i := 0; i < n; i++ {
		var m int
		fmt.Fscan(read, &m)

		var sl = make([]string, m)

		for j := 0; j < m; j++ {
			var w string
			fmt.Fscan(read, &w)
			sl[i] = w
		}

		var result int

		for _, word := range sl {
			if isSheep(word) {
				result++
			}
		}

		fmt.Println(fmt.Sprintf("Case %v: This list contains %v sheep.", i+1, result))
	}

}

func isSheep(word string) bool {

	if word == "sheep" {
		return true
	}

	return false
}
