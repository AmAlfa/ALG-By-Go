package _removeDuplicatesFromSortedArray

import (
	"testing"
)
//给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//
//示例:
//
//输入: [0,1,0,3,12]
//输出: [1,3,12,0,0]
//说明:
//
//必须在原数组上操作，不能拷贝额外的数组。
//尽量减少操作次数。

func TestHello(t *testing.T) {
	n := []int{9,6,4,2,3,5,7,0,1}
	moveZeroes(n)
	t.Log(n)
}

func moveZeroes(nums []int) {
	i := 0
	l := len(nums)
	for _, v := range nums {
		if v != 0 {
			nums[i] = v
			i++
		}
	}
	for j := i; j < l; j++ {
		nums[j] = 0
	}
}
