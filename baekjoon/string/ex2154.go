package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var n string
	var read = bufio.NewReader(os.Stdin)
	fmt.Fscanln(read, &n)

	var sb strings.Builder
	var temp strings.Builder

	for i := 1; i <= 10000-len([]rune(n)); i++ {

		d := strconv.Itoa(i)
		temp.WriteString(d)

		if d[0:1] != n[0:1] {
			continue
		}

		sb.WriteString(d)

		for j := i + 1; j <= 10000; j++ {
			if sb.Len() == len([]rune(n)) {
				if sb.String() == n {
					fmt.Println(temp.Len() - (len([]rune(n)) - 1))
					return
				}

				sb.Reset()
				break
			}

			sb.WriteString(strconv.Itoa(j))
		}
	}

}
