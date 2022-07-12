package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var read = bufio.NewReader(os.Stdin)
	var i = 1
	for {

		var f string
		var s string
		fmt.Fscanln(read, &f, &s)

		if f == "END" && s == "END" {
			break
		}

		if f == getReversedWord(s) {
			fmt.Println(fmt.Sprintf("Case %v: same", i))
		} else {
			fmt.Println(fmt.Sprintf("Case %v: different", i))
		}

		i++
	}

}

func getReversedWord(s string) string {

	var result strings.Builder

	for i := len([]rune(s)) - 1; i >= 0; i-- {
		result.WriteString(s[i : i+1])
	}

	return result.String()
}
