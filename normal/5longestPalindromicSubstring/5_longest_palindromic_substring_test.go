package longestPalindromicSubstring

import (
	"fmt"
	"testing"
)

// 给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
//
// 示例 1：
//
// 输入: "babad"
// 输出: "bab"
// 注意: "aba" 也是一个有效答案。
// 示例 2：
//
// 输入: "cbbd"
// 输出: "bb"


func Test(t *testing.T) {
	s := ""
	// for key, val := range s {
	// 	t.Log(key, val)
	// }
	res := longestPalindrome(s)
	fmt.Println(res)
	t.Log(res)

}

// 最优解：
func longestPalindrome(s string) string {
	ll := len(s)
	if ll == 0 {
		return ""
	}

	var l, r, pl, pr int
	for r < ll {
		// gobble up dup chars
		for r+1 < ll && s[l] == s[r+1] {
			r++
		}
		// find size of this palindrome
		for l-1 >= 0 && r+1 < ll && s[l-1] == s[r+1] {
			l--
			r++
		}
		if r-l > pr-pl {
			pl, pr = l, r
		}
		// reset to next mid point
		l = (l+r)/2 + 1
		r = l
	}
	return s[pl : pr+1]
}


// 最优解：
// func longestPalindrome(s string) string {
//	if s == "" {
//		return ""
//	}
//
//	start, end := 0, 0
//	if len(s)%2 == 0 {
//		//偶数
//		center := len(s) / 2
//		for i := 0; i <= center-1; i++ {
//			right := center + i
//			left := center - i - 1
//
//			//left two-elem center
//			length := expandCenter(s, left, left+1)
//			if length > end-start {
//				start = left - (length-1)/2
//				end = left + length/2
//				if length == (left*2 + 2) {
//					return s[start : end+1]
//				}
//			}
//
//			//left single center
//			length = expandCenter(s, left, left)
//			if length > end-start {
//				start = left - (length-1)/2
//				end = left + length/2
//				if length == left*2+1 {
//					//longest
//					return s[start : end+1]
//				}
//			}
//
//			//right single center
//			length = expandCenter(s, right, right)
//			if length > end-start {
//				start = right - (length-1)/2
//				end = right + length/2
//				if length == (len(s)-right)*2-1 {
//					//longest
//					return s[start : end+1]
//				}
//			}
//
//			//right two-elem center
//			length = expandCenter(s, right, right+1)
//			if length > end-start {
//				start = right - (length-1)/2
//				end = right + length/2
//				if length == (len(s)-right)*2-2 {
//					//longest
//					return s[start : end+1]
//				}
//			}
//		}
//	} else {
//		//奇数
//		center := len(s) / 2
//		for i := 0; i <= center-1; i++ {
//			right := center + i
//			left := center - 1 - i
//
//			//right single center
//			length := expandCenter(s, right, right)
//			if length > end-start {
//				start = right - (length-1)/2
//				end = right + length/2
//				if length == (len(s)-right)*2-1 {
//					//longest
//					return s[start : end+1]
//				}
//			}
//			//right two-elem center
//			length = expandCenter(s, right, right+1)
//			if length > end-start {
//				start = right - (length-1)/2
//				end = right + length/2
//				if length == (len(s)-right)*2-2 {
//					//longest
//					return s[start : end+1]
//				}
//			}
//
//			//left two-elem center
//			length = expandCenter(s, left, left+1)
//			if length > end-start {
//				start = left - (length-1)/2
//				end = left + length/2
//				if length == (left*2 + 2) {
//					return s[start : end+1]
//				}
//			}
//
//			//left single center
//			length = expandCenter(s, left, left)
//			if length > end-start {
//				start = left - (length-1)/2
//				end = left + length/2
//				if length == left*2+1 {
//					//longest
//					return s[start : end+1]
//				}
//			}
//		}
//	}
//
//	return s[start : end+1]
// }
//
// //解法4 中心扩展
// // func longestPalindrome(s string) string {
// //     if s == ""{
// //         return ""
// //     }
// //     start,end := 0,0
// //     for i:=0;i<len(s);i++{
// //         len1 := expandCenter(s,i,i)
// //         len2 := expandCenter(s,i,i+1)
// //         if len1 < len2{
// //             len1 = len2
// //         }
// //         if len1 > end-start{
// //             start = i - (len1-1)/2
// //             end = i+ len1/2
// //         }
// //     }
// //     return s[start:end+1]
// // }
//
// func expandCenter(s string, left, right int) int{
//    for ;left>=0&&right<len(s)&&s[left]==s[right];{
//        left--
//        right++
//    }
//    return right-left-1
// }
//
// //解法3 动态规划，从最小开始，遍历所有子串，并判断回文
// //存储子串的回文判断结果，加速回文判断
// //时间 O(n^2)
// //空间 O(n)
// // func longestPalindrome(s string) string {
// // 	length := len(s)
// // 	isP := make([]bool,length)
// // 	max := 0;
// // 	maxStr := ""
// // 	for i := length-1; i>=0 ; i--{
// // 		for j := length-1; j >= i ; j-- {
// // 			isP[j] = (j-i<3 || isP[j-1]) && s[i] == s[j]
// // 				if isP[j] && j-i+1 > max{
// //                     max = j-i+1
// // 					maxStr = s[i:j+1]
// //                 }
// // 		}
// // 	}
// // 	return maxStr
// // }
//
// // func longestPalindrome(s string) string {
// // 	length := len(s)
// // 	isP := make([][]bool,length,length)
// //     for i:=0;i<length;i++{
// //         isP[i] = make([]bool,length)
// //     }
// // 	max := 0;
// // 	maxStr := ""
// // 	for curLen := 1; curLen <= length ; curLen++ {
// // 		for start := 0; start <= length-curLen; start++ {
// // 			end := start+curLen-1;
// // 			isP[start][end] = (curLen==1 || curLen == 2 || isP[start+1][end-1]) && s[start] == s[end]
//
// // 				if isP[start][end] && curLen > max{
// //                     max = curLen
// // 					maxStr = s[start:end+1]
// //                 }
// // 		}
// // 	}
// // 	return maxStr
// // }
//
// //解法2 最长公共子串
// //时间 O(n^2)
// //空间 O(n)
// // func longestPalindrome(s1 string) string{
// //     if s1==""{
// //         return ""
// //     }
//
// //     s2 := reverseString(s1)
// // 	max:=0
// // 	endIndex := 0
// // 	record:=make([]int,len(s2))
//
// // 	for i :=0; i <len(s1); i++{
// // 		for j :=len(s2)-1; j >= 0; j--{
// // 			if s1[i]==s2[j]{
// // 				if i == 0 || j == 0{
// // 					record[j] = 1
// // 				}else{
// // 					record[j]=record[j-1]+1
// // 				}
// // 				if record[j] > max{
// //                     if len(s1)-j-1+record[j]-1 == i{
// // 					    max = record[j];
// // 					    endIndex = i;
// //                     }
// // 				}
// // 			}else{
// // 				record[j] = 0
// // 			}
// // 		}
// // 	}
// // 	return s1[endIndex-max+1:endIndex+1];
// // }
//
// // func reverseString(s string) string {
// //     runes := []rune(s)
// //     for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
// //         runes[from], runes[to] = runes[to], runes[from]
// //     }
// //     return string(runes)
// // }
//
// //解法1.1 暴力搜索改进，从最长子串开始搜索，一旦判断是回文，可以立刻返回，减少无用遍历
// //时间 O(n^3), 最差情况才是 O(n^3)
// //空间 O(1)
// // func longestPalindrome(s string) string {
// // 	for end := len(s); end >=0 ; end-- {
// // 		for start,tempEnd := 0,end; tempEnd <= len(s); tempEnd++ {
// // 			if isPalindrome(s[start:tempEnd]) {
// // 				return s[start:tempEnd]
// // 			}
// // 			start++
// // 		}
// // 	}
// // 	return ""
// // }
//
//
// //解法1 暴力搜索
// //时间 O(n^3)
// //空间 O(1)
// // func longestPalindrome(s string) string {
// // 	var max int = 0
// // 	var result string = ""
// // 	for i := 0; i < len(s); i++ {
// // 		for j := i + 1; j <= len(s); j++ {
// // 			if isPalindrome(s[i:j]) {
// // 				if j-i > max {
// // 					max = j - i;
// // 					result = s[i:j]
// // 				}
// // 			}
// // 		}
// // 	}
// // 	return result
// // }
//
// func isPalindrome(s string) bool {
//	for i := 0; i < len(s)/2; i++ {
//		if s[i] != s[len(s)-i-1] {
//			return false
//		}
//	}
//	return true
// }


