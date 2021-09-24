package sort

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

func quickSort(nums []int, left int, rigth int) {
	if left < rigth {
		pivot := getPivot(nums, left, rigth)
		quickSort(nums, left, pivot)
		quickSort(nums, pivot+1, rigth)
	}
}

func getPivot(nums []int, left int, right int) int {
	index := 0
	for i := left; i < right; i++ {

	}
}

func swap(nums []int, i int, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
