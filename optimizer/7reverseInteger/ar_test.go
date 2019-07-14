package _reverseInteger

import (
	"math"
	"testing"
)

//给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
//
//示例 1:
//
//输入: 123
//输出: 321
// 示例 2:
//
//输入: -123
//输出: -321
//示例 3:
//
//输入: 120
//输出: 21
//注意:
//
//假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−2^31,  2^31 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。

func TestHello(t *testing.T) {
	a := -29349238
	t.Log(reverse(a))
}

//时间最优
//1
func reverse(x int) int {
	rev := 0
	if x > 0 {
		rev = revPos(x)
		if rev > 2147483647 {return 0}
	} else if x < 0 {
		rev = -revPos(-x)
		if rev < -2147483648 {return 0}
	}
	return rev
}

func revPos(x int) int {
	a := 0
	rev := 0
	for x > 0 {
		a = x%10
		x /= 10
		rev *= 10
		rev += a
	}
	return rev
}
//2
func reverse1(x int) int {
	ans := 0
	for; x!=0; {
		ans = ans*10+x%10
		x/=10
	}
	if ans > 2147483647 || ans < -2147483648 {
		return 0
	}
	return ans
}
//3
func reverse3(x int) int {
	var rev int
	for x != 0{
		pop := x % 10
		x /= 10
		rev = rev * 10  + pop

		if !(-(1<<31) <= rev && rev <= ((1<<31)-1)) {
			return 0
		}
	}
	return rev
}
//空间最优
//1
func revers4(x int) int {
	// symbol := 1
	// if x < 0 {
	//     symbol = -1
	//     x *= symbol
	// }
	// for x % 10 == 0 {
	//     x /= 10
	// }
	res := 0
	for x != 0 {
		pop := x % 10
		x = x / 10
		res = res * 10 + pop
	}
	intmax := (math.Pow(2, 31)) - 1
	intmin := math.Pow(-2, 31)
	if (res > int(intmax)) || (res < int(intmin)) {
		res = 0
	}

	return res
	// return res * symbol
}
//2
func reverse5(x int) int {
	// symbol := 1
	// if x < 0 {
	//     symbol = -1
	//     x *= symbol
	// }
	// for x % 10 == 0 {
	//     x /= 10
	// }
	//     res := 0
	//     for x != 0 {
	//         pop := x % 10
	//         x = x / 10
	//         res = res * 10 + pop
	//     }

	res := 0
	for x != 0 {
		res *= 10
		res += x % 10
		x = x / 10
	}

	intmax := (math.Pow(2, 31)) - 1
	intmin := math.Pow(-2, 31)
	if (res > int(intmax)) || (res < int(intmin)) {
		res = 0
	}

	return res
	// return res * symbol
}

// func reverse(x int) int {
//     var temp, pop, rev int
//     for ; x != 0; {
//         pop = x % 10;
//         x /= 10;
//         temp = rev * 10 + pop;
//         rev = temp
//     }
//     intmax := (math.Pow(2, 31)) - 1
//     intmin := math.Pow(-2, 31)
//     if (rev > int(intmax)) || (rev < int(intmin)) {
//         rev = 0
//     }
//     return rev

// }
