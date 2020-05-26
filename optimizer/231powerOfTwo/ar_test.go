package _removeDuplicatesFromSortedArray

import (
	"strconv"
	"strings"
	"testing"
)
//给定一个整数，编写一个函数来判断它是否是 2 的幂次方。
//
//示例 1:
//
//输入: 1
//输出: true
//解释: 20 = 1
//示例 2:
//
//输入: 16
//输出: true
//解释: 24 = 16
//示例 3:
//
//输入: 218
//输出: false

func TestHello(t *testing.T) {
	n := 102
	t.Log(isPowerOfTwo(n))
}

//时间最优
//国外1
func isPowerOfTwo1(n int) bool {
	if n <= 0 {
		return false
	}
	if n&(n-1) > 0 {
		return false
	}
	return true
}
//国外2
//国内3
func isPowerOfTwo3(n int) bool {
	s := strconv.FormatInt(int64(n), 2)
	return strings.HasPrefix(s, "1") && strings.Count(s, "1") == 1
}
//国内4
//内存最优
//国外5
//国外6
//国内7
//国内8
