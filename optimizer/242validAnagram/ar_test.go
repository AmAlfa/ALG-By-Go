package _removeDuplicatesFromSortedArray

import (
	"testing"
)
//给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
//
//示例 1:
//
//输入: s = "anagram", t = "nagaram"
//输出: true
//示例 2:
//
//输入: s = "rat", t = "car"
//输出: false
//说明:
//你可以假设字符串只包含小写字母。
//
//进阶:
//如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？

func TestHello(t *testing.T) {
	s := "asda"
	ta := "saad"
	t.Log(isAnagram(s, ta))
}

//时间最优
//国外1
func isAnagram1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var hm [26]byte

	for i := 0; i < len(s); i++ {
		hm[s[i] - 'a']++
		hm[t[i] - 'a']--
	}

	for i := 0; i < len(hm); i++ {
		if hm[i] != 0 {
			return false
		}
	}

	return true
}
//国外2
func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	seen := make(map[rune]int)
	for _, l := range s {
		seen[l]++
	}

	for _, l := range t {
		seen[l]--
		if seen[l] < 0 {
			return false
		}
	}

	return true
}
//国内3
func isAnagram3(s string, t string) bool {
	if len(s)!=len(t) {
		return false
	}
	sMap:=make([]int,26)
	for i:=0;i<len(s);i++ {
		sMap[int(s[i]-'a')]++
	}
	for i:=0;i<len(t);i++ {
		if sMap[t[i]-'a']>0 {
			sMap[t[i]-'a']--
		}else {
			return false
		}
	}
	return true
}
//国内4
func isAnagram4(s string, t string) bool {
	if len(s)!=len(t){
		return false
	}
	maps :=[26]int{}
	for i := 0; i < len(s); i++{
		maps[s[i]-'a']+=1
	}
	for j := 0; j < len(t); j++{
		maps[t[j]-'a']-=1
	}
	for _,v := range(maps){
		if v>0{
			return false
		}
	}
	return true
}
//内存最优
//国外5
func isAnagram5(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	table := make(map[int]int)
	for _, val := range s {
		table[int(val)]++
	}

	for _, val := range t {
		table[int(val)]--
		if table[int(val)] < 0 {
			return false
		}
	}

	return true
}
//国外6
func isAnagram6(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	table := make(map[int]int)
	for _, val := range s {
		table[int(val)]++
	}

	for _, val := range t {
		table[int(val)]--
		if table[int(val)] < 0 {
			return false
		}
	}

	return true
}
//国内7
//国内8
