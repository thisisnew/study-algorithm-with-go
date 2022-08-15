package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var mbtis = []string{"ISTJ", "ISTP", "ISFJ", "ISFP", "INTJ", "INTP", "INFJ", "INFP", "ESTJ", "ESTP", "ESFJ", "ESFP", "ENTJ", "ENTP", "ENFJ", "ENFP"}

func main() {

	var n, m int
	var read = bufio.NewReader(os.Stdin)

	fmt.Fscanln(read, &n, &m)

	var sl = make([][]string, n)

	for i := 0; i < len(sl); i++ {
		var props = make([]string, m)
		var input string
		fmt.Fscanln(read, &input)

		for i, s := range input {
			props[i] = string(s)
		}

		sl[i] = props
	}

	var result int

	//가로
	if m > 3 {
		result += readWordHorizontally(sl)
	}

	//세로
	if n > 3 {
		result += readWordVertically(sl, m)
	}

	if n > 3 && m > 3 {
		//대각선

	}
}

func countMBTIs(word string) int {

	var result int

	for _, mbti := range mbtis {
		result += strings.Count(word, mbti)
	}

	return result
}

func readWordHorizontally(sl [][]string) int {

	var result int

	for _, props := range sl {
		result += countMBTIs(strings.Join(props, ""))
	}

	return result
}

func readWordVertically(sl [][]string, m int) int {

	var result int

	for i := 0; i < m; i++ {

		var word strings.Builder

		for _, props := range sl {
			word.WriteString(props[i])
		}

		result += countMBTIs(word.String())
	}

	return result
}
