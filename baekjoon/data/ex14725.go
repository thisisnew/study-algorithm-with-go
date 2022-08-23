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
		inputs := strings.Split(string(text), " ")
		key := inputs[1]
		prev, ok := result[key]

		if !ok {
			result[key] = [][]string{inputs[2:]}
		} else {
			prev = append(prev, inputs[2:])
			result[key] = prev
		}
	}

	keys := getAntPropertiesKeys(result)
	printAntProperties(keys, result)
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

		sort.SliceStable(props, func(i, j int) bool {
			return props[i][0] < props[j][0]
		})

		for _, prop := range props {
			for i, p := range prop {
				fmt.Println(generateBars(i) + p)
			}
		}

	}
}
