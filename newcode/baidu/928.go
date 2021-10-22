package baidu

import (
	"container/list"
	"fmt"
)

func main5() {
	var a, b string
	for {
		if _, err := fmt.Scan(&a, &b); err != nil {
			break
		}
		la := len(a)
		lb := len(b)
		res := 0
		dic := make(map[string]bool)
		for i := 0; i <= la-lb; i++ {
			tmp := 0
			for j := 0; j < lb; j++ {
				if a[i+j] == b[j] || b[j] == '?' {
					tmp++
					continue
				}
				break
			}
			if tmp == lb {
				tmps := a[i : i+lb]
				if _, ok := dic[tmps]; !ok {
					res++
					dic[tmps] = true
				}
			}
		}
		fmt.Println(res)
	}
}

func main8() {
	var a string
	for {
		if _, err := fmt.Scan(&a); err != nil {
			break
		}
		stack := list.New()
		res := list.New()
		n := len(a)
		left := byte('[')
		right := byte(']')
		for i := 0; i < n; i++ {
			if a[i] == left {
				stack.PushFront(left)
				res.PushFront(left)
			}
			if a[i] == right {
				if stack.Len() != 0 {
					tmp := stack.Front()
					if tmp.Value.(byte) == left {
						res.PushBack(right)
						stack.Remove(tmp)
					} else {
						res.PushBack(left)
						stack.PushFront(left)
					}
				} else {
					res.PushBack(right)
					stack.PushFront(right)
				}
			}
		}
		for stack.Len() != 0 {
			tmp := stack.Front()
			if tmp.Value.(byte) == left {
				res.PushBack(right)
				stack.Remove(tmp)
			} else {
				res.PushFront(left)
				stack.Remove(tmp)
			}
		}
		for res.Len() != 0 {
			fmt.Print(string(res.Front().Value.(byte)))
			res.Remove(res.Front())
		}
	}
}
