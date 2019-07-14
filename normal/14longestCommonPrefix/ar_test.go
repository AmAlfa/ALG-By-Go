package _longestCommonPrefix

import (
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

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	commonProfix := strs[0]
	for _, Val := range strs {
		if len(Val) == 0 {
			return ""
		}
		//flower
		tmpProfix := ""
		for strValKey, strVal := range Val {//0,f
			if len(commonProfix) > strValKey && string(strVal) == string(commonProfix[strValKey]) {
				tmpProfix += string(strVal)
			} else {
				break
			}
		}
		commonProfix = tmpProfix
	}
	return commonProfix
}
