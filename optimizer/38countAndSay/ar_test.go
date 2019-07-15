package _removeDuplicatesFromSortedArray

import (
	"strconv"
	"strings"
	"testing"
)
//报数序列是一个整数序列，按照其中的整数的顺序进行报数，得到下一个数。其前五项如下：
//
//1.     1
//2.     11
//3.     21
//4.     1211
//5.     111221
//1 被读作  "one 1"  ("一个一") , 即 11。
//11 被读作 "two 1s" ("两个一"）, 即 21。
//21 被读作 "one 2",  "one 1" （"一个二" ,  "一个一") , 即 1211。
//
//给定一个正整数 n（1 ≤ n ≤ 30），输出报数序列的第 n 项。
//
//注意：整数顺序将表示为一个字符串。
//
// 
//
//示例 1:
//
//输入: 1
//输出: "1"
//示例 2:
//
//输入: 4
//输出: "1211"

func TestHello(t *testing.T) {
	n := 7
	t.Log(countAndSay(n))
}

//时间最优
//国外1

func countAndSay1(n int) string {
	/**
	时间复杂度 O(n * K), K 为 n的 字符串长度
	 */

	if n <= 1 {
		return "1"
	}
	res := "1"
	for i := 2; i <= n; i ++ {
		var sb strings.Builder
		idx, count, j := 0, 1, 1
		for ; j < len(res); j ++{
			if res[idx] == res[j] {
				count ++
				continue
			}
			sb.WriteString(strconv.Itoa(count))
			sb.WriteByte(res[idx])
			count = 1
			idx = j
		}
		sb.WriteString(strconv.Itoa(count))
		sb.WriteByte(res[idx])
		res = sb.String()
	}
	return res
}
//国外2
func countAndSay2(n int) string {
	res := "1"
	for start := 1; start < n; start++ {
		fmt.Println(res)
		tmp := ""
		for i, start := 0, 0; i <= len(res); i++ {
			if i == len(res) || res[i] != res[start] {
				tmp += strconv.Itoa(i - start)
				tmp += string(res[start])
				start = i
			}
		}
		res = tmp
	}
	return res
}
//国内3
func countAndSay3(n int) string {
	str:=[]byte{'1'}
	for n>1{
		tmp:=make([]byte,0,len(str)*2)
		i,j:=0,1
		for i<len(str){
			for j<len(str)&&str[i]==str[j]{
				j++
			}
			tmp=append(tmp,byte(j-i+'0'),str[i])
			i=j
		}
		str=tmp
		n--
	}
	return string(str)
}
//国内4
func countAndSay4(n int) string {
	var res string
	var i int
	res = "1"
	for i = 0; i < n - 1; i++{
		res = count(res)
	}
	return res
}

func count(s string) string {
	var res string
	var count int
	count = 0
	res = ""
	for index, i := range s{
		count++
		if index == len(s) - 1 || s[index + 1] != s[index]{
			res += string(count + 48) + string(i)
			count = 0
		}
	}
	return res
}
//内存最优
//国外5
func countAndSay5(n int) string {
	if n == 1 {
		return "1"
	}
	if n == 2 {
		return "11"
	}
	if n < 0 || n > 30 {
		return "0"
	}

	var out strings.Builder
	out.WriteString("11")

	for i := 2; i < n; i++ {
		var count int
		var cur rune

		for _, v := range out.String() {
			if cur == 0 {
				out.Reset()
				cur = v
			}

			if cur == v {
				count++
			} else {
				out.WriteString(strconv.Itoa(count))
				out.WriteRune(cur)
				cur = v
				count = 1
			}
		}
		out.WriteString(strconv.Itoa(count))
		out.WriteRune(cur)
	}

	return out.String()
}
//国外6
func countAndSay6(n int) string {
	if n == 1 {
		return "1"
	}
	if n == 2 {
		return "11"
	}
	if n < 0 || n > 30 {
		return "0"
	}

	var out strings.Builder
	out.WriteString("11")

	for i := 2; i < n; i++ {
		var count int
		var cur rune

		for _, v := range out.String() {
			if cur == 0 {
				out.Reset()
				cur = v
			}

			if cur == v {
				count++
			} else {
				out.Write([]byte(strconv.Itoa(count)))
				out.WriteRune(cur)
				cur = v
				count = 1
			}
		}
		out.Write([]byte(strconv.Itoa(count)))
		out.WriteRune(cur)
	}

	return out.String()
}
//国内7
//国内8
