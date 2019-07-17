package _removeDuplicatesFromSortedArray

import (
	"testing"
)
//给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
//
//案例:
//
//s = "leetcode"
//返回 0.
//
//s = "loveleetcode",
//返回 2.

func TestHello(t *testing.T) {
	nums := "sadadsdadad"
	t.Log(firstUniqChar(nums))
}

func firstUniqChar(s string) int {
	var arr [26]int
	for _, val := range s {
		arr[val - 'a']++
	}
	for key, val := range s {
		if arr[val - 'a'] == 1 {
			return key
		}
	}
	return -1
}
