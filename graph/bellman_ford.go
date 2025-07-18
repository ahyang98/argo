package graph

import (
	"bufio"
	"fmt"
	"os"
)

type Edge struct {
	a int
	b int
	w int //weight
}

type BellManFord struct {
	n     int
	m     int
	k     int
	edges []Edge
}

func NewBellManFord() *BellManFord {
	return &BellManFord{}
}

func (f *BellManFord) MinDist() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &f.n, &f.m, &f.k)
	f.edges = make([]Edge, f.m)
	for i := 0; i < f.m; i++ {
		fmt.Fscanln(reader, &f.edges[i].a, &f.edges[i].b, &f.edges[i].w)
	}

	f.bellmanFord()
}

func (f *BellManFord) bellmanFord() {
	const (
		N   = 510
		INF = 0x3f3f3f3f
	)

	var (
		dist = make([]int, N)
		last = make([]int, N)
	)

	for i := 0; i < N; i++ {
		dist[i] = INF
	}

	dist[1] = 0

	for i := 0; i < f.k; i++ {
		for j := 0; j < f.m; j++ {
			last[j] = dist[j]
		}
		for j := 0; j < f.m; j++ {
			a := f.edges[j].a
			b := f.edges[j].b
			w := f.edges[j].w
			if dist[b] > last[a]+w {
				dist[b] = last[a] + w
			}
		}
	}

	if dist[f.n] > INF/2 {
		fmt.Println("impossible")
	} else {
		fmt.Println(dist[f.n])
	}
}
