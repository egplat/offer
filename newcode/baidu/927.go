package baidu

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main1() {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	s := scan.Text()
	sa := strings.Split(s, " ")
	n, _ := strconv.Atoi(sa[0])
	t, _ := strconv.Atoi(sa[1])
	c, _ := strconv.Atoi(sa[2])

	scan.Scan()
	s = scan.Text()
	sa = strings.Split(s, " ")
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		num, _ := strconv.Atoi(sa[i])
		nums[i] = num
	}

	record := 0
	res := 0
	start := 0
	for i := 0; i < n; i++ {
		record += nums[i]
		for record > t {
			record -= nums[start]
			start++
		}
		if i-start+1 == c {
			record -= nums[start]
			start++
			res++
		}
	}
	fmt.Print(res)
}

func main3() {
	var n, m, x, y, t int
	for {
		if _, err := fmt.Scan(&n, &m, &x, &y, &t); err != nil {
			break
		}
		p := make([]float64, n*m)
		var pcc float64
		for i := 0; i < n; i++ {
			var tmp []float64
			for j := 0; j < m; j++ {
				var k float64
				fmt.Scan(&k)
				tmp = append(tmp, k)
				if i == x-1 && j == y-1 {
					pcc = k
				}
			}
			p = append(p, tmp...)
		}
		var total float64
		for i := 0; i < len(p); i++ {
			total += p[i]
		}
		average := total / (float64(n * m))
		ssp := 1 - pow(1.0-average, t)
		ccp := 1 - pow(1.0-pcc, t)
		if ssp > ccp {
			fmt.Println("ss")
			fmt.Printf("%.2f\n", ssp)
		} else if ssp < ccp {
			fmt.Println("cc")
			fmt.Printf("%.2f\n", ccp)
		} else {
			fmt.Println("equal")
			fmt.Printf("%.2f\n", ccp)
		}
	}
}

func pow(x float64, n int) float64 {
	res := 1.0
	for i := 0; i < n; i++ {
		res *= x
	}
	return res
}

func main4() {
	var n, m, k, x, y int
	for {
		if _, err := fmt.Scan(&n, &m, &k); err != nil {
			break
		}
		if n == 1 || m == 1 {
			if k == 0 {
				fmt.Printf("%0.2f\n", 1.00)
			} else {
				fmt.Printf("%0.2f\n", 1.00)
			}
			continue
		}
		matrix, prob := make([][]float64, 0), make([][]float64, 0)
		for i := 0; i < n; i++ {
			tmp := make([]float64, m)
			matrix = append(matrix, tmp)
			prob = append(prob, tmp)
		}
		for i := 0; i < k; i++ {
			// 标记蘑菇
			fmt.Scan(&x, &y)
			matrix[x-1][y-1] = 1
		}
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if i == 0 && j == 0 {
					prob[i][j] = 1
					continue
				}
				if matrix[i][j] == 1 {
					prob[i][j] = 0
					continue
				}
				if i == n-1 && j == m-1 {
					prob[i][j] = prob[n-1][m-2] + prob[n-2][m-1]
					continue
				}
				if i == 0 {
					prob[i][j] = prob[i][j-1] * 0.5
					continue
				}
				if j == 0 {
					prob[i][j] = prob[i-1][j] * 0.5
					continue
				}
				if i == n-1 {
					prob[i][j] = prob[i][j-1] + prob[i-1][j]*0.5
					continue
				}
				if j == m-1 {
					prob[i][j] = prob[i-1][j] + prob[i][j-1]*0.5
					continue
				}
				prob[i][j] = prob[i-1][j]*0.5 + prob[i][j-1]*0.5
			}
		}
		fmt.Printf("%0.2f\n", prob[n-1][m-1])
	}
}
