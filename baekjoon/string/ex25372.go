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

	var pws = make([]string, n)

	for i := 0; i < n; i++ {
		var pw string
		fmt.Fscanln(read, &pw)

		pws[i] = pw
	}

	for _, pw := range pws {
		fmt.Println(isValidPassword(pw))
	}

}

func isValidPassword(pw string) string {

	if len([]rune(pw)) > 9 || len([]rune(pw)) < 6 {
		return "no"
	}

	return "yes"
}
