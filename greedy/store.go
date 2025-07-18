package greedy

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func MinDisSum() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanln(reader, &n)
	stores := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &stores[i])
	}

	sort.Ints(stores)

	mid := n / 2
	val := stores[mid]

	res := 0
	for i := 0; i < n; i++ {
		res = res + int(math.Abs(float64(stores[i]-val)))
	}
	fmt.Println(res)
}
