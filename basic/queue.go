package basic

import (
	"bufio"
	"fmt"
	"os"
)

type SlideWindow struct {
}

func NewSlideWindow() *SlideWindow {
	return &SlideWindow{}
}

func (w SlideWindow) Calc() {
	const N = 1000010
	var (
		q [N]int
		a [N]int
	)

	hh, tt := 0, -1
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	var n, k int
	fmt.Fscanln(reader, &n, &k)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	fmt.Fscanln(reader)
	for i := 0; i < n; i++ {
		if hh <= tt && (i-k+1) > q[hh] {
			hh++
		}

		for hh <= tt && a[q[tt]] >= a[i] {
			tt--
		}
		tt++
		q[tt] = i
		if i >= k-1 {
			fmt.Fprintf(writer, "%d ", a[q[hh]])
		}
	}
	fmt.Fprintln(writer)
	hh, tt = 0, -1
	for i := 0; i < n; i++ {
		if hh <= tt && (i-k+1) > q[hh] {
			hh++
		}

		for hh <= tt && a[q[tt]] <= a[i] {
			tt--
		}
		tt++
		q[tt] = i
		if i >= k-1 {
			fmt.Fprintf(writer, "%d ", a[q[hh]])
		}
	}
	fmt.Fprintln(writer)
	writer.Flush()
}
