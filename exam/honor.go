package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	sn := scan.Text()
	n, _ := strconv.Atoi(sn)
	sarr := make([]string, 0)
	rarr := make([]string, 0)
	for i := 0; i < n; i++ {
		scan.Scan()
		sn = scan.Text()
		sarr = append(sarr, sn)
	}
	for i := 0; i < n; i++ {
		scan.Scan()
		sn = scan.Text()
		rarr = append(rarr, sn)
	}
	res := 0
	for i, pre := range sarr {
		cur := rarr[i]
		min := 1<<(32-1) - 1
		tmp := 0
		index := 0
		for j := 0; j < len(pre); j++ {
			for p := 0; p < len(cur) && p < len(pre)-j; p++ {
				if cur[p] == pre[p+j] {
					continue
				}
				index = p
				tmp++
			}
			if len(cur) > len(pre) {
				tmp += len(cur) - len(pre)
			} else {
				tmp += len(pre) - len(cur)
			}
			if len(cur) == len(pre) {
				tmp += j
				tmp += len(cur) - index
			}
			if tmp < min {
				min = tmp
			}
			tmp = 0
			fmt.Println(index, min)
		}
		res += min
	}
	fmt.Print(res)
}

func kmp(a string, b string) (int, int) {
	index, lenth := 0, 0
	for i := 1; i < len(b); i++ {
		for j := 0; j < len(b)-i; j++ {
			sub := strings.Index(a, b[j:j+i])
			if sub >= 0 && i > lenth {
				slen := strings.Count(a, b[j:j+1]) * i
				if slen > i {
					lenth = slen / i
				}
				lenth = i
				index = sub
			}
		}
	}
	return index, lenth
}
