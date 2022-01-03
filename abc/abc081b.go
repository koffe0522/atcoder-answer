package abc

import "fmt"

func countDivisor(nums []int, count int) int {
	isFinish := false
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 != 0 {
			isFinish = true
			break
		}

		nums[i] = nums[i] / 2
	}

	if isFinish {
		return count
	}

	return countDivisor(nums, count+1)
}

// Abc081b is `ABC081B - Shift only` answer
// URL: https://atcoder.jp/contests/abs/tasks/abc081_b
func Abc081b() {
	// N は可変調
	var n int
	fmt.Scanf("%d", &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}

	fmt.Println(countDivisor(a, 0))
}
