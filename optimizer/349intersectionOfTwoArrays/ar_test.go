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

//时间最优
//国外1
//国外2
//国内3
//国内4
//内存最优
//国外5
func intersection(nums1 []int, nums2 []int) []int {
	numsMin := nums1
	numsMax := nums2
	if(len(numsMin) > len(numsMax)){
		numsMin = nums2
		numsMax = nums1
	}
	sort(numsMin, 0, len(numsMin)-1)
	sort(numsMax, 0, len(numsMax)-1)
	var res []int
	for i:=0;i<len(numsMin);i++ {
		if(i>0){
			if(numsMin[i] != numsMin[i-1] && find(numsMax, numsMin[i])){
				res = append(res, numsMin[i])
			}
		} else if(find(numsMax, numsMin[i])){
			res = append(res, numsMin[i])
		}
	}
	return res
}

func sort(nums []int, left int, right int) {
	if(left >= right){
		return
	}
	tmpBase := left
	tmpLeft := left+1
	tmpRight := right
	base := nums[tmpBase]
	for ;tmpLeft<=tmpRight; {
		for ;tmpLeft<=tmpRight; {
			if(nums[tmpRight] >= base){
				tmpRight--
			} else {
				nums[tmpBase] = nums[tmpRight]
				nums[tmpRight] = base
				tmpBase = tmpRight
				tmpRight--
				break
			}
		}
		for ;tmpLeft<=tmpRight; {
			if(nums[tmpLeft] < base){
				tmpLeft++
			} else {
				nums[tmpBase] = nums[tmpLeft]
				nums[tmpLeft] = base
				tmpBase = tmpLeft
				tmpLeft++
				break
			}
		}

	}
	sort(nums, left, tmpBase-1)
	sort(nums, tmpBase+1, right)
}

func find(nums []int , target int) bool{
	left := 0
	right := len(nums) - 1
	for ;left<=right; {
		mid := left+(right-left)/2
		if(nums[mid] == target){
			return true
		} else if(nums[mid] < target) {
			left = mid+1
		} else {
			right = mid-1
		}
	}
	return false
}
//国外6
func intersection(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	ans := []int{}

	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			if len(ans) < 1 || nums1[i] != ans[len(ans) - 1] {
				ans = append(ans, nums1[i])
			}
			i++
			j++
			continue
		}
		if nums1[i] > nums2[j] {
			j++
		} else {
			i++
		}
	}

	return ans
}
//国内7
//国内8
