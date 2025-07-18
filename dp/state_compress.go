package dp

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type StatusCompressDp struct {
}

func NewStatusCompressDp() *StatusCompressDp {
	return &StatusCompressDp{}
}

func (d *StatusCompressDp) Dp() {
	d.MendelianDp1()
}

func (d *StatusCompressDp) MendelianDp1() {
	const (
		N = 12
		// M 每一列的每一个空格有两种选择，放和不放，所以是2^n
		M = 1 << N
	)

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	var (
		n, m  int
		state = make(map[int]map[int]int) //存储所有合法状态
		//第 i-2 列伸到 i-1 列的状态为 k ， 是否能成功转移到 第 i-1 列伸到 i 列的状态为 j
		//st[j|k]=true 表示能成功转移
		st [M]bool   // 这一列空着的是不是偶数个
		f  [N][M]int //f[N][M] f[i][j]表示 i-1列的方案数已经确定，从i-1列伸出，并且第i列的状态是j的所有方案数
	)
	for {
		fmt.Fscanln(reader, &n, &m)
		if n == 0 || m == 0 {
			break
		}
		//第一部分：预处理1
		//对于每种状态，先预处理每列不能有奇数个连续的0
		for i := 0; i < 1<<n; i++ {
			isValid := true // 某种状态没有奇数个连续的0则标记为true
			cnt := 0
			for j := 0; j < n; j++ {
				//i >> j位运算，表示i（i在此处是一种状态）的二进制数的第j位；
				// &1为判断该位是否为1
				if (i >> j & 1) == 1 {
					//这一位为1，看前面连续的0的个数，如果是奇数（cnt &1为真）则该状态不合法
					if cnt&1 > 0 {
						isValid = false
						break
					}
					// 既然该位是1，并且前面不是奇数个0（经过上面的if判断），计数器清零。
					cnt = 0
				} else {
					cnt++
				}
			}
			//最下面的那一段判断一下连续的0的个数
			if cnt&1 > 0 {
				isValid = false
			}
			//状态i是否有奇数个连续的0的情况,输入到数组st中
			st[i] = isValid
		}

		//第二部分：预处理2
		// 经过上面每种状态 连续0的判断，已经筛掉一些状态。
		//下面来看进一步的判断：看第i-2列伸出来的和第i-1列伸出去的是否冲突
		for j := 0; j < 1<<n; j++ {
			delete(state, j)
			state[j] = make(map[int]int)
			for k := 0; k < 1<<n; k++ {
				// 第i-2列伸出来的 和第i-1列伸出来的不冲突(不在同一行)
				// 考虑的是第i-1列（第i-1列是这里的主体）中从第i-2列横插过来的 合并后 即或在一起
				// 这个 j|k 就是当前 第i-1列的到底有几个1，即哪几行是横着放格子的
				if (j&k) == 0 && st[j|k] {
					state[j][k] = k
				}
			}
		}

		//按定义这里是：前第-1列都摆好，且从-1列到第0列伸出来的状态为0的方案数。
		//首先，这里没有-1列，最少也是0列。
		//其次，没有伸出来，即没有横着摆的。即这里第0列只有竖着摆这1种状态。
		f[0][0] = 1

		for i := 1; i <= m; i++ {
			for j := 0; j < 1<<n; j++ {
				f[i][j] = 0
				for k, _ := range state[j] {
					f[i][j] += f[i-1][k]
				}
			}
		}

		fmt.Fprintln(writer, f[m][0])
	}
	writer.Flush()
}

func (d *StatusCompressDp) Hamilton() {
	reader := bufio.NewReader(os.Stdin)
	const (
		N = 20
		M = 1 << N
	)
	var (
		n int
		w [N][N]int
		f [M][N]int
	)

	fmt.Fscanln(reader, &n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &w[i][j])
		}
		fmt.Fscanln(reader)
	}

	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			f[i][j] = math.MaxInt
		}
	}

	f[1][0] = 0

	for i := 0; i < 1<<n; i++ {
		for j := 0; j < n; j++ {
			if i>>j&1 == 1 {
				for k := 0; k < n; k++ {
					if i>>k&1 == 1 {
						f[i][j] = d.min(f[i][j], f[i-(1<<j)][k]+w[k][j])
					}
				}
			}
		}
	}
	fmt.Println(f[1<<n-1][n-1])

}

func (d *StatusCompressDp) min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
