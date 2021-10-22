package baidu

import (
	"container/list"
)

func foundMonotoneStack(nums []int) [][]int {
	stack := list.New()
	n := len(nums)
	res := make([][]int, 0)
	for i := 0; i < n; i++ {
		res = append(res, []int{-1, -1})
	}
	for i := 0; i < n; i++ {
		tmp := nums[i]
		if stack.Len() != 0 {
			for stack.Back() != nil {
				min := stack.Back().Value.([]int)
				if tmp >= min[1] {
					break
				}
				res[min[0]][1] = i
				stack.Remove(stack.Back())
			}
		}
		stack.PushBack([]int{i, tmp})
	}
	stack.Init()
	for i := n - 1; i >= 0; i-- {
		tmp := nums[i]
		if stack.Len() != 0 {
			for stack.Back() != nil {
				min := stack.Back().Value.([]int)
				if tmp >= min[1] {
					break
				}
				res[min[0]][0] = i
				stack.Remove(stack.Back())
			}
		}
		stack.PushBack([]int{i, tmp})
	}
	return res
}
