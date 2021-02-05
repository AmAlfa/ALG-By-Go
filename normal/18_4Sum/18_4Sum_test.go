package romanToInteger

import (
	"fmt"
	"sort"
	"testing"
)

// 给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。
//
// 注意：
//
// 答案中不可以包含重复的四元组。
//
// 示例：
//
// 给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。
//
// 满足要求的四元组集合为：
// [
//  [-1,  0, 0, 1],
//  [-2, -1, 1, 2],
//  [-2,  0, 0, 2]
// ]



func Test(t *testing.T) {
	s := []int{1, 0, -1, 0, -2, 2}
	target := 0
	res := fourSum(s, target)
	fmt.Println(res)
	t.Log(res)

}

// 排序+双指针
func fourSum(nums []int, target int) [][]int {
	var result [][]int
	numsLen := len(nums)
	// 判断nums满足4元素
	if numsLen < 4 {
		return result
	}

	// 判断第一个元素是否大于target
	sort.Ints(nums)
	reLasSub := numsLen - 1
	reSecSub := numsLen - 2
	reThrSub := numsLen - 3
	reLastNum := nums[numsLen - 1]
	reSecNum := nums[reSecSub]
	reThrNum := nums[reThrSub]
	sumLS := reLastNum + reSecNum
	sumLST := reLastNum + reSecNum + reThrNum
	for i := 0; i < reThrSub; i++  {
		// 去重
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}

		// 判断最小的值大于目标值
		min := nums[i] + nums[i + 1] + nums[i + 2] + nums[i + 3]
		if min > target {
			break
		}

		// 判断最大的值小于目标值
		max := nums[i] + sumLST
		if max < target {
			continue
		}

		for j := i + 1; j < reSecSub; j++ {
			if j > i + 1 && nums[j] == nums[j - 1] {
				continue
			}

			// 判断最小的值大于目标值
			min := nums[i] + nums[j] + nums[j + 1] + nums[j + 2]
			if min > target {
				break
			}

			// 判断最大的值小于目标值
			max := nums[i] + nums[j] + sumLS
			if max < target {
				continue
			}

			l := j + 1
			r := reLasSub
			for l < r {
				surplus := target - nums[i] - nums[j] - nums[l] - nums[r]
				if surplus == 0 {
					result = append(result, []int{nums[i], nums[j], nums[l], nums[r]})
					// 破坏平衡，进行下一个比对
					cur := l
					l++
					for ; l < r && nums[l] == nums[cur]; l++ {
					}
				} else if surplus > 0 {
					cur := l
					l++
					for ; l < r && nums[l] == nums[cur]; l++ {
					}
				} else {
					cur := r
					r--
					for ; l < r && nums[r] == nums[cur]; r-- {
					}
				}
			}

		}
	}

	return result
}
