package dp

import (
	"bufio"
	"fmt"
	"os"
)

func JosephCycle() {
	reader := bufio.NewReader(os.Stdin)
	var (
		count int
		a     [1010]int
		n     int
		m     int
	)

	fmt.Fscanln(reader, &count)
	for count > 0 {
		count--
		fmt.Fscan(reader, &n, &m)
		for i := 0; i < m; i++ {
			fmt.Fscan(reader, &a[i])
		}

		res := 0 //起始位置f(1, (n-1)%m)
		for i, j := 1, (n-1)%m; i < n; {
			i++
			j = (j + m - 1) % m
			res = (res + a[j]) % i
		}
		fmt.Println(res)
	}

}
