import sys

def solve():
    # Читаем строку из стандартного ввода
    s = sys.stdin.readline().strip()
    if not s:
        return
        
    n = len(s)
    tbank = "tbank"
    study = "study"
    m = 5 # Длина обеих целевых строк одинакова

    # Если строка слишком короткая (хотя по условию n >= 10)
    if n < 10:
        print(-1)
        return

    # Предварительно вычисляем стоимость замены для каждой позиции
    # cost_t[i] - сколько замен нужно, чтобы s[i:i+5] стало "tbank"
    cost_t = [0] * (n - m + 1)
    cost_s = [0] * (n - m + 1)

    for i in range(n - m + 1):
        c_t = 0
        c_s = 0
        sub = s[i:i+m]
        for j in range(m):
            if sub[j] != tbank[j]:
                c_t += 1
            if sub[j] != study[j]:
                c_s += 1
        cost_t[i] = c_t
        cost_s[i] = c_s

    # Создаем массивы суффиксного минимума
    # suf_min_s[i] — минимальная стоимость сделать "study" на отрезке [i, n]
    suf_min_t = [0] * (n - m + 1)
    suf_min_s = [0] * (n - m + 1)

    suf_min_t[n - m] = cost_t[n - m]
    suf_min_s[n - m] = cost_s[n - m]

    for i in range(n - m - 1, -1, -1):
        suf_min_t[i] = min(cost_t[i], suf_min_t[i + 1])
        suf_min_s[i] = min(cost_s[i], suf_min_s[i + 1])

    # Ищем минимальную сумму двух непересекающихся подстрок
    ans = float('inf')

    # Проходим по всем возможным позициям первой строки (i)
    # Вторая строка должна начинаться как минимум в позиции i + m
    for i in range(n - 2 * m + 1):
        # Вариант 1: сначала tbank, потом study
        ans = min(ans, cost_t[i] + suf_min_s[i + m])
        # Вариант 2: сначала study, потом tbank
        ans = min(ans, cost_s[i] + suf_min_t[i + m])

    print(ans)

if __name__ == "__main__":
    solve()
