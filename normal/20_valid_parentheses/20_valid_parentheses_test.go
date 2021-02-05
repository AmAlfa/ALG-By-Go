package validParentheses

import (
	"fmt"
	"testing"
)

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
//
// 有效字符串需满足：
//
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 注意空字符串可被认为是有效字符串。
//
// 示例 1:
//
// 输入: "()"
// 输出: true
// 示例 2:
//
// 输入: "()[]{}"
// 输出: true
// 示例 3:
//
// 输入: "(]"
// 输出: false
// 示例 4:
//
// 输入: "([)]"
// 输出: false
// 示例 5:
//
// 输入: "{[]}"
// 输出: true



func Test(t *testing.T) {
	s := "()"
	res := isValid(s)
	fmt.Println(res)
	t.Log(res)

}

func isValid(s string) bool {
	var stack []byte
	var bracketMap = map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	for i := range s {
		if _, ok := bracketMap[s[i]]; !ok {
			stack = append(stack, s[i])
			continue
		}

		if len(stack) == 0 {
			return false
		}

		if stack[len(stack)-1] != bracketMap[s[i]] {
			return false
		}
		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}

// func isValid(s string) bool {
// 	end := len(s) - 1
// 	var queue = []string{"*"}
// 	var twain = map[string]string{")" : "(", "}" : "{", "]" : "["}
// 	for i := 0; i <= end; i++ {
// 		if string(s[i]) == "(" || string(s[i]) == "{" || string(s[i]) == "[" {
// 			queue = append(queue, string(s[i]))
// 		} else if string(s[i]) == ")" || string(s[i]) == "}" || string(s[i]) == "]" {
// 			qLength := len(queue) - 1
// 			if len(queue) == 1 {
// 				return false
// 			} else {
// 				if queue[qLength] == twain[string(s[i])] {
// 					queue = queue[:qLength]
// 				} else {
// 					return false
// 				}
// 			}
// 		}
//
// 		continue
// 	}
//
// 	if len(queue) == 1 {
// 		return true
// 	} else {
// 		return false
// 	}
// }
