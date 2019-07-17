package _removeDuplicatesFromSortedArray

import (
	"testing"
)
//给定一个包含 0, 1, 2, ..., n 中 n 个数的序列，找出 0 .. n 中没有出现在序列中的那个数。
//
//示例 1:
//
//输入: [3,0,1]
//输出: 2
//示例 2:
//
//输入: [9,6,4,2,3,5,7,0,1]
//输出: 8
//说明:
//你的算法应具有线性时间复杂度。你能否仅使用额外常数空间来实现?

func TestHello(t *testing.T) {
	n := []int{9,6,4,2,3,5,7,0,1}
	t.Log(missingNumber(n))
}

//时间最优
//国外1
func missingNumber1(nums []int) int {
	var data int

	for i := 0; i <= len(nums); i++ {
		if i == 0 {
			data = i
		} else {
			data = data ^ i
		}
	}

	for _, v :=range nums {
		data = data ^v
	}
	return data
}
//国外2
func missingNumber2(nums []int) int {
	n := len(nums)

	sum := 0
	for i:=0;i<n;i++ {
		sum += nums[i]
	}
	t := 0
	if n%2==0 {
		t = (n+1)*(n/2)
	}else{
		t = (n+1)*(n/2)+(n+1)/2
	}

	return t-sum
}
//国内3
func missingNumber3(nums []int) int {
	if len(nums)==0 {
		return -1
	}
	missingnum:=len(nums)
	for i:=0;i<len(nums);i++ {
		missingnum^=i^nums[i]
	}
	return missingnum
}
//国内4
//内存最优
//国外5
//国外6
//国内7
//国内8
