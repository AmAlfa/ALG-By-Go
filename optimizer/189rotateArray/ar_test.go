package _removeDuplicatesFromSortedArray

import (
	"testing"
)
//给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
//
//示例 1:
//
//输入: [1,2,3,4,5,6,7] 和 k = 3
//输出: [5,6,7,1,2,3,4]
//解释:
//向右旋转 1 步: [7,1,2,3,4,5,6]
//向右旋转 2 步: [6,7,1,2,3,4,5]
//向右旋转 3 步: [5,6,7,1,2,3,4]
//示例 2:
//
//输入: [-1,-100,3,99] 和 k = 2
//输出: [3,99,-1,-100]
//解释:
//向右旋转 1 步: [99,-1,-100,3]
//向右旋转 2 步: [3,99,-1,-100]
//说明:
//
//尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
//要求使用空间复杂度为 O(1) 的 原地 算法。

func TestHello(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	//k := 3
	//rotate(nums, k)
	t.Log(nums)
}

//时间最优
//国外1
func rotate(nums []int, k int)  {
	k = k % len(nums)
	temp := append(nums[len(nums)-k:],nums[:len(nums)-k]...)
	copy(nums,temp)
	//fmt.Println(nums)

}
//国外2
//国内3
func rotate3(nums []int, k int)  {
	k = k % len(nums)
	reverse(nums[:len(nums)-k])
	reverse(nums[len(nums)-k:])
	reverse(nums)
}

func reverse(nums []int) {
	head,tail := 0,len(nums)-1
	for head < tail {
		nums[head],nums[tail] = nums[tail],nums[head]
		head++
		tail--
	}
}
//国内4
func rotate4(nums []int, k int) {
	len1 := len(nums)
	k = k % len1
	nums1 := make([]int, k)
	copy(nums1, nums[len1-k:])
	for i := len1 - k - 1; i >= 0; i-- {
		nums[i+k] = nums[i]
	}
	for j := 0; j < len(nums1); j++ {
		nums[j] = nums1[j]
	}
}
//内存最优
//国外5
//国外6
//国内7
//国内8
