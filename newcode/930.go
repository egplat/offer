package newcode

var matrix [][]int
var rcd [][]int
var dif [][]int
var res = 1<<31 - 1

func shortestPath(grid [][]int, k int) int {
	n := len(grid)
	m := len(grid[0])
	dif = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	matrix = grid
	rcd = make([][]int, 0)
	for i := 0; i < n; i++ {
		tmp := make([]int, m)
		rcd = append(rcd, tmp)
	}

	dfs(0, 0, 0)
}

func dfs(s int, e int, step int) {
}
