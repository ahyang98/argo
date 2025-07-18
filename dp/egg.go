package dp

import (
	"bufio"
	"fmt"
	"os"
)

func TestEgg() {
	const (
		N = 101
		M = 11
	)
	var f [N][M]int
	reader := bufio.NewReader(os.Stdin)
	for {
		var n, m int
		count, err := fmt.Fscanln(reader, &n, &m)
		if count < 2 || err != nil {
			break
		}
		for i := 1; i <= n; i++ {
			f[i][1] = i
		}
		for i := 1; i <= m; i++ {
			f[1][i] = 1
		}
		for i := 2; i <= n; i++ {
			for j := 2; j <= m; j++ {
				f[i][j] = f[i][j-1]
				for k := 1; k <= i; k++ {
					f[i][j] = min(f[i][j], max(f[k-1][j-1], f[i-k][j])+1)
				}
			}
		}
		fmt.Println(f[n][m])
	}
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func TestEgg2() {
	const (
		N = 101
		M = 11
	)
	var f [N][M]int
	reader := bufio.NewReader(os.Stdin)
	for {
		var n, m int
		count, err := fmt.Fscanln(reader, &n, &m)
		if count < 2 || err != nil {
			break
		}
		for i := 1; i <= n; i++ {
			for j := 1; j <= m; j++ {
				f[i][j] = 1 + f[i-1][j-1] + f[i-1][j]
			}
			if f[i][m] >= n {
				fmt.Println(f[i][m])
				break
			}
		}
	}
}
