package _removeDuplicatesFromSortedArray

import (
	"strconv"
	"testing"
)
//写一个程序，输出从 1 到 n 数字的字符串表示。
//
//1. 如果 n 是3的倍数，输出“Fizz”；
//
//2. 如果 n 是5的倍数，输出“Buzz”；
//
//3.如果 n 同时是3和5的倍数，输出 “FizzBuzz”。
//
//示例：
//
//n = 15,
//
//返回:
//[
//    "1",
//    "2",
//    "Fizz",
//    "4",
//    "Buzz",
//    "Fizz",
//    "7",
//    "8",
//    "Fizz",
//    "Buzz",
//    "11",
//    "Fizz",
//    "13",
//    "14",
//    "FizzBuzz"
//]


func TestHello(t *testing.T) {
	nums := 3
	t.Log(fizzBuzz(nums))
}

//时间最优
//国外1
//国外2
func fizzBuzz2(n int) []string {
	if n == 0 {
		return make([]string, 0)
	}

	s := make([]string, n)
	for i:=1; i<=n; i++ {
		if i % 15 == 0 {
			s[i-1] = "FizzBuzz"
		} else if i % 5 == 0 {
			s[i-1] = "Buzz"
		} else if i % 3 == 0 {
			s[i-1] = "Fizz"
		} else {
			s[i-1] = strconv.Itoa(i)
		}

	}

	return  s
}
//国内3
func fizzBuzz3(n int) []string {
	res := []string{}
	for i := 1; i <= n; i++ {
		s := strconv.Itoa(i)
		if i % 15 == 0 {
			s = "FizzBuzz"
		} else if i % 5 == 0 {
			s = "Buzz"
		} else if i % 3 == 0 {
			s = "Fizz"
		}
		res = append(res, s)
	}
	return res
}
//国内4
func fizzBuzz4(n int) []string {
	i := 1
	strArr := []string{}
	for i <= n {
		if i%5 == 0 && i%3 == 0 {
			strArr = append(strArr, "FizzBuzz")
		} else if i%5 == 0 {
			strArr = append(strArr, "Buzz")
		} else if i%3 == 0 {
			strArr = append(strArr, "Fizz")
		} else {
			strArr = append(strArr, strconv.Itoa(i))
		}
		i++
	}
	return strArr
}
//内存最优
//国外5
func fizzBuzz5(n int) []string {
	answer := []string{}

	for i := 1; i < n+1; i++ {
		var v string

		if i % 3 == 0 { v += "Fizz"	}
		if i % 5 == 0 { v += "Buzz" }

		if v == "" { v = strconv.Itoa(i) }
		answer = append(answer, v)
	}

	return answer

}
//国外6
func fizzBuzz6(n int) []string {
	out := []string{}

	fizz := []byte("Fizz")
	buzz := []byte("Buzz")
	fizzbuzz := []byte("FizzBuzz")

	for i := 1; i < n+1; i++ {
		if i % 15 == 0 {
			out = append(out, string(fizzbuzz))
		} else if i % 3 == 0 {
			out = append(out, string(fizz))
		} else if i % 5 == 0 {
			out = append(out, string(buzz))
		} else {
			out = append(out, strconv.Itoa(i))
		}
	}

	return out
}
//国内7
//国内8
