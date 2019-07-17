package _removeDuplicatesFromSortedArray

import (
	"testing"
)
//统计所有小于非负整数 n 的质数的数量。
//
//示例:
//
//输入: 10
//输出: 4
//解释: 小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。


func TestHello(t *testing.T) {
	n := 10
	t.Log(countPrimes(n))
}

func countPrimes(n int) int {
	a := make([]bool, n)
	if n == 0 || n == 1 {
		return 1
	}
	count := 0
	for i := 2; i < n; i++ {
		if a[i] {
			continue
		}
		for j := 1; j * i < n; j++ {
			a[j * i] = true
		}
		count++
	}
	return count
}
