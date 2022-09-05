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

	text, _, _ := read.ReadLine()
	var result = getPopularStudentMap(strings.Split(string(text), " "))

	for i := 0; i < n; i++ {
		text, _, _ = read.ReadLine()
		result = addPopularPointToStudents(strings.Split(string(text), " "), result)
	}

	reversedResult, points := getPopularPointMap(result)
	printPopularStudents(reversedResult, points)
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

func getPopularPointMap(students map[string]int) (map[int][]string, []int) {

	var result = map[int][]string{}
	var points = make([]int, len(students))

	var idx = 0
	for k, v := range students {

		points[idx] = v
		idx++

		student, ok := result[v]

		if !ok {
			result[v] = []string{k}
			continue
		}

		if k > student[len(student)-1] {
			student = append(student, k)
		} else {
			student = append([]string{k}, student...)
		}

		result[v] = student
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i] > points[j]
	})

	return result, points
}

func printPopularStudents(students map[int][]string, points []int) {

	var prevPoint int
	for _, p := range points {
		if p == prevPoint {
			continue
		}

		student, _ := students[p]

		for _, s := range student {
			fmt.Println(s, p)
		}

		prevPoint = p
	}

}
