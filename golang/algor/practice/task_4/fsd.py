import sys

INF = int(1e9)

def solve():
    # Читаем первую строку
    first_line = sys.stdin.readline()
    if not first_line:
        return
    
    # Разбираем n и m
    parts = first_line.split()
    n = int(parts[0])
    m = int(parts[1])
    
    # Создаем список смежности
    g = [[] for _ in range(n)]
    
    # Читаем m ребер
    for _ in range(m):
        line = sys.stdin.readline()
        while line and line.strip() == '':
            line = sys.stdin.readline()
        if not line:
            break
        a, b = map(int, line.split())
        a -= 1
        b -= 1
        g[a].append(b)
        g[b].append(a)
    
    ans = INF
    
    for start in range(n):
        dist = [-1] * n
        parent = [-1] * n
        
        # Ручная реализация очереди через список
        queue = [0] * n
        head = 0
        tail = 0
        
        queue[tail] = start
        tail += 1
        dist[start] = 0
        
        while head < tail:
            u = queue[head]
            head += 1
            
            if 2 * dist[u] + 1 >= ans:
                continue
            
            for v in g[u]:
                if dist[v] == -1:
                    dist[v] = dist[u] + 1
                    parent[v] = u
                    queue[tail] = v
                    tail += 1
                elif parent[u] != v and parent[v] != u:
                    ans = min(ans, dist[u] + dist[v] + 1)
    
    if ans == INF:
        print(-1)
    else:
        print(ans)

if __name__ == "__main__":
    solve()
