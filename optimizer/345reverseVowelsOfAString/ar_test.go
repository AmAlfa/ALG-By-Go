package _removeDuplicatesFromSortedArray

import (
	"bytes"
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

//时间最优
//国外1
func reverseVowels1(s string) string {

	start := 0
	end := len(s) - 1
	sByte := []byte(s)
	set := []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	for start < end {
		f := bytes.IndexByte(set, sByte[start])
		e := bytes.IndexByte(set, sByte[end])
		// string couldn't assign or change it's value
		// it must to convert to byte
		if f >= 0 && e >= 0 {
			sByte[start], sByte[end] = sByte[end], sByte[start]
			start++
			end--
			continue
		}
		if f < 0 {
			start++
		}
		if e < 0 {
			end--
		}
	}

	return string(sByte)
}
//国外2
//国内3
func reverseVowels3(s string) string {
	chars := []byte(s)

	vowels := "aeiouAEIOU"

	i, j := 0, len(s)-1
	for i < j {
		if strings.IndexByte(vowels, chars[i]) == -1 {
			i++
		}else if strings.IndexByte(vowels, chars[j]) == -1 {
			j--
		}else{
			chars[i], chars[j] = chars[j], chars[i]
			i, j = i+1, j-1
		}
	}
	return string(chars)
}
//国内4
func reverseVowels4(fuck string) string {
	if len(fuck)==0||len(fuck)==1{
		return fuck
	}
	vowels:=[]byte{'a','e','i','o','u','A','E','I','O','U'}
	l,r:=0,len(fuck)-1
	s:=[]byte(fuck)
	for ;l<r;{
		if !in(s[l],vowels){
			l++
		}
		if !in(s[r],vowels){
			r--
		}
		if in(s[r],vowels) &&in(s[l],vowels){
			s[l],s[r]=s[r],s[l]
			l++
			r--
		}
	}
	return string(s)
}
func in(x byte,y[]byte)bool{
	for i:=0;i<len(y);i++{
		if x==y[i]{
			return true
		}
	}
	return false
}
//内存最优
//国外5
func reverseVowels5(s string) string {
	i, j := 0, len(s)-1
	sByte := []byte(s)
	for i < j {
		if isVowel(sByte[i]) && isVowel(sByte[j]) {
			sByte[i], sByte[j] = sByte[j], sByte[i]
			i++
			j--
			continue
		}
		if !isVowel(sByte[i]) {
			i++
		}
		if !isVowel(sByte[j]) {
			j--
		}
	}
	return string(sByte)
}

func isVowel(c uint8) bool {
	if c > 'Z' {
		c = c-32
	}
	return c == 'A' || c == 'E' || c == 'I' || c == 'O' ||  c == 'U'
}
//国外6
//国内7
//国内8
