package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n, m int
	var read = bufio.NewReader(os.Stdin)
	fmt.Fscanln(read, &n, &m)

	var dnas = make([]string, n)

	for i := 0; i < n; i++ {
		var dna string
		fmt.Fscanln(read, &dna)
		dnas[i] = dna
	}

	var hammingDistance int
	var result strings.Builder

	for _, dna := range dnas {

		var a, c, g, t int
		var max = 0
		var maxC string

		for _, token := range dna {

			switch token {
			case 'A':
				a++
				if max < a {
					max = a
					maxC = "A"
				}
			case 'C':
				c++
				if max < c {
					max = c
					maxC = "C"
				}
			case 'G':
				g++
				if max < g {
					max = g
					maxC = "G"
				}
			case 'T':
				t++
				if max < t {
					max = t
					maxC = "T"
				}
			}
		}

		result.WriteString(maxC)
		hammingDistance = n - max
	}

	fmt.Println(result.String())
	fmt.Println(hammingDistance)
}
