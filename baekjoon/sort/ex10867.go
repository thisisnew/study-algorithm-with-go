package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var n int
	var read = bufio.NewReader(os.Stdin)
	fmt.Fscanln(read, &n)

	var arr []int

	for i := 0; i < n; i++ {
		var num int
		fmt.Fscanln(read, &num)
		arr = append(arr, num)
	}

}
