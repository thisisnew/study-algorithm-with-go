package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var read = bufio.NewReader(os.Stdin)

	var input string
	fmt.Fscanln(read, &input)

	ln := len([]rune(input))
	fmt.Println(fmt.Sprintf("%s%s%s", input[0:1], getEsTwice(ln-2), input[ln-1:ln]))

}

func getEsTwice(ln int) string {

	var result strings.Builder

	for i := 0; i < 2*ln; i++ {
		result.WriteString("e")
	}

	return result.String()
}
