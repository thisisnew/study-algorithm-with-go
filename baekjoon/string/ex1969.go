package main

import (
	"bufio"
	"fmt"
	"os"
)

var fDna string

func main() {
	var n, m int
	var read = bufio.NewReader(os.Stdin)
	fmt.Fscanln(read, &n, &m)
	fmt.Fscanln(read, &fDna)

	var minHd int
	var mDna string
	var isGetHd bool

	for i := 0; i < n-1; i++ {
		var dna string
		fmt.Fscanln(read, &dna)

		hd := getHammingDistance(dna)

		if !isGetHd {
			isGetHd = true
			minHd = hd
			mDna = dna
			continue
		}

		if hd < minHd {
			minHd = hd
			mDna = dna
		}
	}

	fmt.Println(mDna)
}

func getHammingDistance(dna string) int {
	var result int

	for i := 0; i < len([]rune(dna)); i++ {
		if dna[i:i+1] != fDna[i:i+1] {
			result++
		}
	}

	return result
}
