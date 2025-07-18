package basic

import (
	"bufio"
	"fmt"
	"os"
)

type Stock3 struct {
}

func NewStock3() *Stock3 {
	return &Stock3{}
}

func (s *Stock3) Calc() {
	reader := bufio.NewReader(os.Stdin)
	const N = 100010
	var (
		n int
		p [N]int
		g [N]int //1~i天之内交易完成的最大收益
		f [N]int //i+1~n交易的最大收益
	)

	fmt.Fscanln(reader, &n)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &p[i])
	}

	for i, minv := 2, p[1]; i <= n; i++ {
		//分为第i天卖出和不是第i天卖出
		f[i] = s.max(f[i-1], p[i]-minv)
		minv = s.min(minv, p[i])
	}

	for i, maxv := n-1, p[n]; i >= 2; i-- {
		//分为第i天买入和不是第i天买入
		g[i] = s.max(g[i+1], maxv-p[i])
		maxv = s.max(maxv, p[i])
	}

	res := 0
	for i := 2; i <= n; i++ {
		res = s.max(res, f[i]+g[i+1])
	}
	fmt.Println(res)

}

func (s *Stock3) max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (s *Stock3) min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
