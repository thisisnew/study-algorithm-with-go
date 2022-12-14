package main

import "fmt"

func main() {
	fmt.Println(네트워크(3, [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}))
}

var (
	isVisited []bool
	result    int
)

func 네트워크(n int, computers [][]int) int {
	isVisited = make([]bool, n+1)

	for i := 0; i < n; i++ {
		if !isVisited[i] {
			result++
			dfs(i, computers)
		}
	}

	return result
}

func dfs(n int, computers [][]int) {
	isVisited[n] = true

	for i := 0; i < len(computers[n]); i++ {
		if !isVisited[i] && computers[n][i] == 1 {
			dfs(i, computers)
		}
	}
}
