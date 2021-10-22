package sort

import (
	"log"
)

func BubbleSort(nums []int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] {
				swap(nums, j, j+1)
			}
		}
	}
}

func QuickSort(nums []int) {
	quickSort(nums, 0, len(nums))
}

func quickSort(nums []int, left int, right int) {
	if left < right {
		pivot := getPivot(nums, left, right)
		quickSort(nums, left, pivot)
		quickSort(nums, pivot+1, right)
	}
}

func getPivot(nums []int, left int, right int) int {
	pivot := left
	index := left + 1
	for i := index; i < right; i++ {
		if nums[i] < nums[pivot] {
			swap(nums, i, index)
			index++
		}
	}
	log.Print(nums, index)
	swap(nums, pivot, index-1)
	return index - 1
}

func PickSort(nums []int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		minindex := i
		for j := i; j < n; j++ {
			if nums[j] < nums[minindex] {
				minindex = j
			}
		}
		swap(nums, i, minindex)
	}
}

func swap(nums []int, i int, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
