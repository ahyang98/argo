package basic

import (
	"bufio"
	"fmt"
	"os"
)

type MonotonousStack struct {
}

func NewMonotonousStack() *MonotonousStack {
	return &MonotonousStack{}
}

func (s *MonotonousStack) Calc() {
	const N = 100010
	var (
		stk [N]int
	)
	tt := 0
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanln(reader, &n)
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(reader, &x)
		for tt > 0 && stk[tt] >= x {
			tt--
		}
		if tt == 0 {
			fmt.Print("-1 ")
		} else {
			fmt.Printf("%d ", stk[tt])
		}
		tt++
		stk[tt] = x
	}
}

type MinStack struct {
	stack    []int
	tt       int
	minStack []int
	minTt    int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{stack: make([]int, 110), minStack: make([]int, 110)}
}

func (this *MinStack) Push(x int) {
	this.tt++
	this.stack[this.tt] = x
	if this.minTt == 0 || this.minStack[this.minTt] > x {
		this.minTt++
		this.minStack[this.minTt] = x
	}
}

func (this *MinStack) Pop() {
	if this.minStack[this.minTt] == this.stack[this.tt] {
		this.minTt--
	}
	this.tt--
}

func (this *MinStack) Top() int {
	return this.stack[this.tt]
}

func (this *MinStack) GetMin() int {
	return this.minStack[this.minTt]
}