// 法五：
// func longestPalindrome(s string) string {
// 	maxL:=0
// 	result:=""
// 	if len(s)==1 {
// 		return string(s[0])
// 	} else if len(s)==0 {
// 		return ""
// 	}
// 	result=string(s[0])
// 	maxL=1
// 	for i:=0;i<len(s);i++{
// 		l1,r1:=calcByStartLength(i,0,s)
// 		r2:=""
// 		l2:=0
// 		if i<len(s)-1 && s[i]==s[i+1]{
// 			l2,r2=calcByStartLength(i,1,s)
// 		}
// 		tempL:=l1
// 		temp:=""
// 		if l1>l2{
// 			temp=r1
// 		} else {
// 			temp=r2
// 			tempL=l2
// 		}
// 		if tempL>maxL{
// 			maxL=tempL
// 			result=temp
// 		}
// 	}
// 	return result
// }
//
// func calcByStartLength(start int,sl int, s string) (int, string) {
// 	end:= start+sl
// 	result:=end-start + 1
// 	for start>0 && end <len(s)-1{
// 		if s[start-1]==s[end+1]{
// 			start--
// 			end++
// 			result+=2
// 		} else {
// 			break
// 		}
// 	}
// 	return result, string(s[start : end+1])
// }

// 法四： 马拉车
// func longestPalindromicSubstring(s string) string {
// 	n:=len(s)
// 	if(n<2){
// 		return s
// 	}
// 	shaped_str:=string_shape(s)
// 	// 在字符间的空位插入井号，便于检索位置
// 	palinds:=make([]int,2*n+1)
// 	R,C:=-1,-1
// 	max_palind_length:=0
// 	posi:=0
// 	for i:=0 ; i < n*2+1 ; i++{
// 		palinds[i]=1
// 		if(R>i){
// 			palinds[i]=min(palinds[2*C-i],R-i+1)
// 		}
// 		for{
// 			if(i+palinds[i]>= 2*n+1 || i<palinds[i]){
// 				break
// 			}
// 			if(shaped_str[i-palinds[i]]==shaped_str[i+palinds[i]]){
// 				palinds[i]++
// 			}else{
// 				break
// 			}
//
// 		}
// 		if(i+palinds[i]>R){
// 			R=i+palinds[i]-1;
// 			C=i;
// 		}
// 		if(palinds[i]>max_palind_length){
// 			max_palind_length=palinds[i]
// 			posi=i
// 		}
// 	}
//
//
//
// 	if(posi%2==1){
// 		posi/=2
// 		max_palind_length=max_palind_length/2-1
// 		s=s[(posi-max_palind_length):(posi+max_palind_length+1)]
// 	}else{
// 		posi/=2
// 		max_palind_length/=2
// 		s=s[(posi-max_palind_length):(posi+max_palind_length)]
// 	}
//
//
// 	return s
// }
//
// func string_shape(s string)string{ // 这个函数用来在字符间的空位插入井号
// 	out_str:="#"
// 	for _, v := range s {
// 		out_str+=string(v)+"#"
// 	}
// 	return out_str
// }
//
// func min(x,y int) int {
// 	if x < y {
// 		return x
// 	}
//
// 	return y
// }

