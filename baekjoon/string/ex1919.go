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

			a = strings.Replace(a, s, "*", i)
			b = strings.Replace(b, s, "*", j)
			break
		}
	}

	a = strings.ReplaceAll(a, "*", "")
	b = strings.ReplaceAll(b, "*", "")

	fmt.Println(len([]rune(a)) + len([]rune(b)))
	return
}
