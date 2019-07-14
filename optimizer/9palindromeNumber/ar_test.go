package _palindromeNumber

import (
	"math"
	"strconv"
	"testing"
)

//判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
//
//示例 1:
//输入: 121
//输出: true

//示例 2:
//输入: -121
//输出: false
//解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。

//示例 3:
//输入: 10
//输出: false
//解释: 从右向左读, 为 01 。因此它不是一个回文数。
//
//进阶:
//你能不将整数转为字符串来解决这个问题吗？

func TestHello(t *testing.T) {
	a := -29349238
	t.Log(isPalindrome(a))
}
//时间
//国内1：
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var num int
	cache := x
	for  {
		n := x % 10

		if x == 0 && n == 0 {
			break
		}
		x = x / 10
		num = num * 10 + n
	}
	if cache == num {
		return true
	}

	return false

}
//国内2：
func isPalindrome1(x int) bool {
	if x < 0 {
		return false
	}
	var y int
	z := x

	for x != 0 {
		y = y * 10 + x % 10
		x = x / 10
	}

	if y == z {
		return true
	}

	return false
}
//国外1
func isPalindrome3(x int) bool {
	if x < 0 {
		return false
	}
	s := []rune(strconv.Itoa(x))
	for i,j := 0, len(s)-1; i < j; i,j = i+1,j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
//国外2
func isPalindrome4(x int) bool {
	if x < 0 {
		return false
	}
	s := []byte(strconv.Itoa(x))
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}
	return true
}


//内存
//国外1
func isPalindromeN(x, n int) bool {

	switch n {
	case 1:
		return true
	case 2:
		return x/10 == x%10
	case 3:
		return x/100 == x%10
	case 4:
		first := x / 1000
		return first == x%10 && isPalindromeN(x/10-first*100, 2)
	case 5:
		first := x / 10000
		return first == x%10 && isPalindromeN(x/10-first*1000, 3)
	case 6:
		first := x / 100000
		return first == x%10 && isPalindromeN(x/10-first*10000, 4)
	case 7:
		first := x / 1000000
		return first == x%10 && isPalindromeN(x/10-first*100000, 5)
	case 8:
		first := x / 10000000
		return first == x%10 && isPalindromeN(x/10-first*1000000, 6)
	case 9:
		first := x / 100000000
		return first == x%10 && isPalindromeN(x/10-first*10000000, 7)
	case 10:
		first := x / 1000000000
		return first == x%10 && isPalindromeN(x/10-first*100000000, 8)
	default:
		digits := make([]int, 0, 17)
		for i := x; i > 0; i = i / 10 {
			digits = append(digits, i%10)
		}
		l := len(digits)
		for i := 0; i < l/2; i++ {
			if digits[i] != digits[l-i-1] {
				return false
			}
		}
		return true
	}

}

func isPalindrome5(x int) bool {

	switch {
	case x < 0:
		return false
	case x == 0:
		return true
	case x < 10:
		return isPalindromeN(x, 1)
	case x < 100:
		return isPalindromeN(x, 2)
	case x < 1000:
		return isPalindromeN(x, 3)
	case x < 10000:
		first := x / 1000
		return first == x%10 && isPalindromeN(x/10-first*100, 2)
	case x < 100000:
		first := x / 10000
		return first == x%10 && isPalindromeN(x/10-first*1000, 3)
	case x < 1000000:
		first := x / 100000
		return first == x%10 && isPalindromeN(x/10-first*10000, 4)
	case x < 10000000:
		first := x / 1000000
		return first == x%10 && isPalindromeN(x/10-first*100000, 5)
	case x < 100000000:
		first := x / 10000000
		return first == x%10 && isPalindromeN(x/10-first*1000000, 6)
	case x < 1000000000:
		first := x / 100000000
		return first == x%10 && isPalindromeN(x/10-first*10000000, 7)
	case x < 10000000000:
		first := x / 1000000000
		return first == x%10 && isPalindromeN(x/10-first*100000000, 8)
	case x < 100000000000:
		first := x / 10000000000
		return first == x%10 && isPalindromeN(x/10-first*1000000000, 9)
	case x < 1000000000000:
		first := x / 100000000000
		return first == x%10 && isPalindromeN(x/10-first*10000000000, 10)
	default:
		digits := make([]int, 0, 17)
		for i := x; i > 0; i = i / 10 {
			digits = append(digits, i%10)
		}
		l := len(digits)
		for i := 0; i < l/2; i++ {
			if digits[i] != digits[l-i-1] {
				return false
			}
		}
		return true
	}

}
//国外2
func isPalindrome6(x int) bool {
	if x < 0 {
		return false
	}

	if x >= 0 && x < 10 {
		return true
	}
	ln10 := math.Log10(float64(x))
	half := ln10 / 2
	ceil := int(math.Ceil(half))
	tail := x % int(math.Pow10(ceil))
	head := x / int(math.Pow10(int(ln10) + 1 - ceil))

	reverse := 0
	for i := 0; i <= int(half); i++ {
		reverse = (reverse * 10) + (tail % 10)
		tail = tail / 10
	}

	return head == reverse
}
