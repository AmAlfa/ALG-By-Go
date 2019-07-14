package _removeDuplicatesFromSortedArray

import (
	"sort"
	"testing"
)
//给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
//
//说明：
//
//你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
//
//示例 1:
//
//输入: [2,2,1]
//输出: 1
//示例 2:
//
//输入: [4,1,2,1,2]
//输出: 4

func TestHello(t *testing.T) {
	nums := []int{9, 9}
	t.Log(singleNumber(nums))
}

//时间最优
//国外1
func singleNumber1(nums []int) int {
	mask := 0
	for _, num := range nums {
		mask = mask ^ num
	}
	return mask
}
//国外2
func singleNumber2(nums []int) int {
	var res int
	for _, num := range nums {
		res = res ^ num
	}
	return res
}
//国内3
func singleNumber3(nums []int) int {
	var t int
	for k, num := range nums {
		if k == 0 {
			t = num
		} else {
			t = t ^ num
		}
	}

	return t
}
//国内4
func singleNumber4(nums []int) int {
	var res int
	for _, v := range nums {
		res ^= v
	}
	return res
}
//内存最优
//国外5
func singleNumber5(nums []int) int {
	sort.Ints(nums)
	var ret int
	var i int
	for i=0;i<len(nums)-1;i=i+2{
		if nums[i] != nums[i+1] {
			ret = nums[i]
			break
		}
	}
	if i == len(nums)-1{
		ret = nums[i]
	}
	return ret
}
//国外6
func singleNumber6(nums []int) int {
	var ret = nums[0]
	for i:=1;i<len(nums);i++{
		ret = ret ^ nums[i]
	}
	return ret
}
//国内7
//国内8
