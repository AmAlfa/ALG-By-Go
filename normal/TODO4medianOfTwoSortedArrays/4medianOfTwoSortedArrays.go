package _strDuplicateRemoval

import (
	"testing"
)

//给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
//
//请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
//
//你可以假设 nums1 和 nums2 不会同时为空。
//
//示例 1:
//
//nums1 = [1, 3]
//nums2 = [2]
//
//则中位数是 2.0
//示例 2:
//
//nums1 = [1, 2]
//nums2 = [3, 4]
//
//则中位数是 (2 + 3)/2 = 2.5


func TestHello(t *testing.T) {
	s := ""
	//for key, val := range s {
	//	t.Log(key, val)
	//}
	t.Log(lengthOfLongestSubstring(s))

}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	m := make(map[int]int)
	strLength := len(s)
	lastRepetition, length := 0, 0
	tmpKey := 0
	for key, val := range s {
		key = int(key)
		if v, ok := m[int(val)]; ok {
			v = int(v)
			if m[int(val)] >= tmpKey {
				currentLength := key - lastRepetition
				lastRepetition = v + 1
				if length < currentLength {
					length = currentLength
				}
				tmpKey = v
			}
		}
		m[int(val)] = key
	}
	residueLength := strLength - lastRepetition
	if lastRepetition <= strLength - 1 && length <= residueLength {
		length = residueLength
	}
	return length
}
