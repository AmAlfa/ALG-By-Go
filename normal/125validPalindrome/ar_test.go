package _removeDuplicatesFromSortedArray

import (
	"regexp"
	"strings"
	"testing"
)
//给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
//
//说明：本题中，我们将空字符串定义为有效的回文串。
//
//示例 1:
//
//输入: "A man, a plan, a canal: Panama"
//输出: true
//示例 2:
//
//输入: "race a car"
//输出: false


func TestHello(t *testing.T) {
	s := "race a car"
	t.Log(isPalindrome(s))
}

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	reg := regexp.MustCompile(`[^0-9a-zA-Z]`)
	str := reg.ReplaceAllString(s, "")
	str = strings.ToLower(str)
	strLen := len(str)
	for i := 0; i < strLen / 2; i++ {
		if string(str[i]) != string(str[strLen - 1 - i]) {
			return false
		}
	}
	return true
}
