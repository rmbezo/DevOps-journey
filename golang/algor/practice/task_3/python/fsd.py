from collections import deque

INF = int(1e9)

def solve():
    data = sys.stdin.read().split()
    if not data:
        return
    
    idx = 0
    n = int(data[idx]); idx += 1
    m = int(data[idx]); idx += 1
    
    # Создаем список смежности
    g = [[] for _ in range(n)]
    
    for _ in range(m):
        a = int(data[idx]) - 1; idx += 1
        b = int(data[idx]) - 1; idx += 1
        g[a].append(b)
        g[b].append(a)
    
    ans = INF
    
    for start in range(n):
        dist = [-1] * n
        parent = [-1] * n
        
        queue = deque()
        queue.append(start)
        dist[start] = 0
        
        while queue:
            u = queue.popleft()
            
            if 2 * dist[u] + 1 >= ans:
                continue
            
            for v in g[u]:
                if dist[v] == -1:
                    dist[v] = dist[u] + 1
                    parent[v] = u
                    queue.append(v)
                elif parent[u] != v and parent[v] != u:
                    ans = min(ans, dist[u] + dist[v] + 1)
    
    if ans == INF:
        print(-1)
    else:
        print(ans)

if __name__ == "__main__":
    solve()
