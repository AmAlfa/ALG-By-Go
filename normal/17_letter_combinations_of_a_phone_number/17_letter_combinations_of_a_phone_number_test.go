package romanToInteger

import (
	"testing"
)

// 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
//
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
//
//
//
// 示例:
//
// 输入："23"
// 输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
// 说明:
// 尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。


func Test(t *testing.T) {
	s := "237"
	res := letterCombinations(s)
	// fmt.Println(res)
	t.Log(res)

}

var mm []string = []string{
	"abc",
	"def",
	"ghi",
	"jkl",
	"mno",
	"pqrs",
	"tuv",
	"wxyz",
}

func letterCombinations(digits string) []string {
	if len(digits) == 0{
		return []string{}
	}
	res := make([]string,0)
	str := ""
	i := 0
	helper(digits,str,i,&res)
	return res
}


func helper(digits,str string,i int, res *[]string){
	if i == len(digits) {
		*res = append(*res,str)
		return
	}else {
		s := mm[digits[i] - '2']
		for _,ch := range s{
			helper(digits,str + string(ch),i+1,res)
		}
	}
}

//
// func letterCombinations(digits string) []string {
// 	m := map[int][]string{
// 		2 : {"a", "b", "c"},
// 		3 : {"d", "e", "f"},
// 		4 : {"g", "h", "i"},
// 		5 : {"j", "k", "l"},
// 		6 : {"m", "n", "o"},
// 		7 : {"p", "q", "r", "s"},
// 		8 : {"t", "u", "v"},
// 		9 : {"w", "x", "y", "z"},
// 	}
// 	var result []string
// 	for _, v := range digits {
// 		var tmp []string
// 		if result == nil {
// 			result = m[int(v - '0')]
// 			continue
// 		}
// 		for _, char := range m[int(v - '0')] {
// 			for _, value := range result{
// 				fmt.Println(value)
// 				tmp = append(tmp, value + char)
// 			}
// 		}
// 		result = tmp
// 	}
//
// 	return result
// }
