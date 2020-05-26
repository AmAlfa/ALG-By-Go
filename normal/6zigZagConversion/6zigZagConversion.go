package _strDuplicateRemoval

import (
	"testing"
)

//给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
//
//示例 1：
//
//输入: "babad"
//输出: "bab"
//注意: "aba" 也是一个有效答案。
//示例 2：
//
//输入: "cbbd"
//输出: "bb"



func TestHello(t *testing.T) {
	s := ""
	//for key, val := range s {
	//	t.Log(key, val)
	//}
	t.Log(lengthOfLongestSubstring(s))

}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	m := make(map[int]int)
	strLength := len(s)
	lastRepetition, length := 0, 0
	tmpKey := 0
	for key, val := range s {
		key = int(key)
		if v, ok := m[int(val)]; ok {
			v = int(v)
			if m[int(val)] >= tmpKey {
				currentLength := key - lastRepetition
				lastRepetition = v + 1
				if length < currentLength {
					length = currentLength
				}
				tmpKey = v
			}
		}
		m[int(val)] = key
	}
	residueLength := strLength - lastRepetition
	if lastRepetition <= strLength - 1 && length <= residueLength {
		length = residueLength
	}
	return length
}
