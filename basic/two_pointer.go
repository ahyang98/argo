package basic

import (
	"bufio"
	"fmt"
	"os"
)

func maxNonRepeatSeq() {
	const N = 100010
	counts := make([]int, N)
	q := make([]int, N)
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanln(reader, &n)
	maxLen := 0
	for i, j := 0, 0; i < n; i++ {
		fmt.Fscan(reader, &q[i])
		counts[q[i]]++
		for j < i && counts[q[i]] > 1 {
			counts[q[j]]--
			j++
		}
		curLen := i - j + 1
		if curLen > maxLen {
			maxLen = curLen
		}
	}
	fmt.Println(maxLen)
}

func findTwoNum4Sum() {
	const N = 100010
	var a, b [N]int
	var n, m, x int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n, &m, &x)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	fmt.Fscanln(reader)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &b[i])
	}
	fmt.Fscanln(reader)

	for i, j := 0, m-1; i < n; i++ {
		for j >= 0 && a[i]+b[j] > x {
			j--
		}
		if a[i]+b[j] == x {
			fmt.Printf("%d %d", i, j)
			return
		}
	}
}

func findSubSeq() {
	const N = 100010
	var a, b [N]int
	var n, m int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n, &m)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	fmt.Fscanln(reader)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &b[i])
	}
	fmt.Fscanln(reader)
	i, j := 0, 0
	for i < n && j < m {
		if a[i] == b[j] {
			i++
		}
		j++
	}
	if i == n {
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")
}
