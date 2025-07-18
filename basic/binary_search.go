package basic

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func FindNums() {
	const N = 100010
	nums := make([]int, N)
	reader := bufio.NewReader(os.Stdin)
	var n, q int
	fmt.Fscanln(reader, &n, &q)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &nums[i])
	}
	fmt.Fscanln(reader)

	var x int
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < q; i++ {
		fmt.Fscanln(reader, &x)
		start := getStart(nums, x, n)
		if start == -1 {
			fmt.Fprintln(writer, "-1 -1")
		} else {
			fmt.Fprintln(writer, start, getEnd(nums, x, n))
		}
	}
	writer.Flush()
}

func getEnd(nums []int, x, n int) int {
	l, r := 0, n-1

	for l < r {
		mid := (l + r + 1) / 2
		if nums[mid] > x {
			r = mid - 1
		} else {
			l = mid
		}
	}
	if nums[l] == x {
		return l
	}
	return -1
}

func getStart(nums []int, x, n int) int {
	l, r := 0, n-1
	for l < r {
		mid := (l + r) / 2
		if nums[mid] < x {
			l = mid + 1
		} else {
			r = mid
		}
	}
	if nums[l] == x {
		return l
	}
	return -1
}

type MinOfMatrix struct {
	matrix [][]int
}

func NewMinOfMatrix() *MinOfMatrix {
	return &MinOfMatrix{matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}}
}

func (m *MinOfMatrix) query(a int, b int) int {
	return m.matrix[a][b]
}

func (m *MinOfMatrix) GetMinimumValue(N int) []int {
	const INF = math.MaxInt64
	l := 0
	r := N - 1
	for l < r {
		mid := (l + r) >> 1
		var val = INF
		k := 0
		for i := 0; i < N; i++ {
			x := m.query(i, mid)
			if x < val {
				val = x
				k = i
			}
		}
		var left = INF
		if mid > 0 {
			left = m.query(k, mid-1)
		}

		var right = INF
		if mid+1 < N {
			right = m.query(k, mid+1)
		}

		if val < left && val < right {
			return []int{k, mid}
		}
		if left < val {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	val := INF
	k := 0
	for i := 0; i < N; i++ {
		x := m.query(i, r)
		if x < val {
			val = x
			k = i
		}
	}
	return []int{k, r}
}

type PeakEle struct {
	nums []int
}

func NewPeakEle() *PeakEle {
	return &PeakEle{
		nums: []int{1, 2, 3, 1},
	}
}

func (p *PeakEle) query(x int) int {
	return p.nums[x]
}

func (p *PeakEle) Len() int {
	return len(p.nums)
}

func (p *PeakEle) FindPeakElement(N int) int {
	l := 0
	r := N - 1
	for l < r {
		mid := (l + r) >> 1

		if p.query(mid) > p.query(mid+1) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
