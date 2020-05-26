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
//时间最优
//国外1
func findUnsortedSubarray1(nums []int) int {

	left, right := 0, -1
	max, min := nums[0], nums[len(nums)-1]

	for i := range nums {
		if max <= nums[i] {
			max = nums[i]
		} else {
			right = i
		}

		j := len(nums) - i - 1
		if min >= nums[j] {
			min = nums[j]
		} else {
			left = j
		}
	}

	return right - left + 1
}
//国外2
func findUnsortedSubarray2(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	var left int
	for i:=1;i<len(nums);i++ {
		if nums[i] >=nums[i-1] {
			left = i
		}else{
			break
		}
	}

	if left == len(nums)-1 {
		return 0
	}

	for j:=left;j<len(nums)&&left>=0;j++ {
		for left>=0&&nums[j] < nums[left] {
			left--
		}
	}

	var right int = len(nums)-1
	for i:=len(nums)-2;i>=0;i-- {
		if nums[i] <= nums[i+1]{
			right =i
		}else{
			break
		}
	}

	if right == 0 {
		return 0
	}

	for j:=right;j>=0&&right<=len(nums)-1;j--{
		for right<=len(nums)-1 &&nums[j] >nums[right]{
			right ++
		}
	}

	return right-left-1
}
//国内3
func findUnsortedSubarray3(nums []int) int {
	length := len(nums)
	left, right, min, max := -1, -2, nums[length-1], nums[0]
	for i := 1; i < length; i++ {
		if max < nums[i] {
			max = nums[i]
		}
		if min > nums[length-i-1] {
			min = nums[length-i-1]
		}
		if nums[i] < max {
			right = i
		}
		if nums[length-i-1] > min {
			left = length-i-1
		}
	}
	return right - left + 1
}

//国内4
func findUnsortedSubarray4(nums []int) int {
	length := len(nums)
	if length < 2{
		return 0
	}
	m := nums[0]
	n:=nums[length-1]
	l := -1
	r := -2
	for i:=1;i<length;i++{
		m = max(m,nums[i])
		n = min(n,nums[length-i-1])
		if m!=nums[i]{
			r = i
		}
		if n != nums[length-i-1]{
			l = length-i-1
		}
	}

	return r-l+1
}

func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}


func min(a,b int) int{
	if a<b {
		return a
	}
	return b
}
//内存最优
//国外5
func findUnsortedSubarray5(nums []int) int {
	n := len(nums)
	sorted := make ([]int, len(nums))
	copy(sorted, nums)
	sort.Ints(sorted)
	var i int
	for i = 0; i < n; i++ {
		if sorted[i] != nums[i] {
			break
		}
	}
	if i == n {
		return 0
	}
	start := i
	for i = n-1; i > start; i-- {
		if sorted[i] != nums[i] {
			break
		}
	}
	end := i
	return end - start + 1
}
//国外6
func findUnsortedSubarray6(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	min := len(nums)
	max := 0
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			for j := 0; j < i; j++ {
				if nums[j] > nums[i] {
					if j < min {
						min = j
					}
					max = i - 1
					nums[i], nums[j] = nums[j], nums[i]
				}
			}
		}
	}
	max++

	fmt.Println(nums, min, max)

	if min >= max {
		return 0
	}

	return len(nums[min:max]) + 1
}
//国内7
//国内8
