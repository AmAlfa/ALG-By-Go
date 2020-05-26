package _removeDuplicatesFromSortedArray

import (
	"testing"
)
//给定一个已按照升序排列 的有序数组，找到两个数使得它们相加之和等于目标数。
//
//函数应该返回这两个下标值 index1 和 index2，其中 index1 必须小于 index2。
//
//说明:
//
//返回的下标值（index1 和 index2）不是从零开始的。
//你可以假设每个输入只对应唯一的答案，而且你不可以重复使用相同的元素。
//示例:
//
//输入: numbers = [2, 7, 11, 15], target = 9
//输出: [1,2]
//解释: 2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。

func TestHello(t *testing.T) {
	nums := []int{2, 2, 7, 11, 15}
	target := 9
	t.Log(twoSum(nums, target))
}
//时间最优
//国外1
func twoSum1(numbers []int, target int) []int {
	m := make(map[int]int)
	for i, n := range numbers {
		if j,ok := m[target-n]; ok {
			return []int{j, i+1}
		}
		m[n]=i+1
	}
	return []int{}
}
//国外2
func twoSum2(numbers []int, target int) []int {
	for i, j := 0, len(numbers) - 1; i < j; {
		if numbers[i] + numbers[j] == target {
			return []int{i + 1, j + 1}
		} else if numbers[i] + numbers[j] > target {
			j--
		} else {
			i++
		}
	}
	return []int{}
}
//国内3
func twoSum3(nums []int, target int) []int {
	i,j := 0,len(nums)-1
	for i<j{
		a := nums[i] + nums[j]
		if  a > target {
			j--
		}else if a < target {
			i++
		}else {
			return []int{i+1,j+1}
		}
	}
	return nil
}
//国内4
func twoSum4(numbers []int, target int) []int {
	indics:=[]int{0,0}
	for i,j:=0,len(numbers)-1;i<j;{
		sum:=numbers[i]+numbers[j]
		if sum==target{
			indics=[]int{i+1,j+1}
			break
		}else if sum>target{
			j--
		}else{
			i++
		}
	}
	return indics
}
//内存最优
//国外5
func twoSum5(numbers []int, target int) []int {
	myMap := make(map[int]int)
	left, right := 0, 0

	for i := 0; i < len(numbers); i++ {
		if myMap[numbers[i]] > 0 {
			right = i + 1
			left = myMap[numbers[i]]
			break
		}
		myMap[target-numbers[i]] = i + 1
	}

	return []int{left, right}
}
//国外6
func twoSum6(nums []int, target int) []int {
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if idx, ok := m[target - nums[i]]; ok {
			return []int{idx, i + 1}
		} else {
			m[nums[i]] = i + 1
		}
	}
	return []int{}
}
//国内7
//国内8
