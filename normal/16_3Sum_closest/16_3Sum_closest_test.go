package romanToInteger

import (
	"fmt"
	"sort"
	"testing"
)

// 给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。
//
// 例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.
//
// 与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).



func Test(t *testing.T) {
	s := []int{-1, 0, 1, 2, -1, -4}
	target := 9
	res := threeSumClosest(s, target)
	fmt.Println(res)
	t.Log(res)

}

func threeSumClosest(nums []int, target int) int {

	if len(nums)==0 {
		return 0
	}
	if len(nums)==1 {
		return nums[0]
	}
	if len(nums)==2 {
		return nums[0] + nums[1]
	}
	if len(nums)==3 {
		return nums[0] + nums[1] + nums[2]
	}

	sort.Ints(nums)

	res := nums[0] + nums[1] + nums[2]

	for i:=0; i<len(nums); i++ {
		// 避免添加重复的结果
		// i>0 是为了防止nums[i-1]溢出
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1

		for l < r {
			s := nums[i] + nums[l] + nums[r]

			switch {
			case s < target:
				// 较小的 l 需要变大
				var d int
				if res < target {
					d = target - res
				} else {
					d = res - target
				}
				if d > target -s {
					res = s
				}

				l++
			case s > target:
				// 较大的 r 需要变小
				var d int
				if res < target {
					d = target - res
				} else {
					d = res - target
				}
				if d> s - target {
					res = s
				}
				r--
			default:
				return target
			}

		}
	}

	return res
}

//
// func threeSumClosest(nums []int, target int) int {
// 	sort.Ints(nums)
// 	middle, length := 1, len(nums) - 1
// 	sub := absolute(nums[0] + nums[1] + nums[2] - target)
// 	result := nums[0] + nums[1] + nums[2]
// 	for ; middle < length; middle++ {
// 		left, right :=  0, length
// 		for middle > left && middle < right {
// 			sum := nums[left] + nums[middle] + nums[right]
// 			tmp := absolute(sum - target)
// 			if tmp <= sub {
// 				sub = tmp
// 				result = sum
// 			}
// 			if sum < target {
// 				left ++
// 			} else if sum > target {
// 				right--
// 			} else {
// 				break
// 			}
// 		}
// 	}
//
// 	return result
// }
//
// func absolute(num int) (result int) {
// 	if num < 0 {
// 		result = -num
// 	} else {
// 		result = num
// 	}
//
// 	return
// }
