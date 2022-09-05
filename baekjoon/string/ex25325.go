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

	text, _, _ := read.ReadLine()
	var result = getPopularStudentMap(strings.Split(string(text), " "))

	for i := 0; i < n; i++ {
		text, _, _ = read.ReadLine()
		result = addPopularPointToStudents(strings.Split(string(text), " "), result)
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

func addPopularPointToStudents(students []string, result map[string]int) map[string]int {

	if len(students) <= 1 {
		result[students[0]]++
		return result
	}

	result[students[0]]++
	result[students[1]]++
	return result
}
