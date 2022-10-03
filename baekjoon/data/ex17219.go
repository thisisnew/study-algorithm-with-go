package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var read = bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscanln(read, &n, &m)

	var urls = map[string]string{}
	for i := 0; i < n; i++ {
		var url string
		var pw string
		fmt.Fscanln(read, &url, &pw)

		urls[url] = pw
	}

	for i := 0; i < m; i++ {
		var url string
		fmt.Fscanln(read, &url)
		fmt.Println(urls[url])
	}

}
