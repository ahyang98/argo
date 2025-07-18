package greedy

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func MoveCard() {
	reader := bufio.NewReader(os.Stdin)
	var n, x, avg, sum int
	fmt.Fscanln(reader, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		sum += a[i]
	}
	avg = sum / n
	res := 0
	for i := 1; i < n; i++ {
		x = a[i] - avg + x
		if x != 0 {
			res++
		}
	}

	fmt.Println(res)
}

type int64Slice []int64

func (x int64Slice) Len() int           { return len(x) }
func (x int64Slice) Less(i, j int) bool { return x[i] > x[j] }
func (x int64Slice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func CycleMove() {
	const N = 1e6 + 10
	reader := bufio.NewReader(os.Stdin)
	var (
		n   int
		avg int64
	)
	fmt.Fscanln(reader, &n)
	s := make([]int64, N)
	c := make(int64Slice, N)

	for i := 1; i <= n; i++ {
		fmt.Fscanln(reader, &s[i])
		s[i] = s[i] + s[i-1]
	}
	avg = s[n] / int64(n)
	//C[2]~C[n+1] n个
	for i := 2; i <= n; i++ {
		c[i] = (int64(i)-1)*avg - (s[i] - s[1])
	}
	c = c[2 : n+2]
	sort.Sort(c)
	mid := c[n/2]
	var res int64 = 0
	for i := 0; i < n; i++ {
		res += int64(math.Abs(float64(c[i] - mid)))
	}
	fmt.Println(res)
}
