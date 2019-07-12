package two_sum

import (
	"testing"
)

//func twoSum(nums []int, target int) []int {
//	m := make(map[int]int)
//	for i := 0; i < len(nums); i++ {
//		if v, ok := m[target - nums[i]]; ok && v != i{
//			return []int{v, i}
//		}
//		m[nums[i]] = i
//	}
//	return []int{}
//}

func twoSum(nums []int, target int) []int {
	count := len(nums) - 1//2
	for i := 0; i <= count; i++ {
		//if nums[i] <= target {
			for j := i + 1; j <= count; j++ {
				if nums[i] + nums[j] == target {
					return []int{i, j}
				}
			}
		//}
	}
	return []int{0, 0}
}

func TestTwoSum(t *testing.T)  {
	arr := []int{2, 7, 1, 10}
	t.Logf("%v\n", twoSum(arr, 17))
}
