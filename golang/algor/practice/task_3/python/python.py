def solve():
    t = int(input())
    for _ in range(t):
        s = input().strip()
        n = len(s)
        
        ones = [i for i, ch in enumerate(s) if ch == '1']
        if not ones:
            print(0)
            continue
        
        # ищем максимальный непрерывный отрезок единиц с учетом цикличности
        # для этого дублируем список + n
        ext = ones + [x + n for x in ones]
        max_len = 1
        cur_len = 1
        for i in range(1, len(ext)):
            if ext[i] == ext[i-1] + 1:
                cur_len += 1
            else:
                max_len = max(max_len, cur_len)
                cur_len = 1
        max_len = max(max_len, cur_len)
        
        # максимальная длина в циклическом смысле не может превышать n
        max_len = min(max_len, n)
        
        # теперь по длине L находим площадь
        L = max_len
        # h + w = L + 1, макс h*w при h,w ~ (L+1)/2
        h = (L + 1) // 2
        w = (L + 2) // 2
        ans = h * w
        print(ans)

if name == "main":
    solve()