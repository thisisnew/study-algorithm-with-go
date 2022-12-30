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

	total, totalLen := getTotalNumberString()
	nLen := len([]rune(n))
	var sb strings.Builder

	for i := 0; i < totalLen-nLen; i++ {

		if total[i:i+1] != n[0:1] {
			continue
		}

		sb.WriteString(total[i : i+1])

		for j := i + 1; j <= totalLen; j++ {

			if sb.Len() < nLen {
				sb.WriteString(total[j : j+1])
				continue
			}

			if sb.String() != n {
				sb.Reset()
				break
			}

			fmt.Println(i + 1)
			return
		}
	}

}

func getTotalNumberString() (string, int) {

	var result strings.Builder

	for i := 1; i <= 10000; i++ {

		d := strconv.Itoa(i)
		result.WriteString(d)
	}

	return result.String(), result.Len()
}
