package _removeDuplicatesFromSortedArray

import (
	"strings"
	"testing"
)
//给定两个字符串 s 和 t，判断它们是否是同构的。
//
//如果 s 中的字符可以被替换得到 t ，那么这两个字符串是同构的。
//
//所有出现的字符都必须用另一个字符替换，同时保留字符的顺序。两个字符不能映射到同一个字符上，但字符可以映射自己本身。
//
//示例 1:
//
//输入: s = "egg", t = "add"
//输出: true
//示例 2:
//
//输入: s = "foo", t = "bar"
//输出: false
//示例 3:
//
//输入: s = "paper", t = "title"
//输出: true
//说明:
//你可以假设 s 和 t 具有相同的长度。


func TestHello(t *testing.T) {
	s := "paper"
	a := "title"

	t.Log(isIsomorphic(s, a))
}
//时间最优
//国外1
func isIsomorphic1(s string, t string) bool {

	//  对应位置字母的数量相同，可以通过构建数组 hash 表
	if len(s) != len(t) {
		return false
	}

	m1 := make([]int, 256)
	m2 := make([]int, 256)

	for i := 0; i < len(s); i++ {
		if m1[int(s[i])] != m2[int(t[i])] {
			return false
		}
		m1[int(s[i])] = i + 1
		m2[int(t[i])] = i + 1
	}
	return true
	// if len(s) != len(t) {
	// 	return false
	// }
	// bs := strings.Split(s, "")
	// bt := strings.Split(t, "")

	// if isPatter(bs, bt) && isPatter(bt, bs) {
	// 	return true
	// }
	// return false

}

// func isPatter(bs, bt []string) bool {
// 	size := len(bs)
// 	pattern := make(map[string]string, size)
// 	//  构建 map，如果构建成功
// 	for i := 0; i < size; i++ {
// 		if w, ok := pattern[bs[i]]; ok {
// 			if w != bt[i] {
// 				return false
// 			}
// 		} else {
// 			pattern[bs[i]] = bt[i]
// 		}
// 	}
// 	return true
// }
//国外2
func isIsomorphic2(s string, t string) bool {
	n := len(s)
	if n == 0 {
		return true
	}

	mp1 := make(map[byte]byte)
	mp2 := make(map[byte]bool)

	p := 0
	for p < n {

		if v,ok := mp1[s[p]]; ok {
			if v == t[p] {
				p++
			} else {
				return false
			}
		} else {
			if _,ok := mp2[t[p]]; ok {
				return false
			}
			mp1[s[p]] = t[p]
			mp2[t[p]] = true
			p++
		}

	}
	return true
}
//国内3
func isIsomorphic3(s string, t string) bool {
	slen := len(s)
	for i:= 0;i < slen;i++ {
		if strings.IndexByte(s, s[i]) != strings.IndexByte(t, t[i]) {
			return false
		}
	}
	return true
}
//国内4
//内存最优
//国外5
func isIsomorphic5(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	isomap := make(map[string]string)
	taken := make(map[string]bool)
	for i:=0 ; i < len(s); i++ {
		if v, ok := isomap[string(s[i])]; ok {
			if string(t[i]) != v {
				return false
			}
		} else {
			if _, ok := taken[string(t[i])]; ok {
				return false
			}
			isomap[string(s[i])] = string(t[i])
			taken[string(t[i])] = true
		}
	}
	return true
}
//国外6
func isIsomorphic6(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	n := len(s)
	to := map[byte]byte{}
	from := map[byte]byte{}
	for i := 0; i < n; i++ {
		si := s[i]
		ti := t[i]

		if ti2, ok := to[si]; ok && ti != ti2 {
			return false
		}
		if si2, ok := from[ti]; ok && si != si2 {
			return false
		}
		to[si] = ti
		from[ti] = si
	}
	return true
}

//国内7
//国内8
