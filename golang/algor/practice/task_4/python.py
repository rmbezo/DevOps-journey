from collections import deque
import sys

def solve():
    input = sys.stdin.read
    data = input().split()
    
    n, m = int(data[0]), int(data[1])
    edges = []
    graph = [[] for _ in range(n)]
    
    idx = 2
    for i in range(m):
        a = int(data[idx]) - 1
        b = int(data[idx + 1]) - 1
        idx += 2
        edges.append((a, b))
        graph[a].append((b, i))
        graph[b].append((a, i))
    
    INF = 10**9
    ans = INF
    
    # BFS, запрещая использовать конкретное ребро
    def bfs_without_edge(u, v, forbidden_edge):
        dist = [-1] * n
        dist[u] = 0
        q = deque([u])
        while q:
            cur = q.popleft()
            for nxt, ei in graph[cur]:
                if ei == forbidden_edge:
                    continue
                if dist[nxt] == -1:
                    dist[nxt] = dist[cur] + 1
                    q.append(nxt)
        return dist[v]
    
    for i, (a, b) in enumerate(edges):
        d = bfs_without_edge(a, b, i)
        if d != -1:
            ans = min(ans, d + 1)
    
    print(ans if ans != INF else -1)

if __name__ == "__main__":
    solve()
