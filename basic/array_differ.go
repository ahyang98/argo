package basic

import (
	"bufio"
	"fmt"
	"os"
)

func differ() {
	const (
		N = 100010
	)

	var (
		n int
		m int
		a [N]int
		b [N]int
		l int
		r int
		c int
	)

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n, &m)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	fmt.Fscanln(reader)

	for i := 1; i <= n; i++ {
		insert(&b, i, i, a[i])
	}

	for i := 0; i < m; i++ {
		fmt.Fscanln(reader, &l, &r, &c)
		insert(&b, l, r, c)
	}

	for i := 1; i <= n; i++ {
		b[i] = b[i-1] + b[i]
		fmt.Printf("%d ", b[i])
	}
}

func insert(b *[100010]int, l int, r int, c int) {
	b[l] = b[l] + c
	b[r+1] = b[r+1] - c
}
