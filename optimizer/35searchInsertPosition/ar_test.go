package _removeDuplicatesFromSortedArray

import (
	"testing"
)
//给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
//
//你可以假设数组中无重复元素。
//
//示例 1:
//
//输入: [1,3,5,6], 5
//输出: 2
//示例 2:
//
//输入: [1,3,5,6], 2
//输出: 1
//示例 3:
//
//输入: [1,3,5,6], 7
//输出: 4
//示例 4:
//
//输入: [1,3,5,6], 0
//输出: 0

func TestHello(t *testing.T) {
	nums := []int{1,3,5,6}
	target := 3

	t.Log(searchInsert(nums, target))
}

//时间最优
//国外1
func searchInsert(nums []int, target int) int {
	for key,val := range nums {
		if val >= target {return key}
	}
	return len(nums)
}
//国外2
//国内3
func searchInsert3(nums []int, target int) int {
	if nums==nil||len(nums)<1{
		return 0
	}
	start,end:=0,len(nums)-1
	for start<=end{
		mid:=(start+end)/2
		if nums[mid]==target{
			return mid
		}else if nums[mid]<target{
			start=mid+1
		}else{
			end=mid-1
		}
	}
	return start
}
//国内4
func searchInsert4(nums []int, target int) int {
	len := len(nums)
	if len <= 0 {
		return 0
	}
	if target > nums[len-1] {
		return len
	}
	i := 0
	for i < len {
		if target > nums[i] {
			i += 1
		} else {
			break
		}
	}
	return i
}
//内存最优
//国外5
func searchInsert5(nums []int, target int) int {
	var returnIdx int
	for idx, num := range nums {
		var nextNum int
		if num == target {
			returnIdx = idx
			break
		}
		if (idx + 1) < len(nums) {
			nextNum = nums[idx+1]
		}
		if num < target && nextNum > target {
			returnIdx = idx + 1
		}
		if idx == len(nums) - 1 && target > num {
			returnIdx = len(nums)
		}
	}
	return returnIdx
}
//国外6
func searchInsert6(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			return i
		}
	}
	return len(nums)
}
//国内7
//国内8
