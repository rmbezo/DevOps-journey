import sys

MOD = 1_000_000_007

def colorWays(lengths, maxK):
    lengths.sort()
    
    dp = [0] * (maxK + 1)
    dp[0] = 1
    
    for L in lengths:
        for j in range(maxK, 0, -1):
            available = L - (j - 1)
            if available > 0:
                dp[j] = (dp[j] + dp[j-1] * available) % MOD
    
    return dp

def solve():
    data = sys.stdin.read().split()
    if not data:
        return
    
    n = int(data[0])
    k = int(data[1])
    
    if n == 1:
        if k == 1:
            print(1)
        else:
            print(0)
        return
    
    maxBishops = 2 * n - 2
    if k > maxBishops:
        print(0)
        return
    
    color1 = []
    color2 = []
    
    for d in range(2 * n - 1):
        length = n - abs((n - 1) - d)
        if d % 2 == 0:
            color1.append(length)
        else:
            color2.append(length)
    
    limit = min(k, maxBishops)
    
    ways1 = colorWays(color1, limit)
    ways2 = colorWays(color2, limit)
    
    ans = 0
    for i in range(k + 1):
        if i < len(ways1) and k - i < len(ways2):
            ans = (ans + ways1[i] * ways2[k - i]) % MOD
    
    print(ans)
    
if __name__ == "__main__":
    solve()
