package _removeDuplicatesFromSortedArray

import (
	"testing"
)
//给定一个仅包含大小写字母和空格 ' ' 的字符串，返回其最后一个单词的长度。
//
//如果不存在最后一个单词，请返回 0 。
//
//说明：一个单词是指由字母组成，但不包含任何空格的字符串。
//
//示例:
//
//输入: "Hello World"
//输出: 5

func TestHello(t *testing.T) {
	s := " dad"
	t.Log(lengthOfLastWord(s))
}

func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
	length := 0
	currentLength := 0
	for _, val := range s{
		if string(val) != " " {
			length += 1
		} else {
			if length != 0 {
				currentLength = length
			}
			length = 0
		}
	}
	if length == 0 {
		length = currentLength
	}

	return length
}
