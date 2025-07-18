package graph

import (
	"argo/basic"
	"bufio"
	"fmt"
	"os"
)

const INF = 0x3f3f3f3f

type MinesDistance struct {
	N int
	g [][]int
}

func NewMinesDistance() *MinesDistance {
	const N = 510
	return &MinesDistance{
		N: N,
		g: make([][]int, N),
	}
}

func (d *MinesDistance) Calc() {
	reader := bufio.NewReader(os.Stdin)
	var (
		n, m    int
		x, y, z int
	)
	fmt.Fscanln(reader, &n, &m)
	for i := 0; i < d.N; i++ {
		d.g[i] = make([]int, d.N)
		for j := 0; j < d.N; j++ {
			d.g[i][j] = INF
		}
	}
	for i := 0; i < m; i++ {
		fmt.Fscanln(reader, &x, &y, &z)
		d.g[x][y] = min(d.g[x][y], z)
	}
	fmt.Println(d.dijkstra(n))
}

// n 点个数
func (d *MinesDistance) dijkstra(n int) int {
	var (
		dist = make([]int, d.N)
		st   = make([]bool, d.N)
	)

	for i := 0; i < d.N; i++ {
		dist[i] = INF
	}

	dist[1] = 0

	for i := 0; i < n-1; i++ {
		t := -1
		for j := 1; j <= n; j++ {
			// 在还未确定最短路的点中，寻找距离最小的点
			// 从所有已经计算过最短距离的点出发，找出距离最近的点，这个点也就是确定了
			// 就可以用这个点去更新所有其他点的距离，能更新最多是这个点可达的邻接点
			// 体现的还是宽搜的思想
			if !st[j] && (t == -1 || dist[t] > dist[j]) {
				t = j
			}
		}

		// 用t更新其他点的距离
		for j := 1; j <= n; j++ {
			dist[j] = min(dist[j], dist[t]+d.g[t][j])
		}

		st[t] = true
	}

	if dist[n] == INF {
		return -1
	}
	return dist[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type MinesDistance2 struct {
	N   int
	h   []int
	e   []int
	w   []int
	ne  []int
	idx int
}

func NewMinesDistance2() *MinesDistance2 {
	const N = 1e6 + 10
	return &MinesDistance2{
		N:  N,
		h:  make([]int, N),
		e:  make([]int, N),
		w:  make([]int, N),
		ne: make([]int, N),
	}
}

func (d *MinesDistance2) Calc() {
	reader := bufio.NewReader(os.Stdin)
	var (
		n, m    int
		a, b, c int
	)
	fmt.Fscanln(reader, &n, &m)
	for i := 0; i < d.N; i++ {
		d.h[i] = -1
	}
	for i := 0; i < m; i++ {
		fmt.Fscanln(reader, &a, &b, &c)
		d.add(a, b, c)
	}
	fmt.Println(d.dijkstra(n))
}

// 稀疏图使用邻接表来保存
func (d *MinesDistance2) add(a, b, c int) {
	// 有重边也不要紧，假设1->2有权重为2和3的边，再遍历到点1的时候2号点的距离会更新两次放入堆中
	// 这样堆中会有很多冗余的点，但是在弹出的时候还是会弹出最小值2+x（x为之前确定的最短路径），
	// 并标记st为true，所以下一次弹出3+x会continue不会向下执行。
	d.e[d.idx] = b
	d.ne[d.idx] = d.h[a]
	d.w[d.idx] = c
	d.h[a] = d.idx
	d.idx++
}

// n 点个数
func (d *MinesDistance2) dijkstra(n int) int {
	var (
		dist = make([]int, d.N)
		st   = make([]bool, d.N)
	)

	for i := 0; i < d.N; i++ {
		dist[i] = INF
	}

	minHeap := basic.NewMinHeap()
	minHeap.Push(&Item{
		id:  1,
		dis: 0,
	})
	dist[1] = 0

	for minHeap.Len() > 0 {
		item := minHeap.Pop().(*Item)
		if st[item.id] {
			continue
		}
		st[item.id] = true
		for i := d.h[item.id]; i != -1; i = d.ne[i] {
			j := d.e[i]
			if dist[j] > dist[item.id]+d.w[i] {
				dist[j] = dist[item.id] + d.w[i]
				minHeap.Push(&Item{
					id:  j,
					dis: dist[j],
				})
			}
		}
	}

	if dist[n] == INF {
		return -1
	}
	return dist[n]
}

type Item struct {
	id  int
	dis int
}

func (i *Item) Value() int {
	return i.dis
}
