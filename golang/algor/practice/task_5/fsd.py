import sys

def min3(a, b, c):
    return min(a, b, c)

def max3(a, b, c):
    return max(a, b, c)

def solve():
    first_line = sys.stdin.readline()
    if not first_line:
        return
    
    n = int(first_line.strip())
    
    second_line = sys.stdin.readline()
    while second_line and second_line.strip() == '':
        second_line = sys.stdin.readline()
    
    if not second_line:
        return
    
    a = list(map(int, second_line.split()))
    
    while len(a) < n:
        next_line = sys.stdin.readline()
        if not next_line:
            break
        a.extend(map(int, next_line.split()))
    
    MAXA = 100000
    
    freq = [0] * (MAXA + 1)
    for val in a:
        freq[val] += 1
    
    good = [False] * (MAXA + 1)
    
    for i in range(n):
        x = a[i]
        y = a[(i + 1) % n]
        z = a[(i + 2) % n]
        
        mn = min3(x, y, z)
        mx = max3(x, y, z)
        
        good[mn] = True
        good[mx] = True
    
    result = []
    for i in range(n):
        val = a[i]
        ans = n - freq[val]
        if not good[val]:
            ans += 1
        result.append(str(ans))
    
    print(" ".join(result))

if __name__ == "__main__":
    solve()
