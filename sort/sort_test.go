package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	data := []int{3, 5, 2, 1, 6, 9, 8, 4, 0, 7}
	BubbleSort(data)
	exp := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	assert.Equal(t, exp, data)
}

func TestQuickSort(t *testing.T) {
	data := []int{3, 5, 2, 1, 6, 9, 8, 4, 0, 7}
	QuickSort(data)
	exp := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	assert.Equal(t, exp, data)
}

func TestPickSort(t *testing.T) {
	data := []int{3, 5, 2, 1, 6, 9, 8, 4, 0, 7}
	PickSort(data)
	exp := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	assert.Equal(t, exp, data)
}
