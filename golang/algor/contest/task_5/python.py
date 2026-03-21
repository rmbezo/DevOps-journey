def solve():
    import sys
    import math
    
    input = sys.stdin.read
    data = input().split()
    
    n = int(data[0])
    a = list(map(int, data[1:]))
    
    positions = {}
    for i, val in enumerate(a):
        if val not in positions:
            positions[val] = []
        positions[val].append(i)
    
    possible = {}
    for val in positions:
        if len(positions[val]) == 0:
            possible[val] = False
            continue
            
        if len(positions[val]) == n:
            possible[val] = True
            continue
            
   
        possible[val] = True
    
    ans_for_val = {}
    
    for val in positions:
        if not possible.get(val, False):
            continue
            
        if len(positions[val]) == n:
            ans_for_val[val] = [0] * n
            continue
            
        pos_list = sorted(positions[val])
        
        extended = pos_list + [p + n for p in pos_list]
        
        res = [0] * n
        
        for i in range(len(pos_list)):
            left = pos_list[i]
            right = pos_list[(i + 1) % len(pos_list)]
            
            if right < left:
                right += n
            
            length = right - left - 1
            
            if length <= 0:
                continue
            
            for j in range(1, length + 1):
                idx = (left + j) % n
                dist_to_left = j
                dist_to_right = length + 1 - j
                
                steps_left = (dist_to_left + 1) // 2
                steps_right = (dist_to_right + 1) // 2
                
                steps = min(steps_left, steps_right)
                
                if res[idx] == 0 or steps < res[idx]:
                    res[idx] = steps
        
        ans_for_val[val] = res
    
    result = []
    for i in range(n):
        val = a[i]
        if val not in ans_for_val:
            result.append("-1")
        else:
            result.append(str(ans_for_val[val][i]))
    
    print(" ".join(result))

if __name__ == "__main__":
    solve()
