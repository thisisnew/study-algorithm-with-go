package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var read = bufio.NewReader(os.Stdin)
	var input string

	fmt.Fscanln(read, &input)

	rns := []rune(input)

	var result int

	for i := 0; i < len(rns); i++ {
		result += int(rns[i]-64) * int(math.Pow(26, float64(len(rns)-i-1)))
	}

	fmt.Println(result)
}
