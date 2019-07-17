package _removeDuplicatesFromSortedArray

import (
	"strings"
	"testing"
)
//编写一个函数，以字符串作为输入，反转该字符串中的元音字母。
//
//示例 1:
//
//输入: "hello"
//输出: "holle"
//示例 2:
//
//输入: "leetcode"
//输出: "leotcede"
//说明:
//元音字母不包含字母"y"。

func TestHello(t *testing.T) {
	nums := "dasdaasa"
	t.Log(reverseVowels(nums))
}

func reverseVowels(s string) string {
	chars := []byte(s)

	vowels := "aeiouAEIOU"

	i, j := 0, len(s)-1
	for i < j {
		if strings.IndexByte(vowels, chars[i]) == -1 {
			i++
		} else if strings.IndexByte(vowels, chars[j]) == -1 {
			j--
		} else {
			chars[i], chars[j] = chars[j], chars[i]
			i, j = i + 1, j - 1
		}
	}
	return string(chars)
}
