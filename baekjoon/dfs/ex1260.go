package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	graph   [][]bool
	visited []bool
	writer  *bufio.Writer
)

func main() {

	var n, m, v int
	var read = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(read, &n, &m, &v)
	graph = make([][]bool, n+1)
	visited = make([]bool, n+1)

	for i, _ := range graph {
		graph[i] = make([]bool, n+1)
	}

	for i := 0; i < m; i++ {
		var p1, p2 int
		fmt.Fscanln(read, &p1, &p2)
		graph[p1][p2] = true
		graph[p2][p1] = true
	}

	dfs(v)
}

func dfs(v int) {
	visited[v] = true // dfs에 들어오면 방문했다고 판단
	fmt.Fprintf(writer, "%d ", v)

	for i := 0; i < len(graph[v]); i++ {
		if graph[v][i] && !visited[i] { // 연결되어 있고 방문하지 않은 경우
			dfs(i) // 재귀함수 호출
		}
	}
}
