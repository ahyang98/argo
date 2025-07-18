package greedy

import (
	"bufio"
	"fmt"
	"os"
)

/*
1. 不相交生成新的
2. 相交扩展右端点
3. 包含关系保持不变
*/
func rangeMerge() {
	const (
		N       = 100010
		INVALID = -1 * 10e9
	)

	pairs := make([][]int, N)
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanln(reader, &n)
	for i := 0; i < n; i++ {
		pairs[i] = make([]int, 2)
		fmt.Fscanln(reader, &pairs[i][0], &pairs[i][1])
	}

	quickSort(pairs, 0, n-1)

	var rootL int = INVALID
	var rootR int = INVALID
	count := 0
	for i := 0; i < n; i++ {
		l := pairs[i][0]
		r := pairs[i][1]
		if rootL == INVALID || l > rootR {
			rootL = l
			rootR = r
			count++
		}

		if l <= rootR && r > rootR {
			rootR = r
		}
	}
	fmt.Println(count)
}

func quickSort(pairs [][]int, l, r int) {
	if l >= r {
		return
	}

	mid := (l + r) / 2
	x := pairs[mid][0]
	i := l - 1
	j := r + 1
	for i < j {
		for {
			i++
			if pairs[i][0] >= x {
				break
			}
		}
		for {
			j--
			if pairs[j][0] <= x {
				break
			}
		}
		if i < j {
			swap(pairs, i, j)
		}
	}

	quickSort(pairs, l, j)
	quickSort(pairs, j+1, r)
}

func swap(pairs [][]int, i int, j int) {
	var tmp [2]int
	tmp[0] = pairs[i][0]
	tmp[1] = pairs[i][1]
	pairs[i][0] = pairs[j][0]
	pairs[i][1] = pairs[j][1]
	pairs[j][0] = tmp[0]
	pairs[j][1] = tmp[1]
}
