package _removeDuplicatesFromSortedArray

import (
	"testing"
)
//给定一个整数，写一个函数来判断它是否是 3 的幂次方。
//
//示例 1:
//
//输入: 27
//输出: true
//示例 2:
//
//输入: 0
//输出: false
//示例 3:
//
//输入: 9
//输出: true
//示例 4:
//
//输入: 45
//输出: false
//进阶：
//你能不使用循环或者递归来完成本题吗？
func TestHello(t *testing.T) {
	n := 102
	t.Log(isPowerOfThree(n))
}

//时间最优
//国外1
func isPowerOfThree1(n int) bool {
	if n < 1 {
		return false
	}

	for n % 3 == 0 {
		n /= 3
	}

	return n == 1
}
//国外2
//国内3
//国内4
//内存最优
//国外5
//国外6
//国内7
//国内8
