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

func plusOne(digits []int) []int {
	lastKey := len(digits) - 1
	addCount := 0
	if digits[lastKey] == 9 {
		digits[lastKey] = 10
		for i := lastKey; i >= 0; i-- {
			tmp := (digits[i] + addCount) / 10
			digits[i] = (digits[i] + addCount) % 10
			addCount = tmp
		}
		if digits[0] == 0 && addCount == 1 {
			digits[0] = 1
			digits = append(digits, 0)
		}
	} else {
		digits[lastKey] += 1
	}

	return digits
}
