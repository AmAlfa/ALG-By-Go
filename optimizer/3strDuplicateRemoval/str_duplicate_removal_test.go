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
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func TestHello(t *testing.T) {
	s := ""
	//for key, val := range s {
	//	t.Log(key, val)
	//}
	t.Log(lengthOfLongestSubstring(s))

}
//速度最快
func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	start, result := -1, 0
	data := make(map[rune]int)
	for i, v := range s {
		if index, ok := data[v]; ok && index > start {
			start = index
		}
		data[v] = i
		if i-start > result {
			result = i-start
		}
	}
	return result
}

func lengthOfLongestSubstring1(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	hash := make(map[byte]bool)
	begin, end := 0, 0
	max := 0
	for ; end < len(s) && end >= begin;end++ {
		if _, ok := hash[s[end]]; ok {
			for s[begin] != s[end] {
				delete(hash, s[begin])
				begin++
			}
			begin ++
		} else {
			hash[s[end]] = true
			if end-begin+1 > max {
				max = end - begin + 1
			}
		}
	}
	return max
}


//内存最小
func lengthOfLongestSubstring3(s string) int {
	seen := make([]int, 256)

	for i := 0; i < 256; i++ {
		seen[i] = -1
	}

	start := -1
	maxLen := 0

	for i := range s {
		if seen[int(s[i])] > start {
			start = seen[int(s[i])]
		}

		seen[int(s[i])] = i

		if maxLen < i-start {
			maxLen = i - start
		}
	}

	return maxLen
}


func lengthOfLongestSubstring4(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	count := 1
	maxcount := 1
	i := 0
	for j := 1; i < len(s) && j < len(s); {
		if strings.Contains(s[i:j],s[j:j+1]) {
			if count > maxcount {
				maxcount = count
			}
			count = 1
			i += 1
			j = i + 1
		} else {
			count += 1
			j += 1
		}
	}
	if count > maxcount {
		maxcount = count
	}
	return maxcount
}
