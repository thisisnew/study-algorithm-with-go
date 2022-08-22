package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
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
		result[i] = strconv.Itoa(i) + prop
	}

	return result
}

func addAntPropertiesToPrevAntProperties(props []string, prev []string) []string {

	for i, prop := range props {
		prev = append(prev, strconv.Itoa(i)+prop)
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
		prevIdx := 0
		var printProps []string
		var bars string
		for _, prop := range props {
			idx, _ := strconv.Atoi(prop[0:1])

			if idx < prevIdx {
				printOrderedProps(bars, printProps)
				prevIdx = idx
				printProps = []string{}
				bars = ""
			}

			if len(bars) == 0 {
				bars = generateBars(idx)
			}

			printProps = append(printProps, prop[1:])
		}

		if len(printProps) > 0 {
			printOrderedProps(bars, printProps)
		}
	}
}

func printOrderedProps(bars string, printProps []string) {

	sort.Slice(printProps, func(i, j int) bool {
		return printProps[i] < printProps[j]
	})

	for _, ppr := range printProps {
		fmt.Println(bars + ppr)
	}
}
