package dp

import (
	"bufio"
	"fmt"
	"os"
)

type LineDp struct {
}

func NewLineDp() *LineDp {
	return &LineDp{}
}

func (d *LineDp) Dp() {
	d.MaxPubSeq()
}

func (d *LineDp) DigitDelta() {
	const (
		N   = 510
		INF = 1e9
	)

	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanln(reader, &n)
	a := make([][]int, N)
	for i := 1; i <= n; i++ {
		a[i] = make([]int, N)
		for j := 1; j <= i; j++ {
			fmt.Fscan(reader, &a[i][j])
		}
		fmt.Fscanln(reader)
	}
	f := make([][]int, N)
	for i := 0; i <= n; i++ {
		f[i] = make([]int, N)
		for j := 0; j <= i+1; j++ {
			f[i][j] = INF * -1
		}
	}
	f[1][1] = a[1][1]
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			f[i][j] = d.max(f[i-1][j-1], f[i-1][j]) + a[i][j]
		}
	}
	var res int = INF * -1
	for i := 1; i <= n; i++ {
		res = d.max(res, f[n][i])
	}
	fmt.Println(res)

}

func (d *LineDp) max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func (d *LineDp) MaxAscSubSeq() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanln(reader, &n)
	const N = 1010
	a := make([]int, N)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	f := make([]int, N)
	for i := 1; i <= n; i++ {
		f[i] = 1
		for j := 1; j < i; j++ {
			if a[j] < a[i] {
				f[i] = d.max(f[i], f[j]+1)
			}
		}
	}
	res := 0
	for i := 1; i <= n; i++ {
		res = d.max(res, f[i])
	}
	fmt.Println(res)
}

func (d *LineDp) MaxAscSubSeq2() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanln(reader, &n)
	const N = 1010
	a := make([]int, N)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	//不同长度i上升子序列结尾的最小值q[i]
	q := make([]int, N)
	lenTh := 0
	q[0] = -2e9
	for i := 0; i < n; i++ {
		l := 0
		r := lenTh
		for l < r {
			mid := (l + r + 1) / 2
			if q[mid] < a[i] {
				l = mid
			} else {
				r = mid - 1
			}
		}
		lenTh = d.max(lenTh, r+1)
		q[r+1] = a[i]
	}

	fmt.Println(lenTh)
}

func (d *LineDp) MaxPubSeq() {
	reader := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscanln(reader, &n, &m)
	a := make([]int32, n+1)
	b := make([]int32, m+1)
	for i := 1; i <= n; i++ {
		fmt.Fscanf(reader, "%c", &a[i])
	}
	fmt.Fscanln(reader)
	for i := 1; i <= m; i++ {
		fmt.Fscanf(reader, "%c", &b[i])
	}

	N := 1010
	f := make([][]int, N)
	f[0] = make([]int, N)
	for i := 1; i <= n; i++ {
		f[i] = make([]int, N)
		for j := 1; j <= m; j++ {
			f[i][j] = max(f[i-1][j], f[i][j-1])
			if a[i] == b[j] {
				f[i][j] = max(f[i][j], f[i-1][j-1]+1)
			}
		}
	}
	fmt.Println(f[n][m])
}

func (d *LineDp) MinDis() {
	reader := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscanln(reader, &n)
	a := make([]int32, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscanf(reader, "%c", &a[i])
	}
	fmt.Fscanln(reader)
	fmt.Fscanln(reader, &m)
	b := make([]int32, m+1)
	for i := 1; i <= m; i++ {
		fmt.Fscanf(reader, "%c", &b[i])
	}
	fmt.Fscanln(reader)

	N := 1010
	f := make([][]int, N)
	f[0] = make([]int, N)
	for i := 0; i <= m; i++ {
		f[0][i] = i
	}
	for i := 1; i <= n; i++ {
		f[i] = make([]int, N)
		f[i][0] = i
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			f[i][j] = d.min(f[i-1][j]+1, f[i][j-1]+1)
			if a[i] == b[j] {
				f[i][j] = d.min(f[i][j], f[i-1][j-1])
			} else {
				f[i][j] = d.min(f[i][j], f[i-1][j-1]+1)
			}
		}
	}
	fmt.Println(f[n][m])
}

func (d *LineDp) min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
