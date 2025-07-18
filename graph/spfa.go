package graph

import (
	"bufio"
	"fmt"
	"os"
)

type Spfa struct {
	N   int
	h   []int
	e   []int
	w   []int
	ne  []int
	idx int
}

func NewSpfa() *Spfa {
	const N = 1e5 + 10
	return &Spfa{
		N:  N,
		h:  make([]int, N),
		e:  make([]int, N),
		w:  make([]int, N),
		ne: make([]int, N),
	}
}

func (s *Spfa) ShortestDist() {
	reader := bufio.NewReader(os.Stdin)
	var n, m, a, b, c int
	fmt.Fscanln(reader, &n, &m)
	for i := 0; i < s.N; i++ {
		s.h[i] = -1
	}
	for i := 0; i < m; i++ {
		fmt.Fscanln(reader, &a, &b, &c)
		s.add(a, b, c)
	}
	s.spfa(n, m)
}

func (s *Spfa) add(a, b, c int) {
	s.e[s.idx] = b
	s.w[s.idx] = c
	s.ne[s.idx] = s.h[a]
	s.h[a] = s.idx
	s.idx++
}

func (s *Spfa) spfa(n, m int) {
	const INF = 0x3f3f3f3f
	var (
		dist = make([]int, s.N)
		st   = make([]bool, s.N)
		q    = make([]int, s.N)
		hh   = 0
		tt   = -1
	)
	for i := 0; i < s.N; i++ {
		dist[i] = INF
	}
	dist[1] = 0
	st[1] = true
	tt++
	q[tt] = 1

	for hh <= tt {
		a := q[hh]
		hh++
		st[a] = false
		for i := s.h[a]; i != -1; i = s.ne[i] {
			b := s.e[i]
			if dist[b] > dist[a]+s.w[i] {
				dist[b] = dist[a] + s.w[i]
				if !st[b] {
					tt++
					q[tt] = b
					st[b] = true
				}
			}

		}
	}

	if dist[n] == INF {
		fmt.Println("impossible")
	} else {
		fmt.Println(dist[n])
	}
}
