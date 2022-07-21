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

}

func divideDrm(drm string) (string, string) {
	return drm[:len([]rune(drm))/2+1], drm[len([]rune(drm))/2+1:]
}

func sumTokenAsciiValue(input string) int {

	var result int32

	for _, n := range input {
		result += n
	}

	return int(result)
}

func getNewDrmWord(old string, asciiValue int) string {

	var result strings.Builder

	for _, o := range old {

	}

}
