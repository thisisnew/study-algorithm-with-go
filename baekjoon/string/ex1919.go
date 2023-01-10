package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var a string
	var b string
	var read = bufio.NewReader(os.Stdin)

	fmt.Fscanln(read, &a)
	fmt.Fscanln(read, &b)

	for i := 0; i < len([]rune(a)); i++ {

		s := a[i : i+1]

		for j := 0; j < len([]rune(b)); j++ {

			if s != b[j:j+1] {
				continue
			}

			a = a[0:i] + "*" + a[i+1:]
			b = b[0:j] + "*" + b[j+1:]
			break
		}
	}

	a = strings.ReplaceAll(a, "*", "")
	b = strings.ReplaceAll(b, "*", "")

	fmt.Println(len([]rune(a)) + len([]rune(b)))
	return
}
