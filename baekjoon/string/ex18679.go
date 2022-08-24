package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var dictionary map[string]string

func main() {

	var read = bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanln(read, &n)

	dictionary = getMinioneseDictionary(read, n)

	var t int
	fmt.Fscanln(read, &t)

	for i := 0; i < t; i++ {
		fmt.Println(convertHumanWordsToMinionese(read))
	}
}

func getMinioneseDictionary(read *bufio.Reader, n int) map[string]string {
	result := map[string]string{}

	for i := 0; i < n; i++ {
		text, _, _ := read.ReadLine()
		words := strings.Split(string(text), "=")

		result[strings.TrimSpace(words[0])] = strings.TrimSpace(words[1])
	}

	return result
}

func convertHumanWordsToMinionese(read *bufio.Reader) string {
	var k int
	fmt.Fscanln(read, &k)
	text, _, _ := read.ReadLine()
	humanWords := strings.Split(string(text), " ")

	var result = make([]string, k)

	for i := 0; i < k; i++ {

		key := humanWords[i]
		v, ok := dictionary[key]

		if !ok {
			continue
		}

		result[i] = v
	}

	return strings.Join(result, " ")
}
