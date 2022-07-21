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

}

func divideDrm(drm string) (string, string) {
	return drm[:len([]rune(drm))/2+1], drm[len([]rune(drm))/2+1:]
}

func sumTokenAsciiValue(input string) int32 {

	var result int32

	for _, n := range input {
		result += n
	}

	return result
}

func getNewDrmToken(old string, asciiValue int32) string {

	var result strings.Builder

	for _, o := range old {
		result.WriteString(string(o + asciiValue))
	}

	return result.String()
}

func getNewDrmWord(old string, token string) string {

}
