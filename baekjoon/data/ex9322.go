package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var read = bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscanln(read, &t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscanln(read, &n)

		var fPublicKey []string
		var sPublicKey []string
		var crypto []string

		for j := 0; j < 3; j++ {
			text, _, _ := read.ReadLine()
			values := strings.Split(string(text), " ")

			switch i {
			case 0:
				fPublicKey = values
			case 1:
				sPublicKey = values
			case 2:
				crypto = values
			}
		}

		printPlainTextFromThreeProps(fPublicKey, sPublicKey, crypto)
	}
}

func printPlainTextFromThreeProps(fPublicKey, sPublicKey, crypto []string) {

}
