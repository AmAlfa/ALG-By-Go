package pascal_triangle

import "fmt"

//获取节点
func getNode(row int, li int) int {
	if li > row {
		return 0
	}
	if row == 1 || li == 1 {
		return 1
	}
	leftNum := getNode(row -1, li - 1)
	rightNum := getNode(row -1, li)

	return leftNum + rightNum

}

//获取三角形
func getTriangle(row int) [][]int {
	if row == 0 {
		return [][]int{}
	}
	result := [][]int{{1}}
	if row == 1 {
		return result
	}
	for i:= 1; i < row; i++ {
		tmpRow := make([]int, i + 1)
		tmpRow[0] = 1
		tmpRow[i] = 1
		for j := 1; j < i; j++ {
			tmpRow[j] = result[i - 1][j - 1] + result[i - 1][j]
		}
		result = append(result, tmpRow)
	}
	return result
}

//获取单行
func getRow(row int) []int {
	arr := [][]int{{1}}
	var nowRow []int
	if row == 0 {
		return []int{}
	}
	if row == 1 {
		return []int{1}
	}
	for i:= 1; i < row; i++ {
		tmpRow := make([]int, i + 1)
		tmpRow[0] = 1
		tmpRow[i] = 1
		for j := 1; j < i; j++ {
			tmpRow[j] = arr[i - 1][j - 1] + arr[i - 1][j]
		}
		arr = append(arr, tmpRow)
		nowRow = tmpRow
	}
	return nowRow
}

func main()  {
	fmt.Println(getRow(0))
}
//时间复杂度：O(log(n))
//空间复杂度：O(1)
