package main

import (
	"bufio"
	"fmt"
	"os"
)

const INF = int(1e9)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	in := bufio.NewReaderSize(os.Stdin, 1<<20)
	out := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	g := make([][]int, n)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		a--
		b--
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	ans := INF
	queue := make([]int, n)

	for start := 0; start < n; start++ {
		dist := make([]int, n)
		parent := make([]int, n)
		for i := 0; i < n; i++ {
			dist[i] = -1
			parent[i] = -1
		}

		head, tail := 0, 0
		queue[tail] = start
		tail++
		dist[start] = 0

		for head < tail {
			u := queue[head]
			head++

			if 2*dist[u]+1 >= ans {
				continue
			}

			for _, v := range g[u] {
				if dist[v] == -1 {
					dist[v] = dist[u] + 1
					parent[v] = u
					queue[tail] = v
					tail++
				} else if parent[u] != v && parent[v] != u {
					ans = min(ans, dist[u]+dist[v]+1)
				}
			}
		}
	}

	if ans == INF {
		fmt.Fprintln(out, -1)
	} else {
		fmt.Fprintln(out, ans)
	}
}
