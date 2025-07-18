package basic

import (
	"bufio"
	"fmt"
	"os"
)

//子串 连续的
//子序列 不连续

func MaxIncreaseSubSeq() {
	const (
		N = 1e5 + 10
	)

	var (
		n int
		a = make([]int, N)
		f = make([]int, N)
		g = make([]int, N)
	)

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	for i := 1; i <= n; i++ {
		j := n - i + 1
		if a[i] > a[i-1] {
			f[i] = f[i-1] + 1
		} else {
			f[i] = 1
		}

		if a[j] < a[j+1] {
			g[j] = g[j+1] + 1
		} else {
			g[j] = 1
		}
	}
	res := 0
	for i := 1; i <= n; i++ {
		tmp := 0
		if a[i-1] < a[i+1] {
			tmp = f[i-1] + g[i+1]
		} else {
			tmp = max(f[i-1], g[i+1])
		}
		res = max(res, tmp)
	}
	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
