package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	var input string
	var read = bufio.NewReader(os.Stdin)
	fmt.Fscanln(read, &input)

	data, _ := base64.StdEncoding.DecodeString(input)

	fmt.Println(string(data))
}
