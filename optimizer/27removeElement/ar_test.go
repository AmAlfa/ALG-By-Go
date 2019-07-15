package _removeDuplicatesFromSortedArray

import (
	"sort"
	"testing"
)
//给定一个数组 nums 和一个值 val，你需要原地移除所有数值等于 val 的元素，返回移除后数组的新长度。
//
//不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
//
//元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
//
//示例 1:
//
//给定 nums = [3,2,2,3], val = 3,
//
//函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。
//
//你不需要考虑数组中超出新长度后面的元素。
//示例 2:
//
//给定 nums = [0,1,2,2,3,0,4,2], val = 2,
//
//函数应该返回新的长度 5, 并且 nums 中的前五个元素为 0, 1, 3, 0, 4。
//
//注意这五个元素可为任意顺序。
//
//你不需要考虑数组中超出新长度后面的元素。
//说明:
//
//为什么返回数值是整数，但输出的答案是数组呢?
//
//请注意，输入数组是以“引用”方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。
//
//你可以想象内部操作如下:
//
//// nums 是以“引用”方式传递的。也就是说，不对实参作任何拷贝
//int len = removeElement(nums, val);
//
//// 在函数里修改输入数组对于调用者是可见的。
//// 根据你的函数返回的长度, 它会打印出数组中该长度范围内的所有元素。
//for (int i = 0; i < len; i++) {
//    print(nums[i]);
//}

func TestHello(t *testing.T) {
	nums := []int{0}
	val := 3
	t.Log(removeElement(nums, val))
}
//时间最优
//国外1
//国外2
//国内3
func removeElement3(nums []int, val int) int {
	// 从0开始依次放不等于目标值的数
	pos:=0
	for i:=range(nums){
		if nums[i]!=val{
			nums[pos]=nums[i]
			pos++
		}
	}
	return pos
}
//国内4
//内存最优
//国外5
func removeElement5(nums []int, val int) int {
	sort.Ints(nums)
	result := searchRange(nums, val)
	if result[0] != -1 {
		nums = append(nums[:result[0]], nums[result[1]+1:]...)
	}

	return len(nums)
}

func searchRange(nums []int, target int) []int {
	lenNums := len(nums)

	if lenNums == 0 {
		return []int{-1, -1}
	}

	if lenNums == 1 {
		if nums[0] == target {
			return []int{0, 0}
		}
		return []int{-1, -1}
	}

	if target > nums[lenNums-1] {
		return []int{-1, -1}
	}
	if target < nums[0] {
		return []int{-1, -1}
	}

	resultStart := findTheFirst(nums, target)
	if resultStart != -1 {
		resultEnd := findTheLast(nums, target)
		return []int{resultStart, resultEnd}
	} else {
		return []int{-1, -1}
	}

	return []int{-1, -1}
}

func findTheFirst(nums []int, target int) int {
	lenNums := len(nums)

	start := 0
	end := lenNums - 1

	for start < end-1 {
		if nums[start] == target && (start == 0 || nums[start-1] != target) {
			return start
		}

		middle := (end-start+1)/2 + start
		if target > nums[middle] {
			start = middle
		} else {
			end = middle
		}
	}

	if nums[start] == target && (start == 0 || nums[start-1] != target) {
		return start
	}
	if nums[end] == target && (end == 0 || nums[end-1] != target) {
		return end
	}

	return -1
}

func findTheLast(nums []int, target int) int {
	lenNums := len(nums)

	start := 0
	end := lenNums - 1

	for start < end-1 {
		if nums[end] == target && (end == lenNums-1 || nums[end+1] != target) {
			return end
		}

		middle := (end-start+1)/2 + start
		if target < nums[middle] {
			end = middle
		} else {
			start = middle
		}
	}

	if nums[end] == target && (end == lenNums-1 || nums[end+1] != target) {
		return end
	}
	if nums[start] == target && (start == lenNums-1 || nums[start+1] != target) {
		return start
	}

	return -1
}
//国外6
func removeElement(nums []int, val int) int {
	sort.Ints(nums)
	result := searchRange(nums, val)
	if result[0] != -1 {
		nums = append(nums[:result[0]], nums[result[1]+1:]...)
	}

	return len(nums)
}

func searchRange(nums []int, target int) []int {
	lenNums := len(nums)

	if lenNums == 0 {
		return []int{-1, -1}
	}

	if lenNums == 1 {
		if nums[0] == target {
			return []int{0, 0}
		}
		return []int{-1, -1}
	}

	if target > nums[lenNums-1] {
		return []int{-1, -1}
	}
	if target < nums[0] {
		return []int{-1, -1}
	}

	resultStart := findTheFirst(nums, target)
	if resultStart != -1 {
		resultEnd := findTheLast(nums, target)
		return []int{resultStart, resultEnd}
	} else {
		return []int{-1, -1}
	}

	return []int{-1, -1}
}

func findTheFirst(nums []int, target int) int {
	lenNums := len(nums)

	start := 0
	end := lenNums - 1

	for start < end-1 {
		if nums[start] == target && (start == 0 || nums[start-1] != target) {
			return start
		}

		middle := (end-start+1)/2 + start
		if target > nums[middle] {
			start = middle
		} else {
			end = middle
		}
	}

	if nums[start] == target && (start == 0 || nums[start-1] != target) {
		return start
	}
	if nums[end] == target && (end == 0 || nums[end-1] != target) {
		return end
	}

	return -1
}

func findTheLast6(nums []int, target int) int {
	lenNums := len(nums)

	start := 0
	end := lenNums - 1

	for start < end-1 {
		if nums[end] == target && (end == lenNums-1 || nums[end+1] != target) {
			return end
		}

		middle := (end-start+1)/2 + start
		if target < nums[middle] {
			end = middle
		} else {
			start = middle
		}
	}

	if nums[end] == target && (end == lenNums-1 || nums[end+1] != target) {
		return end
	}
	if nums[start] == target && (start == lenNums-1 || nums[start+1] != target) {
		return start
	}

	return -1
}
//国内7
//国内8
