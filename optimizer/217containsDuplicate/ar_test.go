package _removeDuplicatesFromSortedArray

import (
	"sort"
	"testing"
)
//给定一个整数数组，判断是否存在重复元素。
//
//如果任何值在数组中出现至少两次，函数返回 true。如果数组中每个元素都不相同，则返回 false。
//
//示例 1:
//
//输入: [1,2,3,1]
//输出: true
//示例 2:
//
//输入: [1,2,3,4]
//输出: false
//示例 3:
//
//输入: [1,1,1,3,3,4,3,2,4,2]
//输出: true

func TestHello(t *testing.T) {
	nums := []int{1,1,1,3,3,4,3,2,4,2}
	t.Log(containsDuplicate(nums))
}

//时间最优
//国外1
func containsDuplicate1(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}
	list := make(map[int]int)
	for i,v := range nums{
		if _,ok := list[v]; ok{
			return true
		}else{
			list[v] = i
		}
	}
	return false
}
//国外2
//国内3
func containsDuplicate3(nums []int) bool {
	freq := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if freq[nums[i]] == 1 {
			return true
		}
		freq[nums[i]] = 1
	}

	return false
}
//国内4
func containsDuplicate4(nums []int) bool {
	if len(nums)<=1{
		return false
	}
	m:=make(map[int]bool)
	for i:=0;i<len(nums);i++{
		if _,ok:=m[nums[i]];ok{
			return true
		}else {
			m[nums[i]]=true
		}
	}
	return false
}
//内存最优
//国外5
func containsDuplicate5(nums []int) bool {
	sort.Ints(nums)
	for i, n := range nums {
		if i > 0 && n == nums[i-1] {
			return true
		}
	}
	return false
}
//国外6
func containsDuplicate6(nums []int) bool {
	if len(nums) == 1 {
		return false
	}
	for i := range nums {
		if i != last_index_of(nums, nums[i]) {
			return true
		}
	}
	return false
}

func last_index_of(nums []int, val int) int {
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == val {
			return i
		}
	}
	return -1
}
//国内7
//国内8
