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

	result := map[string][]string{}

	for i := 0; i < n; i++ {
		text, _, _ := read.ReadLine()
		sl := strings.Split(string(text), " ")
		key := sl[1]
		pr, ok := result[key]

		if !ok {
			result[key] = getNewAntProperties(getAntProperties(sl))
		} else {
			result[key] = addAntPropertiesToPrevAntProperties(getAntProperties(sl), pr)
		}
	}

	keys := getAntPropertiesKeys(result)
	printAntProperties(keys, result)
}

func getAntProperties(sl []string) []string {

	var result = make([]string, len(sl)-2)

	for i := 0; i < len(result); i++ {
		result[i] = sl[i+2]
	}

	return result
}

func getNewAntProperties(props []string) []string {

	var result = make([]string, len(props))

	for i, prop := range props {
		result[i] = generateBars(i) + prop
	}

	return result
}

func addAntPropertiesToPrevAntProperties(props []string, prev []string) []string {

	for i, prop := range props {
		prev = append(prev, generateBars(i)+prop)
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

func getAntPropertiesKeys(result map[string][]string) []string {
	var keys []string

	for key, _ := range result {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	return keys
}

func printAntProperties(keys []string, result map[string][]string) {

	for _, key := range keys {
		fmt.Println(key)
		props := result[key]

		for _, prop := range props {
			fmt.Println(prop)
		}
	}
}
