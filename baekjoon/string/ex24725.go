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
		result += readMBTIWordHorizontally(sl)
	}

	//세로
	if n > 3 {
		result += readMBTIWordVertically(sl, m)
	}

	//대각선
	if n > 3 && m > 3 {
		result += readMBTIWordDiagonally(sl, n, m)
	}

	fmt.Println(result)
}

func countMBTIs(word string) int {

	var result int

	for _, m := range mbtis {
		result += strings.Count(word, m)
		result += strings.Count(reverseMBTIWord(word), m)
	}

	return result
}

func readMBTIWordHorizontally(sl [][]string) int {

	var result int

	for _, props := range sl {
		result += countMBTIs(strings.Join(props, ""))
	}

	return result
}

func readMBTIWordVertically(sl [][]string, m int) int {

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

func readMBTIWordDiagonally(sl [][]string, n, m int) int {

	var result int
	var word string
	var hIdx int
	var vIdx int

	for {

		if hIdx == n {
			hIdx = 0
			result += countMBTIs(word)
		}

		if vIdx == m {
			vIdx = 0
			result += countMBTIs(word)
		}

		word += sl[hIdx][vIdx]
		hIdx++
		vIdx++

		break
	}

	//for _, props := range sl {
	//
	//	if idx == len(props) {
	//		fmt.Println("diagonal : ", word.String())
	//		result += countMBTIs(word.String())
	//		word.Reset()
	//		continue
	//	}
	//
	//	word.WriteString(props[idx])
	//	idx++
	//}

	return result
}

func reverseMBTIWord(s string) string {

	var result string

	for _, v := range s {
		result = string(v) + result
	}

	return result
}
