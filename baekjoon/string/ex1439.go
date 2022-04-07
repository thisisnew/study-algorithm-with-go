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

	var isReverse bool

	//0 -> 1변환
	for i, r := range s {

		if i == 0 {
			if string(r) == "1" {
				zeroCnt++
			}
		}

		if string(r) == "0" {
			if isReverse {
				continue
			}
			isReverse = true
		}

		if string(r) == "1" {
			if !isReverse {
				continue
			}
			zeroCnt++
			isReverse = false
		}

	}

	if zeroCnt == 0 || zeroCnt == len(s) {
		fmt.Println(0)
		return
	}

	//1 -> 0
	for i, r := range s {

		if i == 0 {
			if string(r) == "0" {
				oneCnt++
			}
		}

		if string(r) == "1" {
			if isReverse {
				continue
			}
			isReverse = true
		}

		if string(r) == "0" {
			if !isReverse {
				continue
			}
			oneCnt++
			isReverse = false
		}
	}

	if zeroCnt < oneCnt {
		fmt.Println(zeroCnt)
	} else {
		fmt.Println(oneCnt)
	}
}
