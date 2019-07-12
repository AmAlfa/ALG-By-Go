package bubble_sort_test

import (
	"testing"
)

func TestBubbleSort(t *testing.T)  {
	slice := []int64{93, 213, 1, 443, 231, 146, 72}
	//[213 93 443 231 146 72 1]
	len := len(slice)
	for i := 0; i < 3; i++  {
		for j := 0; j < len - 1; j++  {
			if slice[j] < slice[j + 1] {
				slice[j], slice[j + 1] = slice[j + 1], slice[j]
			}
		}
	}
	t.Errorf("%v\n", slice)
}
//时间复杂度：O(n^2)
//空间复杂度：O(1)
