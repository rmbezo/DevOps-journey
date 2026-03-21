package main

import (
	"bufio"
	"io"
	"os"
)

type FastScanner struct {
	data []byte
	pos  int
	n    int
}

func NewFastScanner() *FastScanner {
	data, _ := io.ReadAll(os.Stdin)
	return &FastScanner{
		data: data,
		n:    len(data),
	}
}

func (fs *FastScanner) skipSpaces() {
	for fs.pos < fs.n {
		c := fs.data[fs.pos]
		if c > ' ' {
			break
		}
		fs.pos++
	}
}

func (fs *FastScanner) NextInt64() int64 {
	fs.skipSpaces()
	var sign int64 = 1
	if fs.pos < fs.n && fs.data[fs.pos] == '-' {
		sign = -1
		fs.pos++
	}
	var val int64
	for fs.pos < fs.n {
		c := fs.data[fs.pos]
		if c < '0' || c > '9' {
			break
		}
		val = val*10 + int64(c-'0')
		fs.pos++
	}
	return val * sign
}

func (fs *FastScanner) NextString() string {
	fs.skipSpaces()
	start := fs.pos
	for fs.pos < fs.n && fs.data[fs.pos] > ' ' {
		fs.pos++
	}
	return string(fs.data[start:fs.pos])
}

type SegTree struct {
	sum  []int64
	lazy []int64
	n    int
}

func NewSegTree(n int) *SegTree {
	st := &SegTree{
		sum:  make([]int64, 4*n+5),
		lazy: make([]int64, 4*n+5),
		n:    n,
	}
	st.build(1, 1, n)
	return st
}

func (st *SegTree) build(v, l, r int) {
	st.lazy[v] = 1
	if l == r {
		st.sum[v] = 1
		return
	}
	m := (l + r) >> 1
	st.build(v<<1, l, m)
	st.build(v<<1|1, m+1, r)
	st.sum[v] = st.sum[v<<1] + st.sum[v<<1|1]
}

func (st *SegTree) apply(v int, mul int64) {
	st.sum[v] *= mul
	st.lazy[v] *= mul
}

func (st *SegTree) push(v int) {
	if st.lazy[v] != 1 {
		mul := st.lazy[v]
		st.apply(v<<1, mul)
		st.apply(v<<1|1, mul)
		st.lazy[v] = 1
	}
}

func (st *SegTree) rangeMul(v, l, r, ql, qr int, mul int64) {
	if ql > r || qr < l {
		return
	}
	if ql <= l && r <= qr {
		st.apply(v, mul)
		return
	}
	st.push(v)
	m := (l + r) >> 1
	st.rangeMul(v<<1, l, m, ql, qr, mul)
	st.rangeMul(v<<1|1, m+1, r, ql, qr, mul)
	st.sum[v] = st.sum[v<<1] + st.sum[v<<1|1]
}

func (st *SegTree) pointAdd(v, l, r, pos int, delta int64) {
	if l == r {
		st.sum[v] += delta
		return
	}
	st.push(v)
	m := (l + r) >> 1
	if pos <= m {
		st.pointAdd(v<<1, l, m, pos, delta)
	} else {
		st.pointAdd(v<<1|1, m+1, r, pos, delta)
	}
	st.sum[v] = st.sum[v<<1] + st.sum[v<<1|1]
}

func (st *SegTree) queryPoint(v, l, r, pos int) int64 {
	if l == r {
		return st.sum[v]
	}
	st.push(v)
	m := (l + r) >> 1
	if pos <= m {
		return st.queryPoint(v<<1, l, m, pos)
	}
	return st.queryPoint(v<<1|1, m+1, r, pos)
}

func (st *SegTree) querySum(v, l, r, ql, qr int) int64 {
	if ql > r || qr < l {
		return 0
	}
	if ql <= l && r <= qr {
		return st.sum[v]
	}
	st.push(v)
	m := (l + r) >> 1
	left := st.querySum(v<<1, l, m, ql, qr)
	right := st.querySum(v<<1|1, m+1, r, ql, qr)
	return left + right
}

func (st *SegTree) prefixSum(pos int) int64 {
	if pos <= 0 {
		return 0
	}
	return st.querySum(1, 1, st.n, 1, pos)
}

func (st *SegTree) kth(v, l, r int, k int64) int {
	if l == r {
		return l
	}
	st.push(v)
	m := (l + r) >> 1
	if st.sum[v<<1] >= k {
		return st.kth(v<<1, l, m, k)
	}
	return st.kth(v<<1|1, m+1, r, k-st.sum[v<<1])
}

func main() {
	in := NewFastScanner()
	out := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer out.Flush()

	n := int(in.NextInt64())
	q := int(in.NextInt64())
	s := in.NextString()

	st := NewSegTree(n)

	for ; q > 0; q-- {
		tp := int(in.NextInt64())

		if tp == 1 {
			l := in.NextInt64()
			r := in.NextInt64()

			leftIdx := st.kth(1, 1, n, l)
			rightIdx := st.kth(1, 1, n, r)

			if leftIdx == rightIdx {
				st.pointAdd(1, 1, n, leftIdx, r-l+1)
				continue
			}

			leftCnt := st.queryPoint(1, 1, n, leftIdx)
			leftPref := st.prefixSum(leftIdx - 1)
			leftAdd := (leftPref + leftCnt) - l + 1

			rightPref := st.prefixSum(rightIdx - 1)
			rightAdd := r - rightPref

			if leftIdx+1 <= rightIdx-1 {
				st.rangeMul(1, 1, n, leftIdx+1, rightIdx-1, 2)
			}

			st.pointAdd(1, 1, n, leftIdx, leftAdd)
			st.pointAdd(1, 1, n, rightIdx, rightAdd)

		} else {
			i := in.NextInt64()
			idx := st.kth(1, 1, n, i)
			out.WriteByte(s[idx-1])
			out.WriteByte('\n')
		}
	}
}
