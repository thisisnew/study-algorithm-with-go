package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var read = bufio.NewReader(os.Stdin)
	var drm string
	fmt.Fscanln(read, &drm)

	pre, post := divideDrm(drm)

	//nPre := getNewDrmToken(pre, sumTokenAsciiValue(pre))
	//nPost := getNewDrmToken(post, sumTokenAsciiValue(post))

	getNewDrmToken(pre, sumTokenAsciiValue(pre))
	getNewDrmToken(post, sumTokenAsciiValue(post))
}

func divideDrm(drm string) (string, string) {
	return drm[:len([]rune(drm))/2], drm[len([]rune(drm))/2:]
}

func sumTokenAsciiValue(input string) int {

	var result int

	for _, n := range input {
		result += getDrmIntValues(string(n))
	}

	return result
}

func getNewDrmToken(old string, asciiValue int) string {

	var result strings.Builder

	for _, o := range old {

		v := getDrmIntValues(string(o)) + asciiValue

		if v > 25 {
			v = v - 25
		}

		result.WriteString(strconv.Itoa(v))
	}

	return result.String()
}

func getDrmIntValues(d string) int {

	switch d {
	case "A":
		return 0
	case "B":
		return 1
	case "C":
		return 2
	case "D":
		return 3
	case "E":
		return 4
	case "F":
		return 5
	case "G":
		return 6
	case "H":
		return 7
	case "I":
		return 8
	case "J":
		return 9
	case "K":
		return 10
	case "L":
		return 11
	case "M":
		return 12
	case "N":
		return 13
	case "O":
		return 14
	case "P":
		return 15
	case "Q":
		return 16
	case "R":
		return 17
	case "S":
		return 18
	case "T":
		return 19
	case "U":
		return 20
	case "V":
		return 21
	case "W":
		return 22
	case "X":
		return 23
	case "Y":
		return 24
	case "Z":
		return 25
	}

	return 0
}

//func getNewDrmWord(old string, token string) string {
//
//}
