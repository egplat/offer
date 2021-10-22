package baidu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNC157(t *testing.T) {
	data := []int{3, 4, 1, 5, 6, 2, 7}
	res := foundMonotoneStack(data)
	ac := [][]int{{-1, 2}, {0, 2}, {-1, -1}, {2, 5}, {3, 5}, {2, -1}, {5, -1}}
	assert.Equal(t, ac, res)
}
