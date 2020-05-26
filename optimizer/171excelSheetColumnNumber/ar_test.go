package _removeDuplicatesFromSortedArray

import (
	"math"
	"testing"
)
//给定一个Excel表格中的列名称，返回其相应的列序号。
//
//例如，
//
//    A -> 1
//    B -> 2
//    C -> 3
//    ...
//    Z -> 26
//    AA -> 27
//    AB -> 28
//    ...
//示例 1:
//
//输入: "A"
//输出: 1
//示例 2:
//
//输入: "AB"
//输出: 28

func TestHello(t *testing.T) {
	s := "ECKH"

	t.Log(titleToNumber(s))
}
//时间最优
//国外1

//国外2
//国内3
func titleToNumber3(s string) int {
	if len(s)==0 {
		return 0
	}
	res:=0
	for i:=0;i<len(s);i++ {
		t:=s[i]-'A'+1
		res=res*26+int(t)
	}
	return res
}
//国内4
func titleToNumber4(s string) int {
	res := 0

	b := 1
	for i := len(s)-1;i>-1;i--{
		res += int(s[i] - 'A'+1) * b
		b *= 26
	}

	return res
}
//内存最优
//国外5
//import "strings"
func titleToNumber5(s string) int {
	col_num := 0
	for i := 0; i < len(s); i++ {
		// Codepoint of "A" is 65
		num := int(s[i] - 64)
		weight := int(math.Pow(26,float64(len(s)-i-1)))
		col_num = col_num + weight * num
	}
	return col_num
}
//国外6
func titleToNumber6(s string) int {
	num := 0
	chars := []byte(s)
	for i, c := range chars {
		num = num + int(rune(c)-rune('A')+1)*int(math.Pow(26, float64(len(s)-i-1)))
	}
	return num
}
//国内7
//国内8
