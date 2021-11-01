package main

import (
	"fmt"
)

func main22() {
	var n, a, b, c, d int
	fmt.Scanf("%d", &n)
	fmt.Scanf("%d %d", &a, &b)

	dic := make(map[int][]int)
	for i := 0; i < a; i++ {
		var x, y int
		fmt.Scanf("%d %d", &x, &y)
		if nums, ok := dic[x]; ok {
			nums = append(nums, y)
			dic[x] = nums
		} else {
			tmp := make([]int, 0)
			tmp = append(tmp, y)
			dic[x] = tmp
		}
	}

	fmt.Scanf("%d %d", &c, &d)
	res := make([]int, c)
	for i := 0; i < c; i++ {
		var x, y int
		fmt.Scanf("%d %d", &x, &y)
		if isPre(x, y, dic) {
			res[i] = 1
		} else {
			res[i] = 0
		}
	}
	for i, num := range res {
		if i == 0 {
			fmt.Printf("[%d,", num)
		} else if i == c-1 {
			fmt.Printf("%d]", num)
		} else {
			fmt.Printf("%d,", num)
		}
	}
}

func isPre(x int, y int, dic map[int][]int) bool {
	nums := dic[y]
	return inNums(x, nums, dic)
}

func inNums(x int, nums []int, dic map[int][]int) bool {
	for _, num := range nums {
		if num == x {
			return true
		}
	}
	for _, n := range nums {
		return inNums(n, dic[n], dic)
	}
	return false
}

func isZnum(num int) bool {
	if num <= 2 {
		return true
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
