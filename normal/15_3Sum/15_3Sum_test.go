package romanToInteger

import (
	"fmt"
	"sort"
	"testing"
)

// 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
//
// 注意：答案中不可以包含重复的三元组。
//
//  
//
// 示例：
//
// 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
//
// 满足要求的三元组集合为：
// [
//  [-1, 0, 1],
//  [-1, -1, 2]
// ]



func Test(t *testing.T) {
	s := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(s)
	res := threeSum(s)
	fmt.Println(res)
	t.Log(res)

}

// fastest
func threeSum(nums []int) [][]int {
	if nums == nil || len(nums) < 3 {
		return nil
	}
	sort.Ints(nums)
	var ans [][]int
	length := len(nums)
	for k := 0 ; k < length - 2 ; k++ {
		if nums[k] > 0 {
			break
		}
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}

		l := k + 1
		r := length - 1
		for l < r {
			if nums[l] + nums[r] < -nums[k] {
				l++
				continue
			}else if nums[l] + nums[r] > -nums[k] {
				r--
				continue
			}else if nums[l] + nums[r] == -nums[k] {
				ans = append(ans, []int{nums[k], nums[l], nums[r]})
				for l < length-1 && nums[l] == nums[l+1] {
					l++
				}
				for r > k+1 && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			}

		}
	}
	return ans
}

// func threeSum(nums []int) [][]int {
// 	sort.Ints(nums)
// 	sort.Ints(nums)
// 	subscript := len(nums) - 1
// 	var result [][]int
// 	middle := 1
// 	for ; middle < subscript; middle ++ {
// 		left, right := 0, subscript
// 		if middle > 1 && nums[middle] == nums[middle - 1] {
// 			left = middle - 1
// 		}
//
// 		for middle > left && middle < right {
// 			sum := nums[left] + nums[middle] + nums[right]
// 			if left > 0 && nums[left] == nums[left - 1] {
// 				left++
// 				continue
// 			}
// 			if right < subscript && nums[right] == nums[right + 1] {
// 				right--
// 				continue
// 			}
// 			switch {
// 			case sum == 0:
// 				result = append(result, []int{nums[left], nums[middle], nums[right]})
// 				left++
// 				right--
// 			case sum > 0:
// 				right--
// 			case sum < 0:
// 				left++
// 			}
// 		}
// 	}
//
// 	return result
// }
