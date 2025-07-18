package dp

import (
	"bufio"
	"fmt"
	"os"
)

type Bag struct {
}

func NewBag() *Bag {
	return &Bag{}
}

func (b *Bag) Result() {
	b.d1()
}

func (b *Bag) d2() {
	reader := bufio.NewReader(os.Stdin)
	N := 1010
	var n, m int
	fmt.Fscanln(reader, &n, &m)
	f := make([][]int, N)
	// 体积
	v := make([]int, N)
	// 价值
	w := make([]int, N)

	for i := 1; i <= n; i++ {
		fmt.Fscanln(reader, &v[i], &w[i])
	}

	f[0] = make([]int, N)

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if len(f[i]) < 1 {
				f[i] = make([]int, N)
			}
			f[i][j] = f[i-1][j]
			if j >= v[i] {
				f[i][j] = b.max(f[i][j], f[i-1][j-v[i]]+w[i])
			}
		}
	}

	fmt.Println(f[n][m])
}

func (b *Bag) d1() {
	reader := bufio.NewReader(os.Stdin)
	N := 1010
	var n, m int
	fmt.Fscanln(reader, &n, &m)
	f := make([]int, N)
	// 体积
	v := make([]int, N)
	// 价值
	w := make([]int, N)

	for i := 1; i <= n; i++ {
		fmt.Fscanln(reader, &v[i], &w[i])
	}

	for i := 1; i <= n; i++ {
		for j := m; j >= v[i]; j-- {
			f[j] = b.max(f[j], f[j-v[i]]+w[i])
		}
	}

	fmt.Println(f[m])
}

func (b *Bag) max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func (b *Bag) InfBag2() {
	reader := bufio.NewReader(os.Stdin)
	N := 1010
	var n, m int
	fmt.Fscanln(reader, &n, &m)
	//f[i][j]: 只看前i个物品，背包容量为j的所有方案，取最大值
	f := make([][]int, N)
	// 体积
	v := make([]int, N)
	// 价值
	w := make([]int, N)

	for i := 1; i <= n; i++ {
		fmt.Fscanln(reader, &v[i], &w[i])
	}

	f[0] = make([]int, N)

	for i := 1; i <= n; i++ {
		for j := 0; j <= m; j++ {
			if len(f[i]) < 1 {
				f[i] = make([]int, N)
			}
			f[i][j] = f[i-1][j]
			for k := 0; k*v[i] <= j; k++ {
				f[i][j] = b.max(f[i][j], f[i-1][j-k*v[i]]+k*w[i])
			}
		}
	}

	fmt.Println(f[n][m])
}

func (b *Bag) InfBag1() {
	//j-v 替换 j  f[i][j]=max(f[i][j-v]+w, f[i-1][j])
	reader := bufio.NewReader(os.Stdin)
	N := 1010
	var n, m int
	fmt.Fscanln(reader, &n, &m)
	//f[i][j]: 只看前i个物品，背包容量为j的所有方案，取最大值
	//f := make([][]int, N)
	f := make([]int, N)
	// 体积
	v := make([]int, N)
	// 价值
	w := make([]int, N)

	for i := 1; i <= n; i++ {
		fmt.Fscanln(reader, &v[i], &w[i])
	}

	//f[0] = make([]int, N)

	for i := 1; i <= n; i++ {
		for j := v[i]; j <= m; j++ {
			//if len(f[i]) < 1 {
			//	f[i] = make([]int, N)
			//}
			//f[i][j] = b.max(f[i-1][j], f[i][j-v[i]]+w[i])
			f[j] = b.max(f[j], f[j-v[i]]+w[i])
		}
	}

	fmt.Println(f[m])
}

func (b *Bag) LimitBag1() {
	reader := bufio.NewReader(os.Stdin)
	N := 1010
	var n, m int
	fmt.Fscanln(reader, &n, &m)
	//f[i][j]: 只看前i个物品，背包容量为j的所有方案，取最大值
	f := make([][]int, N)
	// 体积
	v := make([]int, N)
	// 价值
	w := make([]int, N)

	// 每个类别物品的数量
	s := make([]int, N)

	for i := 1; i <= n; i++ {
		fmt.Fscanln(reader, &v[i], &w[i], &s[i])
	}

	f[0] = make([]int, N)

	for i := 1; i <= n; i++ {
		for j := 0; j <= m; j++ {
			if len(f[i]) < 1 {
				f[i] = make([]int, N)
			}
			f[i][j] = f[i-1][j]
			for k := 0; k*v[i] <= j && k <= s[i]; k++ {
				f[i][j] = b.max(f[i][j], f[i-1][j-k*v[i]]+k*w[i])
			}
		}
	}

	fmt.Println(f[n][m])
}

func (b *Bag) LimitBag2() {
	reader := bufio.NewReader(os.Stdin)
	N := 1010
	var n, m int
	fmt.Fscanln(reader, &n, &m)
	f := make([]int, N)
	// 体积
	v := make([]int, N)
	// 价值
	w := make([]int, N)

	cnt := 0
	for i := 1; i <= n; i++ {
		var x, y, s int
		fmt.Fscanln(reader, &x, &y, &s)
		k := 1
		for k <= s {
			cnt++
			v[cnt] = k * x
			w[cnt] = k * y
			s = s - k
			k = k * 2
		}
		if s > 0 {
			cnt++
			v[cnt] = x * s
			w[cnt] = y * s
		}
	}

	n = cnt

	for i := 1; i <= n; i++ {
		for j := m; j >= v[i]; j-- {
			f[j] = b.max(f[j], f[j-v[i]]+w[i])
		}
	}

	fmt.Println(f[m])
}

func (b *Bag) GroupBag() {
	reader := bufio.NewReader(os.Stdin)
	N := 1010
	var n, m int
	fmt.Fscanln(reader, &n, &m)
	//f[i][j]: 只看前i个物品，背包容量为j的所有方案，取最大值
	f := make([]int, N)
	// 体积
	v := make([][]int, N)
	// 价值
	w := make([][]int, N)

	// 每个类别物品的数量
	s := make([]int, N)

	for i := 1; i <= n; i++ {
		fmt.Fscanln(reader, &s[i])
		v[i] = make([]int, N)
		w[i] = make([]int, N)
		for j := 0; j < s[i]; j++ {
			fmt.Fscanln(reader, &v[i][j], &w[i][j])
		}
	}

	for i := 1; i <= n; i++ {
		for j := m; j >= 0; j-- {
			for k := 0; k < s[i]; k++ {
				if j >= v[i][k] {
					f[j] = b.max(f[j], f[j-v[i][k]]+w[i][k])
				}
			}
		}
	}

	fmt.Println(f[m])
}

type XOR struct {
}

func NewXOR() *XOR {
	return &XOR{}
}

func (o XOR) isPrime(x int) bool {
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func (o XOR) Calc() {
	const (
		N   = 5010
		M   = 8192
		MOD = 1e9 + 7
	)

	var (
		n int
		a [N]int
		f [2][M]int
	)
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	f[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j < M; j++ {
			f[i&1][j] = f[(i-1)&1][j]
			if j^a[i] < M {
				f[i&1][j] = (f[i&1][j] + f[(i-1)&1][j^a[i]]) % MOD
			}
		}
	}
	res := 0
	for i := 2; i < M; i++ {
		if o.isPrime(i) {
			res = (res + f[n&1][i]) % MOD
		}
	}

	fmt.Printf("%d", res)
}
