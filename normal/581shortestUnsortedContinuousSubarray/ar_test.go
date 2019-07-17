package _removeDuplicatesFromSortedArray

import (
	"sort"
	"testing"
)
//给定一个整数数组，你需要寻找一个连续的子数组，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。
//
//你找到的子数组应是最短的，请输出它的长度。
//
//示例 1:
//
//输入: [2, 6, 4, 8, 10, 9, 15]
//输出: 5
//解释: 你只需要对 [6, 4, 8, 10, 9] 进行升序排序，那么整个表都会变为升序排序。
//说明 :
//
//输入的数组长度范围在 [1, 10,000]。
//输入的数组可能包含重复元素 ，所以升序的意思是<=。



func TestHello(t *testing.T) {
	s := []int{3,1,25,36,44,321,21}

	t.Log(findUnsortedSubarray(s))
}
func findUnsortedSubarray(nums []int) int {
	tmp := make([]int, len(nums))
	copy(tmp, nums)
	sort.Ints(tmp)
	i, j := 0, len(nums) - 1
	for i < j && nums[i] == tmp[i] {
		i++
	}
	for i <= j && nums[j] == tmp[j] {
		j--
	}
	return j - i + 1
}
