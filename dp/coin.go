package dp

import (
	"bufio"
	"fmt"
	"os"
)

func GetCoin2D() {
	reader := bufio.NewReader(os.Stdin)
	var n1, n2, m int
	fmt.Fscanln(reader, &n1, &n2, &m)
	const (
		N   = 110
		M   = 100010
		MOD = 1e9 + 7
	)

	var (
		v1 [N]int
		v2 [N]int
		f  [N][M]int
	)
	for i := 1; i <= n1; i++ {
		fmt.Fscan(reader, &v1[i])
	}
	fmt.Fscanln(reader)
	for i := 1; i <= n2; i++ {
		fmt.Fscan(reader, &v2[i])
	}

	for i := 1; i <= n1+n2; i++ {
		f[i][0] = 1
	}

	//f[i][j], 前i个物品中选价值为j的方案，
	//要区分i是那种物品即那种场景01还是完全背包
	//此题中就是普通币还是纪念币
	//所以分成两种场景，先把普通币在各种价值下的方案数算出来，再算纪念币中选

	// 先从普通币中选，完全背包问题 用j-v1[i]优化递推
	for i := 1; i <= n1; i++ {
		for j := 1; j <= m; j++ {
			if v1[i] <= j {
				f[i][j] = (f[i-1][j] + f[i][j-v1[i]]) % MOD
			} else {
				f[i][j] = f[i-1][j]
			}
		}
	}

	// 再从纪念币中选，01背包问题
	// 会使用到f[n1][j]，即在普通币已经选择的基础上再继续选
	for i := n1 + 1; i <= n1+n2; i++ {
		for j := 1; j <= m; j++ {
			if v2[i] <= j {
				//选i的方案数就等于把i硬币价值扣除掉后，选前i-1种硬币的数量，
				//如果前i-1种硬币在这个价值下没有方案，那么这个场景下也就是没有方案了。
				f[i][j] = (f[i-1][j] + f[i-1][j-v2[i-n1]]) % MOD
			} else {
				f[i][j] = f[i-1][j]
			}
		}
	}

	fmt.Println(f[n1+n2][m])
}

func GetCoin1D() {
	reader := bufio.NewReader(os.Stdin)
	var n1, n2, m int
	fmt.Fscanln(reader, &n1, &n2, &m)
	const (
		N   = 110
		M   = 100010
		MOD = 1e9 + 7
	)

	var (
		v1 [N]int
		v2 [N]int
		f  [M]int
	)
	for i := 1; i <= n1; i++ {
		fmt.Fscan(reader, &v1[i])
	}
	fmt.Fscanln(reader)
	for i := 1; i <= n2; i++ {
		fmt.Fscan(reader, &v2[i])
	}

	f[0] = 1

	// 使用1维优化，可以认为是滚动数组
	for i := 1; i <= n1; i++ {
		for j := 1; j <= m; j++ {
			if v1[i] <= j {
				f[j] = (f[j] + f[j-v1[i]]) % MOD
			}
		}
	}

	// 进入时f[1~m]中记录了上一轮的方案数，再进一步递推出选择纪念币的
	for i := 1; i <= n2; i++ {
		for j := m; j >= v2[i]; j-- {
			f[j] = (f[j] + f[j-v2[i]]) % MOD
		}
	}

	fmt.Println(f[m])
}
