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
	var idx int

	for idx < m {

		var a, c, g, t int
		var max = 0
		var maxC string

		for _, dna := range dnas {
			switch dna[idx : idx+1] {
			case "A":
				a++
			case "C":
				c++
			case "G":
				g++
			case "T":
				t++
			}
		}

		if max < a {
			max = a
			maxC = "A"
		}

		if max < c {
			max = c
			maxC = "C"
		}

		if max < g {
			max = g
			maxC = "G"
		}

		if max < t {
			max = t
			maxC = "T"
		}

		result.WriteString(maxC)
		hammingDistance += n - max
		idx++
	}

	fmt.Println(result.String())
	fmt.Println(hammingDistance)
}
