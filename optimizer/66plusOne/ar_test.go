package _removeDuplicatesFromSortedArray

import (
	"testing"
)
//给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
//
//最高位数字存放在数组的首位， 数组中每个元素只存储一个数字。
//
//你可以假设除了整数 0 之外，这个整数不会以零开头。
//
//示例 1:
//
//输入: [1,2,3]
//输出: [1,2,4]
//解释: 输入数组表示数字 123。
//示例 2:
//
//输入: [4,3,2,1]
//输出: [4,3,2,2]
//解释: 输入数组表示数字 4321。

func TestHello(t *testing.T) {
	nums := []int{9, 9}
	t.Log(plusOne(nums))
}

////时间最优
////国外1
func plusOne(digits []int) []int {
	carry:=1
	for i:=len(digits)-1; i >= 0  ; i-- {
		digits[i]=digits[i]+carry
		if digits[i] >= 10 {
			digits[i]=0
			carry=1
		}else{
			carry=0
		}

		fmt.Println(digits[i])
	}
	if carry==1{

		digits= append([]int{1}, digits...)
	}
	return digits
}
////国外2
func plusOne2(digits []int) []int {
	newDigits := make([]int, len(digits))
	index := len(digits) - 1
	overflow := 0
	add := 1
	for index >= 0 {
		value := digits[index]

		// add
		plusOne := value + add + overflow
		add = 0
		newValue := plusOne % 10
		overflow = plusOne / 10

		newDigits[index] = newValue

		index--
	}

	if overflow > 0 {
		fmt.Println(newDigits)
		temp := make([]int, 1)
		temp[0] = overflow
		for _, i := range newDigits {
			temp = append(temp, i)
		}
		newDigits = temp
	}
	return newDigits
}
////国内3
func plusOne3(digits []int) []int {
	up:=1
	i:=len(digits)-1
	for up>0&&i>=0{
		add:=digits[i]+up
		digits[i]=add%10
		up=add/10
		i--
	}
	if up>0{
		digits=append([]int{up},digits...)
	}
	return digits
}
////国内4
////内存最优
////国外5
////国外6
////国内7
////国内8
