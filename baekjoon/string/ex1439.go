package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var s string
	var read = bufio.NewReader(os.Stdin)

	fmt.Fscanln(read, &s)

	var zeroCnt int
	var oneCnt int
	var prev = 0

	//0 -> 1변환
	for i, r := range s {
		if string(r) == "0" {
			if i == 0 || (i == prev+1) {
				prev = i
				zeroCnt++
			}
		}

	}

	if zeroCnt == 0 || zeroCnt == len(s) {
		fmt.Println(0)
		return
	}

	//1 -> 0
	prev = 0
	for i, r := range s {
		if string(r) == "1" {
			if i == 0 || (i == prev+1) {
				oneCnt++
			}
		}
		prev = i
	}

	if zeroCnt < oneCnt {
		fmt.Println(zeroCnt)
	} else {
		fmt.Println(oneCnt)
	}
}
