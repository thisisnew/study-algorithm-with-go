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

	bfs(v)
	fmt.Fprintln(writer, "")
}

func bfs(v int) {
	visited[v] = true
	queue := []int{v}

	for len(queue) != 0 {
		front := queue[0]
		fmt.Fprintf(writer, "%d ", front)
		queue = queue[1:]
		for i := 0; i < len(graph[v]); i++ {
			if graph[front][i] == 1 && !visited[i] { // 연결되어 있고 방문하지 않은 경우
				visited[i] = true        // 방문 표시
				queue = append(queue, i) // 큐에 삽입
			}
		}
	}
}
