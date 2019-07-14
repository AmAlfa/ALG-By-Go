package _longestCommonPrefix

import (
	"sort"
	"testing"
)
//编写一个函数来查找字符串数组中的最长公共前缀。
//
//如果不存在公共前缀，返回空字符串 ""。
//
//示例 1:
//
//输入: ["flower","flow","flight"]
//输出: "fl"
//示例 2:
//
//输入: ["dog","racecar","car"]
//输出: ""
//解释: 输入不存在公共前缀。
//说明:
//
//所有输入只包含小写字母 a-z 。

func TestHello(t *testing.T) {
	strs := []string{"aca","cba"}

	t.Log(longestCommonPrefix(strs))
}

//时间最优
//国外1
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	res := ""
	hasCommon := true
	for j := 0; hasCommon && j < len(strs[0]); j++ {
		c := strs[0][j]
		for i := 1; i < len(strs); i++ {
			if j >= len(strs[i]) || strs[i][j] != c {
				hasCommon = false
			}
		}
		if hasCommon {
			res += string(c)
		}
	}
	return res
}
//国外2
//国内3
func longestCommonPrefix3(strs []string) string {
	if strs == nil || len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	for _, str := range strs {
		length := 0
		if len(prefix) < len(str) {
			length = len(prefix)
		} else {
			length = len(str)
		}
		prefix = prefix[0:length]
		for i := 0; i < length; i++ {
			if prefix[i] != str[i] {
				prefix = prefix[0:i]
				break
			}
		}
	}
	return prefix
}
//国内4
func longestCommonPrefix4(strs []string) string {
	if strs==nil||len(strs)<=0{
		return ""
	}
	sort.Strings(strs)
	first:=strs[0]
	last:=strs[len(strs)-1]
	var length,i int
	if len(first)<len(last){
		length=len(first)
	}else{
		length=len(last)
	}
	for i<length&&first[i]==last[i]{
		i++
	}
	return first[:i]
}

//空间最优
//国外5
func longestCommonPrefix5(strs []string) string {
	var shortest string
	var pfxlen int

	if len(strs) < 1 {
		return ""
	}

	shortest = strs[0]
	for _, str := range strs {

		if len(str) < len(shortest) {
			shortest = str
			fmt.Println(shortest)
		}
	}

	for p, _ := range shortest {
		for _, str := range strs {
			if str[p] != shortest[p] {
				return shortest[:pfxlen]
			}
		}

		pfxlen++
	}

	return shortest[:pfxlen]
}
//国外6
func getminlen(strs []string) int {
	min := 10000000000
	for _, cur := range strs {
		if len(cur) < min {
			min = len(cur)
			fmt.Println("curlen ", min)
		}
	}
	return min
}

func adjustcurposition(strs []string, curposition int) bool {
	cur := strs[0][curposition]
	for i := 1; i < len(strs); i++ {
		if strs[i][curposition] != cur {
			return false
		}
	}
	return true
}

func longestCommonPrefix6(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	minlen := getminlen(strs)
	if minlen == 0 {
		return ""
	}
	fmt.Println(minlen)
	var i int = 0
	for ; i < minlen; i++ {
		fmt.Println(i)
		if !adjustcurposition(strs, i) {
			break
		}
	}
	if i == 0 {
		return ""
	}
	return strs[0][:i]
}
//国内7
//国内8
