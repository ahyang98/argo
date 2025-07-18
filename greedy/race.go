package greedy

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type intSlice []int

func (x intSlice) Len() int           { return len(x) }
func (x intSlice) Less(i, j int) bool { return x[i] > x[j] }
func (x intSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func Race() {
	reader := bufio.NewReader(os.Stdin)
	n := -1
	for {
		fmt.Fscanln(reader, &n)
		if n == 0 {
			break
		}
		a := make(intSlice, n)
		b := make(intSlice, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
		}
		fmt.Fscanln(reader)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &b[i])
		}
		fmt.Fscanln(reader)
		sort.Sort(a)
		sort.Sort(b)
		doRace(n, a, b)
	}
}

func doRace(n int, a, b intSlice) {
	var (
		f1  = 0
		f2  = 0
		s1  = n - 1
		s2  = n - 1
		res = 0
	)

	for f1 <= s1 {
		if a[s1] > b[s2] {
			res++
			s1--
			s2--
		} else if a[s1] < b[s2] {
			res--
			s1--
			f2++
		} else {
			if a[f1] > b[f2] {
				f1++
				f2++
				res++
			} else if a[f1] < b[f2] {
				s1--
				f2++
				res--
			} else {
				s1--
				f2++
				if a[s1] < b[f2] {
					res--
				}
			}
		}
	}

	fmt.Println(res * 200)

}
