package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var read = bufio.NewReader(os.Stdin)
	var drm string
	fmt.Fscanln(read, &drm)

	pre, post := divideDrm(drm)

	nPre := getNewDrmToken(pre, sumTokenAsciiValue(pre))
	nPost := getNewDrmToken(post, sumTokenAsciiValue(post))

	fmt.Println(nPre)
	fmt.Println(nPost)
}

func divideDrm(drm string) (string, string) {
	return drm[:len([]rune(drm))/2], drm[len([]rune(drm))/2:]
}

func sumTokenAsciiValue(input string) int {

	var result int

	for _, n := range input {
		result += int(n)
	}

	return result
}

func getNewDrmToken(old string, asciiValue int) string {

	var result strings.Builder

	for _, o := range old {

		v := int(o) + asciiValue - 65

		result.WriteString(string(v))
	}

	return result.String()
}
