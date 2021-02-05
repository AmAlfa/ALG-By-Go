package merge_two_sorted_lists_test

import (
	"fmt"
	"testing"
)
// 实现 strStr() 函数。
// 给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。
// 示例 1:
// 输入: haystack = "hello", needle = "ll"
// 输出: 2
//
// 示例 2:
// 输入: haystack = "aaaaa", needle = "bba"
// 输出: -1
// 说明:
// 当 needle 是空字符串时我们应当返回 0 。

func Test(t *testing.T) {
	s1 := ""
	s2 := "a"
	res := strStr(s1, s2)

	fmt.Println(res)
	t.Log(res)

}

func strStr(haystack string, needle string) int {
	lh, ln := len(haystack), len(needle)
	for i := 0; i < lh-ln+1; i++ {
		if haystack[i:i+ln] == needle {
			return i
		}
	}
	return -1
}

// func strStr(haystack string, needle string) int {
// 	haystackLength := len(haystack) - 1
// 	needleLength := len(needle) - 1
// 	if needleLength == -1 {
// 		return 0
// 	}
// 	for k, v := range haystack {
// 		if string(v) != string(needle[0]) {
// 			continue
// 		}
//
// 		tmp := k + needleLength
// 		if tmp > haystackLength {
// 			return -1
// 		}
// 		if haystack[k : tmp + 1] == needle {
// 			return k
// 		}
// 	}
//
// 	return -1
// }
