package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	var read = bufio.NewReader(os.Stdin)
	fmt.Fscanln(read, &n)

	for i := 0; i < n; i++ {
		var input string
		fmt.Fscanln(read, &input)

		isAddition, a, b := isAddition(input)

		if isAddition {
			fmt.Println(a + b)
			continue
		}

		fmt.Println("skipped")

	}
}

func isAddition(input string) (bool, int, int) {

	addition := strings.Split(input, "+")

	if len(addition) < 2 {
		return false, 0, 0

	}

	a, _ := strconv.Atoi(addition[0])
	b, _ := strconv.Atoi(addition[1])

	return true, a, b

}
