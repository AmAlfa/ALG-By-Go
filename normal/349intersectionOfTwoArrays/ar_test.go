package _removeDuplicatesFromSortedArray

import (
	"testing"
)
//给定两个数组，编写一个函数来计算它们的交集。
//
//示例 1:
//
//输入: nums1 = [1,2,2,1], nums2 = [2,2]
//输出: [2]
//示例 2:
//
//输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//输出: [9,4]
//说明:
//
//输出结果中的每个元素一定是唯一的。
//我们可以不考虑输出结果的顺序。


func TestHello(t *testing.T) {
	nums1 := []int{4,9,5}
	nums2 := []int{9,4,9,8,4}
	t.Log(intersection(nums1, nums2))
}

func intersection(nums1 []int, nums2 []int) []int {
	nums1Tmp := make(map[int]bool)
	result := make([]int, 0)
	for _, val := range nums1 {
		nums1Tmp[val] = true
	}
	for _, val := range nums2 {
		if m, ok := nums1Tmp[val]; ok && m {
			result = append(result, val)
			nums1Tmp[val] = false
		}
	}
	return result
}
