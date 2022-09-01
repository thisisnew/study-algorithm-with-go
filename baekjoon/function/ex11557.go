package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var read = bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscanln(read, &t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscanln(read, &n)
		fmt.Println(getMaximumDrunkSchool(read, n))
	}

}

func getMaximumDrunkSchool(read *bufio.Reader, n int) string {

	var alcohol int
	var school string

	for i := 0; i < n; i++ {
		text, _, _ := read.ReadLine()

		sl := strings.Split(string(text), " ")
		al, _ := strconv.Atoi(sl[1])

		if al > alcohol {
			alcohol = al
			school = sl[0]
		}
	}

	return school
}
