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

	fmt.Fscanln(read, &n)

	for i := 0; i < n; i++ {
		text, _, _ := read.ReadLine()

		sl := strings.Split(string(text), " ")

		var result = make([]string, len(sl))

		for j := len(sl) - 1; j >= 0; j-- {
			result[len(sl)-1-j] = sl[j]
		}

		fmt.Println(fmt.Sprintf("Case #%v: %s", i+1, strings.Join(result, " ")))
	}

}
