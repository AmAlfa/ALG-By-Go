package _removeDuplicatesFromSortedArray

import (
	"strconv"
	"strings"
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

func isPowerOfThree(n int) bool {
	if n < 1 {
		return false
	}
	s := strconv.FormatInt(int64(n), 3)
	return s[:1] == "1" && strings.Count(s, "0") == len(s) - 1
}
