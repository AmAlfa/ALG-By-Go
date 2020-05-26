package _removeDuplicatesFromSortedArray

import (
	"strings"
	"testing"
)
//给定一个正整数，返回它在 Excel 表中相对应的列名称。
//
//例如，
//
//    1 -> A
//    2 -> B
//    3 -> C
//    ...
//    26 -> Z
//    27 -> AA
//    28 -> AB
//    ...
//示例 1:
//
//输入: 1
//输出: "A"
//示例 2:
//
//输入: 28
//输出: "AB"
//示例 3:
//
//输入: 701
//输出: "ZY"

func TestHello(t *testing.T) {
	n := 701
	t.Log(convertToTitle(n))
}

//时间最优
//国外1
//国外2
//国内3
//国内4
//内存最优
//国外5
func toChar(n int) string {
	if n == 0 {
		n = 26
	}
	return string('A' + n - 1)
}

func convertToTitle5(n int) string {
	res := ""

	for n > 26 {
		d := n % 26
		res = toChar(d) + res
		n /= 26
		if d == 0 {
			n -= 1
		}
	}
	d := n % 26
	fmt.Println(n, d, res)
	res = toChar(d) + res
	return res
}
//国外6
func convertToTitle6(n int) string {
	if n <= 0 {
		return ""
	}
	res := []string{}
	reminder := n
	for reminder > 0 {
		reminder -= 1
		b := reminder%26
		reminder = reminder/26
		res = append([]string{string(b + 65)}, res...)
	}
	return strings.Join(res, "")
}
//国内7
//国内8
