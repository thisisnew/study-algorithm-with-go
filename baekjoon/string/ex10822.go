package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var input string
	var read = bufio.NewReader(os.Stdin)
	fmt.Fscan(read, &input)

	stringArray := strings.Split(input, ",")

	var total int

	for _, item := range stringArray {
		num, _ := strconv.Atoi(item)

		total += num
	}

	fmt.Println(total)
}
