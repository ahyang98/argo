package basic

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Init() *ListNode {
	reader := bufio.NewReader(os.Stdin)
	//var line string
	//fmt.Fscanln(reader, &line)
	//line, _, _ := reader.ReadLine()
	line, _ := reader.ReadString('\n')
	strs := strings.Split(line[strings.Index(line, "[")+1:strings.Index(line, "]")], ", ")
	var head, p *ListNode
	for _, str := range strs {
		x, _ := strconv.Atoi(str)
		if head == nil {
			head = &ListNode{
				Val:  x,
				Next: nil,
			}
			p = head
		} else {
			p.Next = &ListNode{
				Val:  x,
				Next: nil,
			}
			p = p.Next
		}
	}
	return head
}

func (l *LinkedList) QuickSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	left := &ListNode{
		Val:  0,
		Next: nil,
	}
	ltail := left
	mid := &ListNode{
		Val:  0,
		Next: nil,
	}
	mtail := mid
	right := &ListNode{
		Val:  0,
		Next: nil,
	}
	rtail := right

	val := head.Val

	for head != nil {
		if head.Val == val {
			mtail.Next = head
			mtail = head
		} else if head.Val < val {
			ltail.Next = head
			ltail = head
		} else {
			rtail.Next = head
			rtail = head
		}
		head = head.Next
	}
	mtail.Next = nil
	ltail.Next = nil
	rtail.Next = nil

	left.Next = l.QuickSort(left.Next)
	right.Next = l.QuickSort(right.Next)
	l.getTail(left).Next = mid.Next
	l.getTail(left).Next = right.Next
	return left.Next
}

func (l *LinkedList) getTail(head *ListNode) *ListNode {
	for head != nil && head.Next != nil {
		head = head.Next
	}
	return head
}

func (l *LinkedList) Print(head *ListNode) {
	fmt.Print("[")
	for head != nil {
		fmt.Printf("%d, ", head.Val)
		head = head.Next
	}
	fmt.Print("]")
}

func entryNodeOfLoop(head *ListNode) *ListNode {
	degrees := make(map[*ListNode]int, 510)
	p := head
	for p != nil {
		degrees[p]++
		if degrees[p] > 1 {
			break
		}
		p = p.Next
	}
	return p
}

func reverseListRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := head

	for p.Next != nil {
		c := p.Next.Next
		p.Next.Next = head
		head = p.Next
		p.Next = c
	}

	return head
}

func reverseListAB(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var a *ListNode = nil
	b := head

	for b != nil {
		c := b.Next
		b.Next = a
		a = b
		b = c
	}

	return a
}
