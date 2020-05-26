package _removeDuplicatesFromSortedArray

import (
	"strconv"
	"testing"
)
//给定一个整数，将其转化为7进制，并以字符串形式输出。
//
//示例 1:
//
//输入: 100
//输出: "202"
//示例 2:
//
//输入: -7
//输出: "-10"
//注意: 输入范围是 [-1e7, 1e7] 。

func TestHello(t *testing.T) {
	a := 5
	t.Log(convertToBase7(a))
}

func convertToBase7(num int) string {
	return strconv.FormatInt(int64(num), 7)
}
