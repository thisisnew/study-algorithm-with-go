package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var input string
	var read = bufio.NewReader(os.Stdin)

	fmt.Fscanln(read, &input)

	nums := strings.Split(input, ",")

	fmt.Println(len(nums))
}
