package _removeDuplicatesFromSortedArray

import (
	"strings"
	"testing"
	"unicode"
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

//时间最优
//国外1
func isPalindrome1(s string) bool {
	if s == "" {
		return true
	}
	s = strings.ToLower(s)

	i, j := 0, len(s)-1
	for i < j {
		for i < j && !isAlphaNumericChar(s[i]) {
			i++
		}
		for i < j && !isAlphaNumericChar(s[j]) {
			j--
		}

		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func isAlphaNumericChar(c byte) bool {
	if ('a' <= c && c <= 'z') || ('0' <= c && c <= '9') {
		return true
	}
	return false
}
//国外2
func isPalindrome2(s string) bool {
	s = Filter(IsValid, s)
	s = strings.ToLower(s)
	return s == Reverse(s)
}

func IsValid(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsNumber(r)
}

func Filter(filter func(rune) bool, s string) string {
	var result []rune
	for _, r := range s {
		if filter(r) {
			result = append(result, r)
		}
	}
	return string(result)
}

func Reverse(s string) string {
	n := len(s)
	runes := make([]rune, n)
	for _, rune := range s {
		n--
		runes[n] = rune
	}
	return string(runes[n:])
}
//国内3
func isPalindrome3(s string) bool {
	s=strings.ToLower(s)
	i,j:=0,len(s)-1
	for i<j{
		if !('a'<=s[i]&&s[i]<='z'||'0'<=s[i]&&s[i]<='9'){
			i++
			continue
		}
		if !('a'<=s[j]&&s[j]<='z'||'0'<=s[j]&&s[j]<='9'){
			j--
			continue
		}
		if s[i]!=s[j]{
			return false
		}
		i++
		j--
	}
	return true
}
//国内4
func isPalindrome4(s string) bool {
	s=strings.ToLower(s)
	r:=[]rune(s)
	l:=len(s)
	i,j:=0,l-1
	for i<j{
		if !isalnum(r[i]){
			i++
			continue
		}
		if !isalnum(r[j]){
			j--
			continue
		}
		if s[i]!=s[j]{
			return false
		}
		i++
		j--
	}
	return true
}
func isalnum(s rune)bool{
	return unicode.IsLetter(s)||unicode.IsDigit(s)
}
//内存最优
//国外5
func isPalindrome5(s string) bool {
	if len(s) <= 1 {
		return true
	}

	i := 0
	k := len(s) - 1

	for i < k && i != k {
		if !validChar(s[i]) && !validChar(s[k]) {
			i++
			k--
			continue
		} else if !validChar(s[i]) {
			i++
			continue
		} else if !validChar(s[k]) {
			k--
			continue
		}

		iChar := strings.ToLower(string(s[i]))
		kChar := strings.ToLower(string(s[k]))

		if iChar != kChar {
			//fmt.Println("ichar", iChar, "kchar", kChar)
			return false
		}

		i++
		k--
	}

	return true
}

func validChar(char byte) bool {
	if char >= 97 && char <= 122 {
		return true
	}

	if char >= 65 && char <= 90 {
		return true
	}

	if char >= 48 && char <= 57 {
		return true
	}

	return false
}
//国外6
func isPalindrome6(s string) bool {
	if len(s) == 0 || len(s) == 1 {
		return true
	}
	alphanumerics := "abcdefghijklmnopqrstuvwxyz0123456789"
	i := 0
	j := len(s)-1
	for i < j {
		lowerI := strings.ToLower(string(s[i]))
		if !strings.Contains(alphanumerics,lowerI) {
			i++
			continue
		}
		lowerJ := strings.ToLower(string(s[j]))
		if !strings.Contains(alphanumerics,lowerJ) {
			j--
			continue
		}
		if lowerI != lowerJ {
			return false
		}
		i++
		j--
	}
	return true

}
//国内7
//国内8
