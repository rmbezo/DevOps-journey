s = input().strip()
n = len(s)

word1 = "tbank"
word2 = "study"
lw = 5

def match_count(pos, word):
    # сколько символов совпадает, если word вставить в позицию pos
    cnt = 0
    for i in range(lw):
        if pos + i < n and s[pos + i] == word[i]:
            cnt += 1
    return cnt

# перебираем позиции для tbank и study
best = 0
for p1 in range(n - lw + 1):
    for p2 in range(n - lw + 1):
        # проверим, не конфликтуют ли они
        ok = True
        # строим массив символов после замен
        temp = [''] * n
        # заполняем tbank
        for i in range(lw):
            temp[p1 + i] = word1[i]
        # заполняем study
        for i in range(lw):
            pos = p2 + i
            if temp[pos] and temp[pos] != word2[i]:
                ok = False
                break
            temp[pos] = word2[i]
        if not ok:
            continue
        # считаем совпадения
        matches = 0
        for i in range(n):
            if temp[i] == s[i]:
                matches += 1
        best = max(best, matches)

# минимальные замены = n - best
print(n - best)
