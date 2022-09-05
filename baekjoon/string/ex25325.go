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

	var result map[string]int

	for i := 0; i <= n; i++ {
		text, _, _ := read.ReadLine()
		students := strings.Split(string(text), " ")

		if i == 0 {
			result = getPopularStudentMap(students)
		} else {

			if len(students) <= 1 {
				result[students[0]]++
				continue
			}

			result[students[0]]++
			result[students[1]]++
		}
	}

	fmt.Println(result)
}

func getPopularStudentMap(students []string) map[string]int {

	var result = map[string]int{}

	for _, student := range students {
		result[student] = 0
	}

	return result
}
