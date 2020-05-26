package _removeDuplicatesFromSortedArray

import (
	"strings"
	"testing"
)
//给定一种规律 pattern 和一个字符串 str ，判断 str 是否遵循相同的规律。
//
//这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 str 中的每个非空单词之间存在着双向连接的对应规律。
//
//示例1:
//
//输入: pattern = "abba", str = "dog cat cat dog"
//输出: true
//示例 2:
//
//输入:pattern = "abba", str = "dog cat cat fish"
//输出: false
//示例 3:
//
//输入: pattern = "aaaa", str = "dog cat cat dog"
//输出: false
//示例 4:
//
//输入: pattern = "abba", str = "dog dog dog dog"
//输出: false


func TestHello(t *testing.T) {
	s := "aaaa"
	a := "hello hello hello hello hello"

	t.Log(wordPattern(s, a))
}

func wordPattern(pattern string, str string) bool {
	strArr := strings.Fields(str)
	if len(pattern)!=len(strArr) {
		return false
	}
	strToArr := make(map[byte]string)
	arrToStr := make(map[string]byte)
	for i := 0; i < len(pattern); i++ {
		arrVal, arrIn := strToArr[pattern[i]]
		strVal, strIn := arrToStr[strArr[i]]
		if (arrIn && arrVal!=strArr[i]) || (strIn && strVal!=pattern[i]) {
			return false
		} else {
			strToArr[pattern[i]] = strArr[i]
			arrToStr[strArr[i]] = pattern[i]
		}
	}
	return true
}
