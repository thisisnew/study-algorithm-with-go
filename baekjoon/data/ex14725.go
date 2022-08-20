package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var read = bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanln(read, &n)

	var result map[string][]string

	for i := 0; i < n; i++ {
		text, _, _ := read.ReadLine()
		sl := strings.Split(string(text), " ")

		pr, ok := result[sl[1]]

		if !ok {
			var props = make([]string, len(sl)-2)

			for j := 0; j < len(sl)-2; j++ {
				props[j] = fmt.Sprintf("%s%s", generateBars(j), sl[j+2])
			}

			result[sl[1]] = props
		} else {

		}

	}

}

func generateBars(j int) string {

	var result strings.Builder

	for i := 0; i < 2*(j+1); i++ {
		result.WriteString("-")
	}

	return result.String()
}
