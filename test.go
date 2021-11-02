package main

import (
	"fmt"
	"math"
	"sync"
)

//1.半连接和长连接
//2.tcp建立连接后 accept listen
//3.equal

var (
	endx, endy, minstep int
	matrix, record      [][]int
	dx, dy              []int
)

func dfs(x int, y int, step int) {
	if x == endx && y == endy {
		if minstep > step {
			minstep = step
		}
		return
	}
	for i, v := range dx {
		cx := x + v
		cy := y + dy[i]
		if cx >= len(matrix) || cy >= len(matrix[0]) {
			return
		}
		if cx < 0 || cy < 0 {
			return
		}
		if record[cx][cy] == 0 && matrix[cx][cy] != 1 {
			record[cx][cy] = 1
			dfs(cx, cy, step+1)
			record[cx][cy] = 0
		}
	}
}

func main() {

	cha, chb, chc := make(chan int, 1), make(chan int, 1), make(chan int, 1)

	var wg sync.WaitGroup
	wg.Add(101)
	cha <- 1

	go func() {
		for {
			select {
			case i := <-cha:
				fmt.Println("a:", i)
				wg.Done()
				chb <- i + 1
				// wg.Done()
			}
		}
	}()

	go func() {
		for {
			select {
			case i := <-chb:
				fmt.Println("b:", i)
				chc <- i + 1
				wg.Done()
			}
		}
	}()

	go func() {
		for {
			select {
			case i := <-chc:
				fmt.Println("c:", i)
				wg.Done()
				cha <- i + 1
			}
		}
	}()

	wg.Wait()
	// time.Sleep(time.Millisecond * 40)

	// endx, endy = 3, 3
	// minstep = 9
	// for i := 0; i < 4; i++ {
	// 	tmp := make([]int, 4)
	// 	matrix = append(matrix, tmp)
	// 	record = append(record, tmp)
	// }
	// dx = []int{0, 1, 0, -1}
	// dy = []int{1, 0, -1, 0}
	// dfs(0, 0, 0)
	// // fmt.Print(minstep)

	// fmt.Print(getWater([]int{2, 1, 2, 3, 2, 1, 2, 3}))
}

func getWater(arr []int) int {
	n := len(arr)
	if n < 3 {
		return 0
	}
	res := 0
	lMax, rMax := arr[0], arr[n-1]
	l, r := 0, n-1
	for l != r {
		lMax = int(math.Max(float64(lMax), float64(arr[l])))
		rMax = int(math.Max(float64(rMax), float64(arr[r])))
		if lMax < rMax {
			res += lMax - arr[l]
			l++
		} else {
			res += rMax - arr[r]
			r--
		}
	}
	return res
}
