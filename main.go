package main

import (
	"container/heap"
	"container/list"
	"fmt"
	"math"
)

func main111() {
	ch := make(chan int)
	go func() {
		tmp := 0
		if _, ok := <-ch; ok {
			tmp = <-ch
		}
		if tmp+1 > 100 {
			return
		}
		fmt.Println(tmp + 1)
		ch <- tmp + 1
	}()
	go func() {
		tmp := 0
		if _, ok := <-ch; ok {
			tmp = <-ch
		}
		if tmp+1 > 100 {
			return
		}
		fmt.Println(tmp + 1)
		ch <- tmp + 1
	}()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getLongest(s string) int {
	rc := make(map[byte]bool)
	res := 0
	r := -1
	for i := 0; i < len(s); i++ {
		if i != 0 {
			delete(rc, s[i-1])
		}
		for r+1 < len(s) && !rc[s[r+1]] {
			rc[s[r+1]] = true
			r++
		}
		res = int(math.Max(float64(res), float64(r-i+1)))
	}
	return res
}

func rv(l *ListNode) *ListNode {
	head := new(ListNode)
	assis := head
	head.Next = l
	cur := l
	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			for cur.Next != nil {
				if cur.Next.Val == cur.Val {
					cur = cur.Next
				} else {
					break
				}
			}
			cur = cur.Next
			assis.Next = cur
		} else {
			assis = cur
			cur = cur.Next
		}
	}
	return head.Next
}

/**
 *
 * @param pHead ListNode类
 * @return ListNode类
 */

func ReverseList(pHead *ListNode) *ListNode {
	// write code here
	var head *ListNode
	next := new(ListNode)
	for pHead != nil {
		next = pHead.Next
		pHead.Next = head
		head = pHead
		pHead = next
	}
	return head
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 *
 * @param root TreeNode类 the root of binary tree
 * @return int整型二维数组
 */
func threeOrders(root *TreeNode) [][]int {
	// write code here
	res := make([][]int, 3)
	Order(root, res)
	return res
}

func Order(root *TreeNode, arr [][]int) {
	if root != nil {
		arr[0] = append(arr[0], root.Val)
		Order(root.Left, arr)
		arr[1] = append(arr[1], root.Val)
		Order(root.Right, arr)
		arr[2] = append(arr[2], root.Val)
	}
}

type ByLen []int

func (s ByLen) Len() int {
	return len(s)
}
func (s ByLen) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByLen) Less(i, j int) bool {
	return (s[i]) > (s[j])
}

func (h *ByLen) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *ByLen) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func GetLeastNumbers_Solution(input []int, k int) []int {
	// write code here
	in := ByLen(input)
	heap.Init(&in)
	fmt.Print(in)
	res := make([]int, 0)
	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(&in).(int))
		fmt.Print(in)
	}
	return res
}

func main2() {
	GetLeastNumbers_Solution([]int{1, 5, 2, 3}, 4)
}

func levelOrder(root *TreeNode) [][]int {
	// write code here
	stack := list.New()
	stack.Init()
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	stack.PushBack(root)
	tmp := make([]int, 0)
	lnode, llnode := root, root
	for stack.Len() > 0 {
		tnode := stack.Front().Value.(*TreeNode)
		stack.Remove(stack.Front())
		tmp = append(tmp, tnode.Val)
		if tnode.Left != nil {
			stack.PushBack(tnode.Left)
			llnode = tnode.Left
		}
		if tnode.Right != nil {
			stack.PushBack(tnode.Right)
			llnode = tnode.Right
		}
		if tnode == lnode {
			lnode = llnode
			res = append(res, tmp)
			tmp = []int{}
		}
	}
	return res
}

func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// write code here
	head := new(ListNode)
	lhead := head
	for pHead1 != nil && pHead2 != nil {
		num1 := pHead1.Val
		num2 := pHead2.Val
		if num1 < num2 {
			lhead.Next = pHead1
			lhead = lhead.Next
			pHead1 = pHead1.Next
		} else {
			lhead.Next = pHead2
			lhead = lhead.Next
			pHead2 = pHead2.Next
		}
	}
	for pHead1 != nil {
		lhead.Next = pHead1
		lhead = lhead.Next
		pHead1 = pHead1.Next
	}
	for pHead2 != nil {
		lhead.Next = pHead2
		lhead = lhead.Next
		pHead2 = pHead1.Next
	}
	return head.Next
}
