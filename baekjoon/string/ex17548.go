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

}
