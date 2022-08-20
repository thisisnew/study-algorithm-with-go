package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var read = bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanln(read, &n)

	result := map[string][]string{}

	for i := 0; i < n; i++ {
		text, _, _ := read.ReadLine()
		sl := strings.Split(string(text), " ")
		key := sl[1]
		pr, ok := result[key]

		if !ok {
			result[key] = getNewAntProperties(sl)
		} else {
			result[key] = addAntPropertiesToPrevAntProperties(sl, pr)
		}
	}

	fmt.Println(result)
}

func getNewAntProperties(sl []string) []string {
	var result = make([]string, len(sl)-2)

	for j := 0; j < len(sl)-2; j++ {
		result[j] = fmt.Sprintf("%s%s", generateBars(j), sl[j+2])
	}

	return result
}

func addAntPropertiesToPrevAntProperties(sl, prev []string) []string {
	for j := 0; j < len(sl)-2; j++ {
		bars := generateBars(j)

		if len(prev) >= j+1 {
			prev[j] = fmt.Sprintf("%s%s%s", prev[j], bars, sl[j+2])
		} else {
			prev = append(prev, fmt.Sprintf("%s%s", bars, sl[j+2]))
		}
	}

	return prev
}

func generateBars(j int) string {

	var result strings.Builder

	for i := 0; i < 2*(j+1); i++ {
		result.WriteString("-")
	}

	return result.String()
}
