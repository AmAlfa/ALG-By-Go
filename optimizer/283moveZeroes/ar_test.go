package _removeDuplicatesFromSortedArray

import (
	"testing"
)
//给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//
//示例:
//
//输入: [0,1,0,3,12]
//输出: [1,3,12,0,0]
//说明:
//
//必须在原数组上操作，不能拷贝额外的数组。
//尽量减少操作次数。

func TestHello(t *testing.T) {
	n := []int{9,6,4,2,3,5,7,0,1}
	missingNumber(n)
	t.Log(n)
}

//时间最优
//国外1
func moveZeroes1(nums []int)  {
	j := 0
	for _, num := range nums {
		if num != 0 {
			nums[j] = num
			j++
		}
	}
	for ; j < len(nums); j++ {
		nums[j] = 0
	}
}
//国外2
//国内3
func moveZeroes3(nums []int)  {
	count := 0
	l := len(nums)
	for i := 0; i < l - 1; i++{
		if nums[i-count] == 0{
			nums = append(nums[0:i-count],nums[i-count+1:]...)
			count++
		}
	}
	for i :=0 ; i < count; i++{
		nums = append(nums,0)
	}
}
//国内4
func moveZeroes4(nums []int)  {
	n := len(nums)
	if n <= 1 {
		return
	}

	lastNonZero := 0

	for i:=0; i<n; i++ {
		if nums[i] != 0  {
			if i != lastNonZero {
				nums[i], nums[lastNonZero] = nums[lastNonZero], nums[i]
			}
			lastNonZero++
		}
	}
}
//内存最优
//国外5
func moveZeroes5(nums []int)  {
	stack := []int{}

	for _, v := range nums {
		if v != 0 {
			stack = append(stack, v)
		}
	}

	var i2 int
	for i := range stack {
		nums[i] = stack[i]
		i2++
	}

	fmt.Println(i2)

	for i2 < len(nums) {
		nums[i2] = 0
		i2++
	}

}
//国外6
func moveZeroes6(nums []int)  {
	count:=0
	for i:=0;i<len(nums);i++{
		if nums[count]==0&&i+1<len(nums){
			nums=append(nums[:count],nums[count+1:]...)
			nums=append(nums,0)
			fmt.Print(nums)
		}else{
			count++
		}

	}
}
//国内7
//国内8
