package main

import (
	"fmt"
	"sync"
)

type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}
func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
	return (s[i]) < (s[j])
}
func main1() {
	// abc := []string{"d", "c", "b"}
	// sort.Sort(ByLength(abc))
	// fmt.Println(abc)
	a := make([][]int, 10)
	// for i := 0; i < 10; i++ {
	// 	tmp := make([]int, 9)
	// 	a[i] = append(a[i], tmp...)
	// }
	a[0] = []int{1, 2}
	fmt.Print(a[0])

}

func test() {
	ss := []int{1, 2, 3, 4, 5}
	var wg sync.WaitGroup
	wg.Add(len(ss))
	for i := 0; i < len(ss); i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println(ss[i])
		}(i)
	}
	wg.Wait()
}
