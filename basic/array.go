package basic

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type SnakeMatrix struct {
}

func NewSnakeMatrix() *SnakeMatrix {
	return &SnakeMatrix{}
}

func (s *SnakeMatrix) Run() {
	reader := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscanln(reader, &n, &m)
	dx := []int{-1, 0, 1, 0}
	dy := []int{0, 1, 0, -1}
	x, y := 0, 0
	d := 1
	q := make([][]int, n)
	for i := 0; i < n; i++ {
		q[i] = make([]int, m)
	}
	for i := 1; i <= n*m; i++ {
		q[x][y] = i
		a := x + dx[d]
		b := y + dy[d]
		if a < 0 || a >= n || b < 0 || b >= m || q[a][b] > 0 {
			d = (d + 1) % 4
		}
		x = x + dx[d]
		y = y + dy[d]
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("%d ", q[i][j])
		}
		fmt.Println()
	}
}

func reverseWords1(s string) string {
	words := strings.Split(s, " ")
	res := ""
	for i := len(words); i >= 0; i-- {
		res = res + words[i] + " "
	}
	return res[0 : len(res)-1]
}

func ReverseWords2(s string) string {
	byteStr := []byte(s)
	reverse(byteStr, 0, len(byteStr)-1)
	for i := 0; i < len(byteStr); i++ {
		j := i
		for j < len(byteStr) && byteStr[j] != ' ' {
			j++
		}
		reverse(byteStr, i, j-1)
		i = j
	}
	return string(byteStr)
}

func reverse(byteStr []byte, i, j int) {
	for i < j {
		byteStr[i], byteStr[j] = byteStr[j], byteStr[i]
		i++
		j--
	}
}
