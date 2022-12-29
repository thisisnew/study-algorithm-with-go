package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var n int
	var read = bufio.NewReader(os.Stdin)
	fmt.Fscanln(read, &n)

	var sb strings.Builder
	var temp strings.Builder
	var div = 10
	var isMatching bool
	var nStr = strconv.Itoa(n)

	for i := 1; i <= 100000; i++ {

		if i == div*10 {
			div *= 10
		}

		sb.WriteString(strconv.Itoa(i))

		d := strconv.Itoa(i / div)

		if d == nStr[0:1] {
			isMatching = true
		}

		if isMatching {
			if temp.Len() == len([]rune(nStr)) {
				if temp.String() != nStr {
					temp.Reset()
				}

				break
			}

			temp.WriteString(strconv.Itoa(i))
		}

	}

}
