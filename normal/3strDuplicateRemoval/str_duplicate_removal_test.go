package _strDuplicateRemoval

import (
	"testing"
)

//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//
//示例 1:
//
//输入: "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//示例 2:
//
//输入: "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//示例 3:
//
//输入: "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

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
