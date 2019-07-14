package _removeDuplicatesFromSortedArray

import (
	"testing"
)
//给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
//
//不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
//
//示例 1:
//
//给定数组 nums = [1,1,2],
//
//函数应该返回新的长度 2, 并且原数组 nums 的前两个元素被修改为 1, 2。
//
//你不需要考虑数组中超出新长度后面的元素。
//示例 2:
//
//给定 nums = [0,0,1,1,1,2,2,3,3,4],
//
//函数应该返回新的长度 5, 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4。
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
//// nums 是以“引用”方式传递的。也就是说，不对实参做任何拷贝
//int len = removeDuplicates(nums);
//
//// 在函数里修改输入数组对于调用者是可见的。
//// 根据你的函数返回的长度, 它会打印出数组中该长度范围内的所有元素。
//for (int i = 0; i < len; i++) {
//    print(nums[i]);
//}

func TestHello(t *testing.T) {
	nums := []int{0,0,1,1,1,2,2,3,3,4}
	t.Log(removeDuplicates(nums))
}
//时间最优
//国外1
func removeDuplicates1(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	// tmp := nums[0]
	n := 0
	for i := 0; i < len(nums); i++ {
		nums[n] = nums[i]
		n++
		for i < len(nums) - 1 && nums[i+1] == nums[i] {
			i++
		}
	}
	return n
}
//国外2
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	val := nums[0]
	indx := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == val {
			continue
		}
		val = nums[i]
		nums[indx] = val
		indx++
	}
	return indx
}
//国内3
func removeDuplicates3(nums []int) int {
	length := len(nums)
	replaceIndex := length //元素替换的位置

	for i := 1 ; i < length; i++  {
		if nums[i] == nums[i-1] && replaceIndex == length { //找到第一个需要替换的位置
			replaceIndex = i
		} else if replaceIndex != length && nums[replaceIndex-1] != nums[i] { //交换边界
			tmp := nums[replaceIndex]
			nums[replaceIndex] = nums[i]
			nums[i] = tmp
			replaceIndex += 1
		}
	}

	return replaceIndex
}
//国内4
func removeDuplicates4(nums []int) int {
	i := 0
	j := 0
	m := make(map[int]bool)
	if len(nums)==0{
		return 0
	}
	for {
		if !m[nums[i]] {
			m[nums[i]] = true
			i++
		}
		if m[nums[j]] {
			j++
		}
		if j >= len(nums) || i >= len(nums) {
			break
		}
		if !m[nums[j]] {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	return len(m)
}
//内存最优
//国外5
func removeDuplicates5(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] == nums[i] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}

	}
	return len(nums)
}
//国外6
func removeDuplicates6(nums []int) int {

	i := 0
	j := i + 1
	l := len(nums)

	for j < l {
		if nums[i] != nums[j] {
			nums[i + 1] = nums[j]
			i ++;
		}
		j++;
	}

	return i + 1;
}

//国内7
//国内8
