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
//时间最优
//国外1
//国外2
//国内3
//国内4
//内存最优
//国外5
func wordPattern5(pattern string, str string) bool {
	a := make(map[byte]int)
	b := make(map[string]int)

	j := 0
	for i := 0; i < len(pattern); i++ {
		var sb strings.Builder
		for j < len(str) && str[j] == 32 {
			j++
		}
		for j < len(str) && str[j] != 32 {
			sb.WriteByte(str[j])
			j++
		}
		word := sb.String()
		if word == "" {
			return false
		}

		if a[pattern[i]] != b[word] {
			return false
		}

		a[pattern[i]] = i + 1
		b[word] = i + 1
	}


	if j != len(str) {
		return false
	}

	return true
}
//国外6
func wordPattern6(pattern string, str string) bool {
	a := make(map[byte]int)
	b := make(map[string]int)

	j := 0
	for i := 0; i < len(pattern); i++ {
		var sb strings.Builder
		for j < len(str) && str[j] == 32 {
			j++
		}
		for j < len(str) && str[j] != 32 {
			sb.WriteByte(str[j])
			j++
		}
		word := sb.String()
		if word == "" {
			return false
		}

		if a[pattern[i]] != b[word] {
			return false
		}

		a[pattern[i]] = i + 1
		b[word] = i + 1
	}


	if j != len(str) {
		return false
	}

	return true
}
//国内7
//国内8