// 法三：
// 动态规划
// func longestPalindrome(s string) string {
//    length := len(s)
//
//    if length <= 1 {
//        return s
//    }
//
//    dp := make([][]bool, length)
//
//    start := 0
//    maxLen := 1
//
//    for r := 0;r<length;r++ {
//        dp[r] = make([]bool, length)
//        dp[r][r] = true
//        for l:=0;l<r;l++ {
//            if s[l] == s[r] && (r -l <= 2 || dp[l+1][r-1]) {
//                dp[l][r] = true
//            }else{
//                dp[l][r] = false
//            }
//
//            if dp[l][r] {
//                curLen := r-l+1
//                if curLen > maxLen {
//                    maxLen = curLen
//                    start = l
//                }
//            }
//        }
//    }
//    return s[start:start+maxLen]
// }

// 法二：
// // 最长公共字串 缺少对反串后的原索引是否跟原字符串索引相同比较
// func longestPalindromicSubstring(s string) string {
// 	length := len(s)
// 	res := ""
// 	maxLen := 0
// 	maxEnd := 0
//
// 	reverse := reverse(s, length)
// 	slice :=make([][]int, length)
// 	for i := 0; i < length; i++ {
// 		slice[i] = make([]int, length)
// 	}
// 	for i := 0; i < length; i++ {
// 		for j := 0; j < length; j++ {
// 			if string(s[j]) == string(reverse[i]) {
// 				if i == 0 || j == 0 {
// 					slice[i][j] = 1
// 				} else {
// 					slice[i][j] = slice[i - 1][j - 1] + 1
// 				}
// 				if maxLen < slice[i][j] {
// 					maxEnd = j
// 					maxLen = slice[i][j]
// 					fmt.Println("maxEnd:" + strconv.Itoa(maxEnd))
// 					fmt.Println("slice[i][j]:" + strconv.Itoa(slice[i][j]))
// 				}
// 			}
// 		}
// 	}
// 	fmt.Println(maxEnd)
// 	fmt.Println(slice)
// 	maxEnd = maxEnd + 1
// 	if maxLen > 1 {
// 		res = s[maxEnd - maxLen: maxEnd]
// 	}
//
// 	return res
// }
//
// func reverse(s string, len int) string {
// 	res := ""
// 	for x:= len - 1; x >= 0; x-- {
// 		res += string(s[x])
// 	}
//
// 	return res
// }

// 法一：
// Brute-force（BF算法）(暴力算法)破解
// func longestPalindromicSubstring1(s string) string {
// 	len := len(s)
// 	if len == 0 {
// 		return ""
// 	}
//
// 	maxLen := 0
// 	maxString := ""
// 	for i := 0; i < len; i++ {
// 		for j := 0; j < len; j++ {
// 			if isPalindromic1(s, i, j) {
// 				if maxLen < j - i + 1 {
// 					maxLen = j - i + 1
// 					maxString = s[i:j+1]
// 				}
// 			}
// 		}
// 	}
//
// 	return maxString
// }
//
// func isPalindromic1(s string, l, r int) bool {
// 	for l < r {
// 		if s[l] != s [r] {
// 			return false
// 		}
//
// 		l ++
// 		r --
// 	}
//
// 	return true
// }
