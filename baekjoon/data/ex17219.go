package main

import (
	"bufio"
	"fmt"
	"os"
)

type urlPwInfo struct {
	url string
	pw  string
}

func main() {

	var read = bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscanln(read, &n, &m)

	var urls = make([]urlPwInfo, n)
	for i := 0; i < n; i++ {
		var url string
		var pw string
		fmt.Fscanln(read, &url, &pw)

		urls[i] = urlPwInfo{url, pw}
	}

	for i := 0; i < m; i++ {
		var url string
		fmt.Fscanln(read, &url)

		for j := 0; j < len(urls); j++ {
			if urls[j].url != url {
				continue
			}

			fmt.Println(urls[j].pw)
			break

		}
	}

}
