package str_print_reverse

import (
	"testing"
)

func TestStrPrintReverse(t *testing.T)  {
	slice := []int64{93, 213, 1, 443, 231, 146, 72}
	//[213 93 443 231 146 72 1]
	for len := len(slice) - 1; len > 0; len--  {
		t.Errorf("%v\n", slice[len])
	}
}
//时间复杂度：O(n)
//空间复杂度：O(1)
