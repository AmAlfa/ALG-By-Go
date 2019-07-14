package _removeDuplicatesFromSortedArray

import (
	"testing"
)
//给定一个大小为 n 的数组，找到其中的众数。众数是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
//
//你可以假设数组是非空的，并且给定的数组总是存在众数。
//
//示例 1:
//
//输入: [3,2,3]
//输出: 3
//示例 2:
//
//输入: [2,2,1,1,1,2,2]
//输出: 2

func TestHello(t *testing.T) {
	nums := []int{1}
	t.Log(majorityElement(nums))
}

func majorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	numsLen := len(nums) / 2
	m := make(map[int]int)
	for _, val := range nums {
		//t.Log(m)
		if _, ok := m[val]; ok {
			m[val] += 1
			if m[val] > numsLen {
				return val
			}
		} else {
			m[val] = 1
		}
	}
	return 0
}
