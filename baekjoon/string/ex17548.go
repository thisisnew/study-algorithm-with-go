package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var read = bufio.NewReader(os.Stdin)

	var input string
	fmt.Fscanln(read, &input)

	ln := len([]rune(input))
	fmt.Println(fmt.Sprintf("%s%s%s", input[0:1], input[1:ln-1], input[ln:ln+1]))

}
