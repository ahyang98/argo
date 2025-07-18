package graph

import (
	"bufio"
	"fmt"
	"os"
)

type SurroundTable struct {
	N        int
	n        int
	relation [][]bool
	res      int
	seat     []int
	st       []bool
}

func NewSurroundTable() *SurroundTable {
	const N = 11
	return &SurroundTable{
		N:        N,
		relation: make([][]bool, N),
		seat:     make([]int, N),
		st:       make([]bool, N),
	}
}

func (t *SurroundTable) Init() {
	for i := 0; i < t.N; i++ {
		t.relation[i] = make([]bool, t.N)
	}
	reader := bufio.NewReader(os.Stdin)
	var m, a, b int
	fmt.Fscanln(reader, &t.n, &m)

	for i := 0; i < m; i++ {
		fmt.Fscanln(reader, &a, &b)
		t.relation[a][b] = true
		t.relation[b][a] = true
	}
	t.seat[1] = 1
	t.st[1] = true
	t.dfs(2)
	fmt.Println(t.res)
}

func (t *SurroundTable) dfs(pos int) {
	if pos > t.n {
		if !t.relation[1][t.seat[t.n]] {
			t.res++
		}
		return
	}

	for i := 1; i <= t.n; i++ {
		if t.st[i] || t.relation[i][t.seat[pos-1]] {
			continue
		}
		t.seat[pos] = i
		t.st[i] = true
		t.dfs(pos + 1)
		t.st[i] = false
	}
}
