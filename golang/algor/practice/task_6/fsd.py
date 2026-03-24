import sys

class FastScanner:
    def __init__(self):
        self.data = sys.stdin.buffer.read()
        self.pos = 0
        self.n = len(self.data)
    
    def skip_spaces(self):
        while self.pos < self.n:
            c = self.data[self.pos]
            if c > 32:  # > ' '
                break
            self.pos += 1
    
    def next_int(self):
        self.skip_spaces()
        sign = 1
        if self.pos < self.n and self.data[self.pos] == ord('-'):
            sign = -1
            self.pos += 1
        val = 0
        while self.pos < self.n:
            c = self.data[self.pos]
            if c < ord('0') or c > ord('9'):
                break
            val = val * 10 + (c - ord('0'))
            self.pos += 1
        return val * sign
    
    def next_string(self):
        self.skip_spaces()
        start = self.pos
        while self.pos < self.n and self.data[self.pos] > 32:
            self.pos += 1
        return self.data[start:self.pos].decode()


class SegTree:
    def __init__(self, n):
        self.n = n
        self.sum = [0] * (4 * n + 5)
        self.lazy = [0] * (4 * n + 5)
        self._build(1, 1, n)
    
    def _build(self, v, l, r):
        self.lazy[v] = 1
        if l == r:
            self.sum[v] = 1
            return
        m = (l + r) >> 1
        self._build(v << 1, l, m)
        self._build(v << 1 | 1, m + 1, r)
        self.sum[v] = self.sum[v << 1] + self.sum[v << 1 | 1]
    
    def _apply(self, v, mul):
        self.sum[v] *= mul
        self.lazy[v] *= mul
    
    def _push(self, v):
        if self.lazy[v] != 1:
            mul = self.lazy[v]
            self._apply(v << 1, mul)
            self._apply(v << 1 | 1, mul)
            self.lazy[v] = 1
    
    def range_mul(self, ql, qr, mul):
        self._range_mul(1, 1, self.n, ql, qr, mul)
    
    def _range_mul(self, v, l, r, ql, qr, mul):
        if ql > r or qr < l:
            return
        if ql <= l and r <= qr:
            self._apply(v, mul)
            return
        self._push(v)
        m = (l + r) >> 1
        self._range_mul(v << 1, l, m, ql, qr, mul)
        self._range_mul(v << 1 | 1, m + 1, r, ql, qr, mul)
        self.sum[v] = self.sum[v << 1] + self.sum[v << 1 | 1]
    
    def point_add(self, pos, delta):
        self._point_add(1, 1, self.n, pos, delta)
    
    def _point_add(self, v, l, r, pos, delta):
        if l == r:
            self.sum[v] += delta
            return
        self._push(v)
        m = (l + r) >> 1
        if pos <= m:
            self._point_add(v << 1, l, m, pos, delta)
        else:
            self._point_add(v << 1 | 1, m + 1, r, pos, delta)
        self.sum[v] = self.sum[v << 1] + self.sum[v << 1 | 1]
    
    def query_point(self, pos):
        return self._query_point(1, 1, self.n, pos)
    
    def _query_point(self, v, l, r, pos):
        if l == r:
            return self.sum[v]
        self._push(v)
        m = (l + r) >> 1
        if pos <= m:
            return self._query_point(v << 1, l, m, pos)
        return self._query_point(v << 1 | 1, m + 1, r, pos)
    
    def query_sum(self, ql, qr):
        return self._query_sum(1, 1, self.n, ql, qr)
    
    def _query_sum(self, v, l, r, ql, qr):
        if ql > r or qr < l:
            return 0
        if ql <= l and r <= qr:
            return self.sum[v]
        self._push(v)
        m = (l + r) >> 1
        left = self._query_sum(v << 1, l, m, ql, qr)
        right = self._query_sum(v << 1 | 1, m + 1, r, ql, qr)
        return left + right
    
    def prefix_sum(self, pos):
        if pos <= 0:
            return 0
        return self.query_sum(1, pos)
    
    def kth(self, k):
        return self._kth(1, 1, self.n, k)
    
    def _kth(self, v, l, r, k):
        if l == r:
            return l
        self._push(v)
        m = (l + r) >> 1
        if self.sum[v << 1] >= k:
            return self._kth(v << 1, l, m, k)
        return self._kth(v << 1 | 1, m + 1, r, k - self.sum[v << 1])


def solve():
    sc = FastScanner()
    n = sc.next_int()
    q = sc.next_int()
    s = sc.next_string()
    
    st = SegTree(n)
    out_lines = []
    
    for _ in range(q):
        tp = sc.next_int()
        
        if tp == 1:
            l = sc.next_int()
            r = sc.next_int()
            
            left_idx = st.kth(l)
            right_idx = st.kth(r)
            
            if left_idx == right_idx:
                st.point_add(left_idx, r - l + 1)
                continue
            
            left_cnt = st.query_point(left_idx)
            left_pref = st.prefix_sum(left_idx - 1)
            left_add = (left_pref + left_cnt) - l + 1
            
            right_pref = st.prefix_sum(right_idx - 1)
            right_add = r - right_pref
            
            if left_idx + 1 <= right_idx - 1:
                st.range_mul(left_idx + 1, right_idx - 1, 2)
            
            st.point_add(left_idx, left_add)
            st.point_add(right_idx, right_add)
        
        else:  # tp == 2
            i = sc.next_int()
            idx = st.kth(i)
            out_lines.append(s[idx - 1])
    
    sys.stdout.write("\n".join(out_lines))


if __name__ == "__main__":
    solve()