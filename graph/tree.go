package graph

import (
	"bufio"
	"fmt"
	"os"
)

type Tree struct {
	val   uint8
	left  *Tree
	right *Tree
}

type BuildTree struct {
	middleMap map[uint8]int
}

func NewBuildTree() *BuildTree {
	return &BuildTree{
		middleMap: make(map[uint8]int, 26),
	}
}

func (t *BuildTree) Init() {
	reader := bufio.NewReader(os.Stdin)
	var middle, level string
	fmt.Fscanln(reader, &middle)
	fmt.Fscanln(reader, &level)
	for i := 0; i < len(middle); i++ {
		t.middleMap[middle[i]] = i
	}

	root := t.build(level)
	t.printTree(root)
}

func (t *BuildTree) build(level string) *Tree {
	if len(level) < 1 {
		return nil
	}
	node := &Tree{
		val:   level[0],
		left:  nil,
		right: nil,
	}

	if len(level) == 1 {
		return node
	}

	root := t.middleMap[level[0]]
	var left, right string

	for i := 1; i < len(level); i++ {
		if t.middleMap[level[i]] < root {
			left = left + string(level[i])
		}
		if t.middleMap[level[i]] > root {
			right = right + string(level[i])
		}
	}
	node.left = t.build(left)
	node.right = t.build(right)
	return node
}

func (t *BuildTree) printTree(root *Tree) {
	if nil == root {
		return
	}

	fmt.Print(string(root.val))
	t.printTree(root.left)
	t.printTree(root.right)
}

func (t *BuildTree) bfs(level string) *Tree {
	var (
		st = make([]bool, 26)
		q  = make([]*Tree, 26)
	)
	q[0] = &Tree{val: level[0]}
	i := 0
	for j := 1; j < len(level); {
		end := j
		for ; i < end; i++ {
			p := t.middleMap[level[i]]
			st[p] = true
			if p > 0 && !st[p-1] {
				q[i].left = &Tree{val: level[j]}
				q[j] = q[i].left
				j++
			}
			if p+1 < len(level) && !st[p+1] {
				q[i].right = &Tree{val: level[j]}
				q[j] = q[i].right
				j++
			}
		}
	}
	return q[0]
}

// TreeNode Definition for a binary tree node.
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func BuildTreePI(preorder []int, inorder []int) *TreeNode {
	const N = 110
	orderMap := make(map[int]int, N)
	for i := 0; i < len(inorder); i++ {
		orderMap[inorder[i]] = i
	}

	return dfs(0, len(preorder)-1, 0, len(inorder)-1, preorder, orderMap)
}

func dfs(pl, pr, il, ir int, preorder []int, inorderMap map[int]int) *TreeNode {
	if pl > pr {
		return nil
	}
	// 左子树的长度
	k := inorderMap[preorder[pl]] - il

	root := &TreeNode{
		Val: preorder[pl],
	}

	root.Left = dfs(pl+1, pl+k, il, il+k-1, preorder, inorderMap)
	root.Right = dfs(pl+k+1, pr, il+k+1, ir, preorder, inorderMap)
	return root
}

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val    int
	Left   *TreeNode
	Right  *TreeNode
	Father *TreeNode
}

func inorderSuccessor(p *TreeNode) *TreeNode {
	if p.Right != nil {
		p = p.Right
		for p.Left != nil {
			p = p.Left
		}
		return p
	}
	for p.Father != nil && p == p.Father.Right {
		p = p.Father
	}
	return p.Father
}
