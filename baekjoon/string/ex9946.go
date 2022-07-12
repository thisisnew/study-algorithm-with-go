package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var read = bufio.NewReader(os.Stdin)
	var sl []string
	var endCnt = 0

	for {
		var f string
		fmt.Fscanln(read, &f)

		if f == "END" {
			endCnt++
		} else {
			endCnt = 0
		}

		if endCnt == 2 {
			break
		}

		sl = append(sl, f)
	}

	var i = 0
	var c = 1
	for i < len(sl) {

		var f = sl[i]
		var s = sl[i+1]

		r := getReversedWord(s)

		if f == r {
			fmt.Println(fmt.Sprintf("Case %v: same", c))
		} else {
			fmt.Println(fmt.Sprintf("Case %v: different", c))
		}

		c++
		i += 2
	}

}

func getReversedWord(s string) string {

	var result strings.Builder

	for i := len([]rune(s)) - 1; i >= 0; i-- {
		result.WriteString(s[i : i+1])
	}

	return result.String()
}
