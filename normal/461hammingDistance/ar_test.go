package _removeDuplicatesFromSortedArray

import (
	"testing"
)
//两个整数之间的汉明距离指的是这两个数字对应二进制位不同的位置的数目。
//
//给出两个整数 x 和 y，计算它们之间的汉明距离。
//
//注意：
//0 ≤ x, y < 231.
//
//示例:
//
//输入: x = 1, y = 4
//
//输出: 2
//
//解释:
//1   (0 0 0 1)
//4   (0 1 0 0)
//       ↑   ↑
//
//上面的箭头指出了对应二进制位不同的位置。

func TestHello(t *testing.T) {
	a := 123
	b := 13
	t.Log(hammingDistance(a, b))
}

func hammingDistance(x int, y int) int {
	v := x^y
	v = (v & 0x55555555) + ((v >> 1) & 0x55555555)
	v = (v & 0x33333333) + ((v >> 2) & 0x33333333)
	v = (v & 0x0f0f0f0f) + ((v >> 4) & 0x0f0f0f0f)
	v = (v & 0x00ff00ff) + ((v >> 8) & 0x00ff00ff)
	v = (v & 0x0000ffff) + ((v >> 16) & 0x0000ffff)

	return v
}
