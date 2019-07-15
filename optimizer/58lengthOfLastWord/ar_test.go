package _removeDuplicatesFromSortedArray

import (
	"strings"
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
//时间最优
//国外1
func lengthOfLastWord1(s string) int {
	//word := make([]string, 5)
	word := strings.Fields(s)
	//word = append(word,)
	if len(word) == 0 {
		return 0
	} else if len(word) == 1 {
		return len(word[0])
	} else {
		return len(word[(len(word)-1)])
	}
}
//国外2
func lengthOfLastWord2(s string) int {
	words := strings.Fields(s)
	l := len(words)
	if l == 0 {
		return 0
	}
	return len(words[l-1])

}
//国内3
//国内4
//内存最优
//国外5
//国外6
//国内7
//国内8
