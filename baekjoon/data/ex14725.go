package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {

	var read = bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanln(read, &n)

	result := map[string][][]string{}

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

	keys := getAntPropertiesKeys(result)
	printAntProperties(keys, result)
}

func getNewAntProperties(sl []string) [][]string {
	var result = make([][]string, len(sl)-2)

	for j := 0; j < len(sl)-2; j++ {
		var sj []string
		sj = append(sj, fmt.Sprintf("%s%s", generateBars(j), sl[j+2]))
		result[j] = sj
	}

	return result
}

func addAntPropertiesToPrevAntProperties(sl []string, prev [][]string) [][]string {
	for i := 0; i < len(sl)-2; i++ {
		bars := generateBars(i)

		if len(prev) >= i+1 {
			prev[i] = append(prev[i], fmt.Sprintf("%s%s", bars, sl[i+2]))
		} else {
			prev = append(prev, []string{fmt.Sprintf("%s%s", bars, sl[i+2])})
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

func getAntPropertiesKeys(result map[string][][]string) []string {
	var keys []string

	for key, _ := range result {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	return keys
}

func printAntProperties(keys []string, result map[string][][]string) {

	for _, key := range keys {
		fmt.Println(key)
		props := result[key]

		//sort.Slice(props, func(i, j int) bool {
		//
		//	pi := strings.ReplaceAll(props[i], "--", "")
		//	pj := strings.ReplaceAll(props[j], "--", "")
		//
		//	return pi < pj
		//})

		fmt.Println(props)

		//for _, prop := range props {
		//	fmt.Println(prop)
		//}
	}

}
