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

//时间最优
//国外1
func firstUniqChar1(s string) int {
	dic := [26]int{}
	for _, c := range s {
		dic[c - 'a']++
	}
	for i, c := range s {
		if dic[c - 'a'] == 1 {
			return i
		}
	}
	return -1
}
//国外2
//国内3
//国内4
//内存最优
//国外5
func firstUniqChar5(s string) int {
	data := make(map[string]int)
	for i := 0 ; i < len(s) ; i++ {
		if _ , ok := data[string(s[i])] ; ok{
			data[string(s[i])] = data[string(s[i])] + 1
		}else {
			data[string(s[i])] = 1
		}
	}

	for i := 0 ; i < len(s) ; i++ {
		if  data[string(s[i])] == 1 {
			return i
		}
	}

	return -1
}
//国外6
func firstUniqChar6(s string) int {
	ret := make(map[int]int, 10)
	for i := range s {
		key := int(s[i]-'a')
		if _,ok := ret[key]; ok {
			ret[key] = -1
		} else {
			ret[key] = i
		}
	}
	for i := range s {
		key := int(s[i]-'a')

		if ret[key] != -1 {
			return i
		}
	}
	return -1
}
//国内7
//国内8
