package romanToInteger

import (
	"fmt"
	"testing"
)

// 罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。
//
// 字符          数值
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000
// 例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。
//
// 通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：
//
// I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
// X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。 
// C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
// 给定一个罗马数字，将其转换成整数。输入确保在 1 到 3999 的范围内。
//
// 示例 1:
//
// 输入: "III"
// 输出: 3
// 示例 2:
//
// 输入: "IV"
// 输出: 4
// 示例 3:
//
// 输入: "IX"
// 输出: 9
// 示例 4:
//
// 输入: "LVIII"
// 输出: 58
// 解释: L = 50, V= 5, III = 3.
// 示例 5:
//
// 输入: "MCMXCIV"
// 输出: 1994
// 解释: M = 1000, CM = 900, XC = 90, IV = 4.



func Test(t *testing.T) {
	s := "MCMXCIV"
	res := romanToInt(s)
	fmt.Println(res)
	t.Log(res)

}

// 最机智的做法
func romanToInt(s string) int {
	char2num := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000}
	sum := 0
	prev := 0
	for i := len(s) - 1; i >= 0; i-- {
		n := char2num[s[i]]
		if n < prev {
			sum -= n
		} else {
			sum += n
		}
		prev = n
	}
	return sum
}

// func romanToInt(s string) int {
// 	x := 0
// 	for i:=0; i < len(s); i++ {
// 		n := 0
// 		switch s[i] {
// 		case 'I': n = 1
// 		case 'V': n = 5
// 		case 'X': n = 10
// 		case 'L': n = 50
// 		case 'C': n = 100
// 		case 'D': n = 500
// 		case 'M': n = 1000
// 		}
// 		if (i+1 < len(s)) &&
// 				((s[i] == 'I' && s[i+1] == 'V') || (s[i] == 'I' && s[i+1] == 'X') ||
// 						(s[i] == 'X' && s[i+1] == 'L') || (s[i] == 'X' && s[i+1] == 'C') ||
// 						(s[i] == 'C' && s[i+1] == 'D') || (s[i] == 'C' && s[i+1] == 'M')) {
// 			n = -n
// 		}
// 		x += n
// 	}
// 	return x
// }

//
// func romanToInt(s string) int {
// 	i := 0
// 	result := 0
// 	for i < len(s) {
// 		isSpecial := false
// 		if (string(s[i]) == "I" || string(s[i]) == "X" || string(s[i]) == "C") && i <= len(s) - 2 {
// 			switch s[i: i + 2] {
// 			case "CM":
// 				i += 2
// 				isSpecial = true
// 				result += 900
// 			case "CD":
// 				i += 2
// 				isSpecial = true
// 				result += 400
// 			case "XC":
// 				i += 2
// 				isSpecial = true
// 				result += 90
// 			case "XL":
// 				i += 2
// 				isSpecial = true
// 				result += 40
// 			case "IX":
// 				i += 2
// 				isSpecial = true
// 				result += 9
// 			case "IV":
// 				i += 2
// 				isSpecial = true
// 				result += 4
// 			}
// 		}
// 		if !isSpecial {
// 			switch string(s[i]) {
// 			case "M":
// 				result += 1000
// 			case "D":
// 				result += 500
// 			case "C":
// 				result += 100
// 			case "L":
// 				result += 50
// 			case "X":
// 				result += 10
// 			case "V":
// 				result += 5
// 			case "I":
// 				result += 1
// 			}
// 			i += 1
// 		}
// 	}
//
// 	return result
// }
