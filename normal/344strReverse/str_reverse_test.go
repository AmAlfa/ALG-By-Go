package str_reverse

import (
	"testing"
)

func TestStrReverse(t *testing.T) {
	slice := []string{"a", "daw", "wqeeq", "sdfs2", "dsfs", "1231", "ts", "ghrt4d", "h", "rb", "ji", "yky"}
	for i, len := 0, len(slice)-1; i < len; i++ {
		slice[i], slice[len] = slice[len], slice[i]
		len--
		t.Errorf("%v\n", len)
	}
	t.Errorf("%v\n", slice)
}
//时间复杂度：O(log(n))
//空间复杂度：O(1)
