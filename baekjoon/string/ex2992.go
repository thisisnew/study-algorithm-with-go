package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var x string
	var read = bufio.NewReader(os.Stdin)
	fmt.Fscanln(read, &x)

}
