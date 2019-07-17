package _removeDuplicatesFromSortedArray

import (
	"testing"
)
//给定两个字符串 s 和 t，它们只包含小写字母。
//
//字符串 t 由字符串 s 随机重排，然后在随机位置添加一个字母。
//
//请找出在 t 中被添加的字母。
//
// 
//
//示例:
//
//输入：
//s = "abcd"
//t = "abcde"
//
//输出：
//e
//
//解释：
//'e' 是那个被添加的字母。


func TestHello(t *testing.T) {
	nums := "sadadsdadad"
	nums2 := "sadadsdvdad"
	t.Log(findTheDifference(nums, nums2))
}

//时间最优
//国外1
//国外2
//国内3
//国内4
//内存最优
//国外5
//国外6
//国内7
//国内8
