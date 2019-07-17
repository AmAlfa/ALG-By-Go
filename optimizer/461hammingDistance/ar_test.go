package _removeDuplicatesFromSortedArray

import (
	"strconv"
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

//时间最优
//国外1
//国外2
//国内3
//国内4
//内存最优
//国外5
func hammingDistance5(x int, y int) int {
	strx := strconv.FormatInt(int64(x), 2)
	stry := strconv.FormatInt(int64(y), 2)

	if len(strx)> len(stry){
		format:=fmt.Sprint("%0", len(strx), "s")
		stry = fmt.Sprintf(format, stry)
	}else if len(stry) > len(strx) {
		format:=fmt.Sprint("%0", len(stry), "s")
		strx = fmt.Sprintf(format, strx)
	}

	fmt.Println(strx, "    ",stry)
	runex := []rune(strx)
	runey := []rune(stry)

	count:=0
	for i:=0; i< len(runex);i++{
		if runex[i] != runey[i] {
			count++
		}
	}

	return count
}
//国外6
func hammingDistance6(x int, y int) int {
	hd := 0
	for x != 0 || y != 0 {
		hd += (x&1)^(y&1)
		x >>= 1
		y >>= 1
	}
	return hd
}
//国内7
//国内8
