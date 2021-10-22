package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func smain() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	st := scan.Text()
	sa := strings.Split(st, " ")
	n, _ := strconv.Atoi(sa[0])
	m, _ := strconv.Atoi(sa[1])

	nums := make([]int, 0)
	scan.Scan()
	st = scan.Text()
	sa = strings.Split(st, " ")
	for i := 0; i < n; i++ {
		v, _ := strconv.Atoi(sa[i])
		nums = append(nums, v)
	}
	res := getH(n, m, nums)
	fmt.Print(res)
}

func getH(n int, m int, nums []int) int {
	bubbleSort(nums)
	step := m
	res := nums[0]
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			continue
		}
		dif := nums[i+1] - nums[i]
		if step >= dif*(i+1) {
			res = nums[i+1]
			step -= dif * (i + 1)
		} else {
			break
		}
	}
	return res
}

func valIsOk(n int, m int, k int, nums []int) bool {
	res := false
	allnum := make([]int, 0)
	allnum = append(allnum, nums...)
	for i, v := range nums {
		nums[i] = pow3(v)
	}
	allnum = append(allnum, nums...)
	bubbleSort(allnum)
	cal := 0
	for _, v := range nums {
		if v > k {
			return false
		}
		cal += v
		if cal == k {
			return true
		}
		if cal > k {
			return false
		}
	}
	return res
}

func pow3(n int) int {
	return n * n * n
}

func bubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				swap(nums, j, j+1)
			}
		}
	}
}

func swap(nums []int, i int, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func pathIsOk(n int, m int, k int, link [][]int, log []int) bool {
	res := true
	pMap := make(map[int][]int)
	for i := 0; i < m; i++ {
		a := link[i][0]
		b := link[i][1]
		if _, ok := pMap[a]; ok {
			pMap[a] = append(pMap[a], b)
		} else {
			pMap[a] = []int{b}
		}
		if _, ok := pMap[b]; ok {
			pMap[b] = append(pMap[b], a)
		} else {
			pMap[b] = []int{a}
		}
	}
	path := make([]int, 0)
	for _, k := range log {
		path = append(path, pMap[k]...)
	}
	for _, k := range log {
		flag := false
		for j := 0; j < len(path); j++ {
			if path[j] == k {
				flag = true
				break
			}
		}
		if !flag {
			res = false
			break
		}
	}
	return res
}
