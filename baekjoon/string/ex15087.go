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

	fmt.Println(getDrmResultBetweenTwoWords(nPre, nPost))
}

func divideDrm(drm string) (string, string) {
	return drm[:len([]rune(drm))/2], drm[len([]rune(drm))/2:]
}

func sumTokenAsciiValue(input string) int {

	var result int

	for _, n := range input {
		result += int(n) - 65
	}

	return result
}

func getNewDrmToken(old string, asciiValue int) string {

	var result strings.Builder

	for _, o := range old {
		result.WriteString(string(rotateDrmCharacter(int(o) + asciiValue)))
	}

	return result.String()
}

func rotateDrmCharacter(v int) int {

	var result = v

	for {
		if result <= 90 {
			break
		}

		result -= 26
	}

	return result
}

func getDrmResultBetweenTwoWords(first, second string) string {
	var result strings.Builder

	for i, f := range first {
		result.WriteString(string(rotateDrmCharacter(int(f) + int([]rune(second)[i]-65))))
	}

	return result.String()
}
