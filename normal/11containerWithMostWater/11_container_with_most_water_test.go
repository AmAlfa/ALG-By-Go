package container_with_most_water_test

import (
	"fmt"
	"testing"
)

// 给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
//
// 说明：你不能倾斜容器，且 n 的值至少为 2。

// 示例：
//
// 输入：[1,8,6,2,5,4,8,3,7]
// 输出：49



func Test(t *testing.T) {
	s := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	res := maxArea(s)
	fmt.Println(res)
	t.Log(res)

}

// 双指针
func maxArea(height []int) int {
	len := len(height)
	l, r, area := 0, len - 1, 0
	for l < r {
		area = max(min(height[l], height[r]) * (r - l), area)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return  area
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}

// 法一：
// func maxArea(height []int) int {
// 	len := len(height)
// 	area := 0
// 	for i := 0; i < len; i++ {
// 		j := i + 1
// 		for ; j < len; j++ {
// 			tmp := min(height[j], height[i]) * (j - i)
// 			if tmp > area {
// 				area = tmp
// 			}
// 		}
// 	}
//
// 	return area
// }
//
// func min(a, b int) int {
// 	if a > b {
// 		return b
// 	}
//
// 	return a
// }
