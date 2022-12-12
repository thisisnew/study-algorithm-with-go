package main

import (
	"bufio"
	"fmt"
	"os"
)

//백준 1260

var (
	graph   [][]int
	visited []bool
	writer  *bufio.Writer
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m, v int
	fmt.Fscanln(reader, &n, &m, &v)

	graph = make([][]int, n+1)
	for i := range graph {
		graph[i] = make([]int, n+1)
	}
	visited = make([]bool, n+1)

	for i := 0; i < m; i++ {
		var vertex1, vertex2 int
		fmt.Fscanln(reader, &vertex1, &vertex2)
		graph[vertex1][vertex2] = 1
		graph[vertex2][vertex1] = 1
	}

	dfs(v)
}

func dfs(v int) {
	visited[v] = true // dfs에 들어오면 방문했다고 판단
	fmt.Fprintf(writer, "%d ", v)

	for i := 0; i < len(graph[v]); i++ {
		if graph[v][i] == 1 && !visited[i] { // 연결되어 있고 방문하지 않은 경우
			dfs(i) // 재귀함수 호출
		}
	}
}
