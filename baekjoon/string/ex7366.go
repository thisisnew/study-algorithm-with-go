package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var n int
	var read = bufio.NewReader(os.Stdin)

	fmt.Fscan(read, &n)

	for i := 0; i < n; i++ {
		var m int
		fmt.Fscan(read, &m)
		text, _, _ := read.ReadLine()

		var result int

		for _, word := range strings.Split(string(text), " ") {
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
