package _removeDuplicatesFromSortedArray

import (
	"testing"
)
//给定一个大小为 n 的数组，找到其中的众数。众数是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
//
//你可以假设数组是非空的，并且给定的数组总是存在众数。
//
//示例 1:
//
//输入: [3,2,3]
//输出: 3
//示例 2:
//
//输入: [2,2,1,1,1,2,2]
//输出: 2

func TestHello(t *testing.T) {
	nums := []int{1}
	t.Log(majorityElement(nums))
}
//时间最优
//国外1
func majorityElement1(nums []int) int {
	res := nums[0]
	count := 0
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			res = nums[i]
			count = 1
		} else {
			if nums[i] == res {
				count++
			} else {
				count--
			}
		}
	}

	return res
}
//国外2
func majorityElement2(nums []int) int {
	ht := make(map[int]int)
	mark := len(nums) / 2
	for _, v := range nums {
		if count, ok := ht[v]; !ok {
			ht[v] = 1
		}else {
			ht[v]++
			if count + 1 > mark {
				return v
			}
		}
	}

	return nums[0]
}
//国内3
func majorityElement3(nums []int) int {
	maps := make(map[int]int, len(nums))

	limit := len(nums) / 2

	for _,v := range nums {
		if _,ok := maps[v]; !ok {
			maps[v] = 1
		} else {
			maps[v]++
		}

		if maps[v] > limit {
			return v
		}
	}

	return -1
}
//国内4
func majorityElement4(nums []int) int {
	threshold := len(nums) / 2

	numMap := make(map[int]int)
	for _, num := range nums {
		numMap[num]++
		if numMap[num] > threshold {
			return num
		}
	}

	return -1
}
//内存最优
//国外5
func majorityElement5(nums []int) int {
	sm := make(map[int]int)
	for _, num := range nums {
		if val, ok := sm[num]; ok {
			sm[num] = val + 1
		} else {
			sm[num] = 1
		}
	}
	for key, val := range sm {
		if val > len(nums)/2 {
			return key
		}
	}
	return 0
}
//国外6
func majorityElement6(nums []int) int {
	iMap := make(map[int]int)

	for i := 0; i < len(nums); i ++ {
		if val, ok := iMap[nums[i]]; ok {
			val += 1
			iMap[nums[i]] = val
		} else {
			iMap[nums[i]] = 1
		}
	}

	max := 0
	num := 0
	for k, v := range iMap {
		if v > max {
			max = v
			num = k
		}
	}

	return num
}
//国内7
//国内8
