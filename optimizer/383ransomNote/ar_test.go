package _removeDuplicatesFromSortedArray

import (
	"testing"
)

//给定一个赎金信 (ransom) 字符串和一个杂志(magazine)字符串，判断第一个字符串ransom能不能由第二个字符串magazines里面的字符构成。如果可以构成，返回 true ；否则返回 false。
//
//(题目说明：为了不暴露赎金信字迹，要从杂志上搜索各个需要的字母，组成单词来表达意思。)
//
//注意：
//
//你可以假设两个字符串均只含有小写字母。
//
//canConstruct("a", "b") -> false
//canConstruct("aa", "ab") -> false
//canConstruct("aa", "aab") -> true

func TestHello(t *testing.T) {
	nums := "asda"
	target := "dasda"
	t.Log(canConstruct(nums, target))
}

//时间最优
//国外1
func canConstruct1(ransomNote string, magazine string) bool {
	lc := make([]int, 200)
	for x := 0; x < len(magazine); x++ {
		c := int(magazine[x])
		if lc[c] == 0 {
			lc[c] = 1
		} else {
			lc[c]++
		}
	}
	for x := 0; x < len(ransomNote); x++ {
		c := int(ransomNote[x])
		if lc[c] != 0 {
			lc[c]--
		} else {
			return false
		}
	}
	return true
}
//国外2
func canConstruct2(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	count := [26]int{}

	for _, c := range magazine {
		count[c - 'a']++;
	}

	for _, c := range ransomNote {
		count[c - 'a']--;
		if count[c - 'a'] < 0 {
			return false
		}
	}

	return true
}
//国内3
func canConstruct3(ransomNote string, magazine string) bool {
	resource := make([]int, 26)
	for _, item := range magazine {
		resource[item - 'a']++
	}
	for _, item := range ransomNote {
		resource[item - 'a']--
	}
	for _, item := range resource {
		if item < 0 {
			return false
		}
	}
	return true
}
//国内4
func canConstruct4(ransomNote string, magazine string) bool {
	var letters [26]int

	for _,v := range magazine {
		letters[v-'a']++
	}

	for _,v := range ransomNote {
		letters[v-'a']--
	}

	for _,v := range letters {
		if v < 0 {
			return false
		}
	}
	return true
}
//内存最优
//国外5
func canConstruct5(ransomNote string, magazine string) bool {
	if len(ransomNote) == 0 {
		return true
	}
	if len(magazine) < len(ransomNote) {
		return false
	}


	ransomNoteMap := make(map[int32]int)
	for _, c := range ransomNote {
		if _, ok := ransomNoteMap[c]; ok {
			ransomNoteMap[c] = ransomNoteMap[c] + 1
			continue
		}
		ransomNoteMap[c] = 1
	}

	for _, v := range magazine {
		if len(ransomNoteMap) == 0 {
			return true
		}

		if n, ok := ransomNoteMap[v]; ok  {
			if n == 1 {
				delete(ransomNoteMap, v)
				continue
			}
			ransomNoteMap[v] = ransomNoteMap[v] - 1
			continue
		}
	}

	return len(ransomNoteMap) == 0
}
//国外6
func canConstruct6(ransomNote string, magazine string) bool {
if len(ransomNote) == 0 {
return true
}
if len(magazine) < len(ransomNote) {
return false
}

magazineMap := make(map[int32]int)
for _, c := range magazine {
if _, ok := magazineMap[c]; ok {
magazineMap[c] = magazineMap[c] + 1
continue
}
magazineMap[c] = 1
}

for _, v := range ransomNote {
if n, ok := magazineMap[v]; ok {
if n == 1 {
delete(magazineMap, v)
continue
}
magazineMap[v] = magazineMap[v] - 1
continue
}
return false
}

return true
}
//国内7
//国内8
