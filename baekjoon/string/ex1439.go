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

		switch string(r) {
		case "0":
			if isReverse {
				continue
			}

			isReverse = true
		case "1":
			if i == 0 {
				zeroCnt++
				continue
			}

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

		switch string(r) {
		case "0":
			if i == 0 {
				oneCnt++
				continue
			}

			if !isReverse {
				continue
			}

			oneCnt++
			isReverse = false
		case "1":
			if isReverse {
				continue
			}

			isReverse = true
		}
	}

	if zeroCnt < oneCnt {
		fmt.Println(zeroCnt)
	} else {
		fmt.Println(oneCnt)
	}
}
