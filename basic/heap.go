package basic

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type HeapItem interface {
	Value() int
}

type ObjectMinHeap struct {
	q []HeapItem
}

func (h *ObjectMinHeap) Len() int {
	return len(h.q)
}

func (h *ObjectMinHeap) Less(i, j int) bool {
	return h.q[i].Value() < h.q[j].Value()
}

func (h *ObjectMinHeap) Swap(i, j int) {
	h.q[i], h.q[j] = h.q[j], h.q[i]
}

func (h *ObjectMinHeap) Push(x any) {
	h.q = append(h.q, x.(HeapItem))
}

func (h *ObjectMinHeap) Pop() any {
	x := h.q[len(h.q)-1]
	h.q = h.q[0 : len(h.q)-1]
	return x
}

type MinHeap struct {
	h *ObjectMinHeap
}

func NewMinHeap() *MinHeap {
	return &MinHeap{
		h: &ObjectMinHeap{q: make([]HeapItem, 0)},
	}
}

func (m *MinHeap) Push(x HeapItem) {
	heap.Push(m.h, x)
}

func (m *MinHeap) Pop() HeapItem {
	return heap.Pop(m.h).(HeapItem)
}

func (m *MinHeap) Len() int {
	return m.h.Len()
}

type ObjectMaxHeap struct {
	q []HeapItem
}

func (h *ObjectMaxHeap) Len() int {
	return len(h.q)
}

func (h *ObjectMaxHeap) Less(i, j int) bool {
	return h.q[i].Value() > h.q[j].Value()
}

func (h *ObjectMaxHeap) Swap(i, j int) {
	h.q[i], h.q[j] = h.q[j], h.q[i]
}

func (h *ObjectMaxHeap) Push(x any) {
	h.q = append(h.q, x.(HeapItem))
}

func (h *ObjectMaxHeap) Pop() any {
	x := h.q[len(h.q)-1]
	h.q = h.q[0 : len(h.q)-1]
	return x
}

type MaxHeap struct {
	h *ObjectMaxHeap
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		h: &ObjectMaxHeap{q: make([]HeapItem, 0)},
	}
}

func (m *MaxHeap) Push(x HeapItem) {
	heap.Push(m.h, x)
}

func (m *MaxHeap) Pop() HeapItem {
	return heap.Pop(m.h).(HeapItem)
}

func (m *MaxHeap) Peek() HeapItem {
	item := heap.Pop(m.h).(HeapItem)
	heap.Push(m.h, item)
	return item
}

func (m *MaxHeap) Len() int {
	return m.h.Len()
}

type IntItem int

func (i IntItem) Value() int {
	return int(i)
}

func DynamicMiddle() {
	var p int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &p)
	for p > 0 {
		down := NewMaxHeap()
		up := NewMinHeap()
		p--
		var id, n int
		fmt.Fscanln(reader, &id, &n)
		fmt.Println(id, (n+1)/2)
		iCnt := 10
		oCnt := 0
		var x int
		for i := 0; i < n; i++ {
			if iCnt == 0 {
				iCnt = 10
				fmt.Fscanln(reader)
			}
			fmt.Fscanf(reader, "%d", &x)
			iCnt--
			item := IntItem(x)
			if down.Len() < 1 {
				down.Push(item)
			} else {
				mid := down.Peek()
				if x <= mid.Value() {
					down.Push(item)
				} else {
					up.Push(item)
				}
			}

			if down.Len() > up.Len()+1 {
				up.Push(down.Pop())
			}

			if up.Len() > down.Len() {
				down.Push(up.Pop())
			}
			if i%2 == 0 {
				oCnt++
				fmt.Printf("%d ", down.Peek().Value())
				if oCnt%10 == 0 {
					fmt.Println()
				}
			}
		}
		if iCnt%10 > 0 {
			fmt.Fscanln(reader)
		}
		if oCnt%10 > 0 {
			fmt.Println()
		}
	}
}
