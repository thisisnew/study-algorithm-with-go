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

	if len([]rune(input)) == 1 {
		fmt.Println(int([]rune(input)[0]-'A') + 1)
		return
	}

}
