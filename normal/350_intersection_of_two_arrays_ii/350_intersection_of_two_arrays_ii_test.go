package intersectionoftwoarraysii

import (
	"fmt"
	"sort"
	"testing"
)

// 给定两个数组，编写一个函数来计算它们的交集。
//
//  
//
// 示例 1：
//
// 输入：nums1 = [1,2,2,1], nums2 = [2,2]
// 输出：[2,2]
// 示例 2:
//
// 输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
// 输出：[4,9]
//  
//
// 说明：
//
// 输出结果中每个元素出现的次数，应与元素在两个数组中出现次数的最小值一致。
// 我们可以不考虑输出结果的顺序。
// 进阶：
//
// 如果给定的数组已经排好序呢？你将如何优化你的算法？
// 如果 nums1 的大小比 nums2 小很多，哪种方法更优？
// 如果 nums2 的元素存储在磁盘上，磁盘内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？

func Test(t *testing.T) {
	nums1 := []int{1,2,2,1}
	nums2 := []int{2,2,1,2}
	res := intersect(nums1, nums2)
	fmt.Println(res)
	t.Log(res)

}

// func intersect(nums1 []int, nums2 []int) []int {
// 	m := make(map[int]int)
// 	if len(nums1) == 0 || len(nums2) == 0 {
// 		return []int{}
// 	}
// 	for _, v := range nums1{
// 		m[v] += 1
// 	}
//
// 	k := 0
// 	for _, v := range nums2{
// 		if _, ok := m[v]; ok && m[v] > 0 {
// 			m[v]--
// 			nums2[k] = v
// 			k++
// 		}
// 	}
//
// 	return nums2[0:k]
// }

// func intersect(nums1 []int, nums2 []int) []int {
// 	var result []int
// 	m := make(map[int]int)
// 	if len(nums1) == 0 || len(nums2) == 0 {
// 		return result
// 	}
// 	for _, v := range nums1{
// 		m[v] += 1
// 	}
// 	for _, v := range nums2{
// 		if _, ok := m[v]; ok && m[v] > 0 {
// 			m[v]--
// 			result = append(result, v)
// 		}
// 	}
//
// 	return nums2[0:k]
// }

func intersect(nums1 []int, nums2 []int) []int {
	l1 := len(nums1)
	l2 := len(nums2)
	if l1 == 0 || l2 == 0 {
		return []int{}
	}

	sort.Ints(nums1)
	sort.Ints(nums2)
	var result []int

	for i, j := 0, 0; i < l1 && j < l2; {
		switch {
		case nums1[i] == nums2[j]:
			result = append(result, nums1[i])
			i++
			j++
		case nums1[i] < nums2[j]:
			i++
		case nums1[i] > nums2[j]:
			j++
		}
	}

	return result
}
