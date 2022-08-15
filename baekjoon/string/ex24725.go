package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

	//fmt.Println(sl)

	var result int
	result += readHorizontalWord(sl)

}

func countMBTIs(word string) int {

	var result int

	for _, mbti := range []string{"ISTJ", "ISTP", "ISFJ", "ISFP", "INTJ", "INTP", "INFJ", "INFP", "ESTJ", "ESTP", "ESFJ", "ESFP", "ENTJ", "ENTP", "ENFJ", "ENFP"} {
		result += strings.Count(word, mbti)
	}

	return result
}

func readHorizontalWord(sl [][]string) int {

	var result int

	for _, props := range sl {
		result += countMBTIs(strings.Join(props, ""))
	}

	return result
}
