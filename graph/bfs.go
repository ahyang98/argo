package graph

import (
	"bufio"
	"fmt"
	"os"
)

func MatrixMinDis() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	var n, m int
	fmt.Fscanln(reader, &n, &m)
	a := make([][]int, n)
	d := make([][]int, n)
	var str string
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &str)
		for j := 0; j < m; j++ {
			if len(a[i]) < 1 {
				a[i] = make([]int, m)
				d[i] = make([]int, m)
			}
			a[i][j] = int(str[j] - '0')
			d[i][j] = -1
		}
	}
	bfs(a, d)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Fprintf(writer, "%d ", d[i][j])
		}
		fmt.Fprintln(writer)
	}
	writer.Flush()
}

func bfs(a, d [][]int) {
	const N = 1010 * 1010
	var q [N][2]int
	hh := 0
	tt := -1
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			if a[i][j] == 1 {
				tt++
				q[tt][0], q[tt][1] = i, j
				d[i][j] = 0
			}
		}
	}
	dx := []int{-1, 0, 1, 0}
	dy := []int{0, 1, 0, -1}
	for hh <= tt {
		x, y := q[hh][0], q[hh][1]
		hh++
		for i := 0; i < 4; i++ {
			nx, ny := x+dx[i], y+dy[i]
			if nx < 0 || nx >= len(a) || ny < 0 || ny >= len(a[0]) || d[nx][ny] != -1 {
				continue
			}
			d[nx][ny] = d[x][y] + 1
			tt++
			q[tt][0], q[tt][1] = nx, ny
		}
	}
}

type NearStore struct {
	N   int
	INF int
	h   []int
	e   []int
	w   []int
	ne  []int
	dis []int
	idx int
}

func NewNearStore() *NearStore {
	const (
		N = 1e5 + 10
		M = 3e5 + 10
	)

	return &NearStore{
		N:   N,
		INF: 0x3f3f3f3f,
		h:   make([]int, N),
		e:   make([]int, M),
		w:   make([]int, M),
		ne:  make([]int, M),
		dis: make([]int, N),
	}
}

func (s *NearStore) init() {
	for i := 0; i < s.N; i++ {
		s.h[i] = -1
	}

	reader := bufio.NewReader(os.Stdin)
	var (
		n, m, k int
		a, b, d int
	)

	fmt.Fscanln(reader, &n, &m)

	for i := 0; i < m; i++ {
		fmt.Fscanln(reader, &a, &b, &d)
		s.add(a, b, d)
		s.add(b, a, d)
	}
	fmt.Fscanln(reader, &k)
	var x int
	for i := 0; i < k; i++ {
		fmt.Fscanln(reader, &x)
		s.add(0, x, 0)
	}

	s.spfa()

	fmt.Fscanln(reader, &k)
	for i := 0; i < k; i++ {
		fmt.Fscanln(reader, &x)
		fmt.Println(s.dis[x])
	}

}

func (s *NearStore) add(a, b, c int) {
	s.e[s.idx] = b
	s.w[s.idx] = c
	s.ne[s.idx] = s.h[a]
	s.h[a] = s.idx
	s.idx++
}

func (s *NearStore) spfa() {
	var (
		q  = make([]int, 0)
		st = make([]bool, s.N)
	)

	for i := 0; i < s.N; i++ {
		s.dis[i] = s.INF
	}

	q = append(q, 0)
	s.dis[0] = 0
	st[0] = true

	for len(q) > 0 {
		a := q[0]
		q = q[1:]
		st[a] = false
		for i := s.h[a]; i != -1; i = s.ne[i] {
			b := s.e[i]
			if s.dis[b] > s.dis[a]+s.w[i] {
				s.dis[b] = s.dis[a] + s.w[i]
				if !st[b] {
					q = append(q, b)
					st[b] = true
				}
			}
		}
	}
}
