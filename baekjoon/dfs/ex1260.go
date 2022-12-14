package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	graph     [][]bool
	isVisited []bool
	writer    *bufio.Writer
)

func main() {
	var n, m, v int
	var read = bufio.NewReader(os.Stdin)
	fmt.Fscanln(read, &n, &m, &v)

	writer = bufio.NewWriter(os.Stdout)

	graph = make([][]bool, n+1)
	for i := range graph {
		graph[i] = make([]bool, n+1)
	}

	isVisited = make([]bool, n+1)

	for i := 0; i < m; i++ {
		var p1, p2 int
		fmt.Fscanln(read, &p1, &p2)

		graph[p1][p2] = true
		graph[p2][p1] = true
	}

	dfs(v)
}

func dfs(v int) {
	isVisited[v] = true
	fmt.Fprintf(writer, "%d", v)

	for i := 0; i < len(graph[v]); i++ {
		if graph[v][i] && !isVisited[i] {
			dfs(i)
		}
	}
}
