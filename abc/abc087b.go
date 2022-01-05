package abc

import "fmt"

// Abc087b is `ABC087B - Coins` answer
// URL: https://atcoder.jp/contests/abs/tasks/abc087_b
func Abc087b() {
	var a, b, c, x int
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	fmt.Scanf("%d", &c)
	fmt.Scanf("%d", &x)

	resultCount := 0
	// 0 ~ A枚まで
	for i := 0; i <= a; i++ {
		for j := 0; j <= b; j++ {
			for k := 0; k <= c; k++ {
				if (500*i)+(100*j)+(50*k) == x {
					resultCount++
				}
			}
		}
	}
	fmt.Println(resultCount)

}
