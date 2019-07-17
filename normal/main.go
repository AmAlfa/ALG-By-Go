package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{2,6,4,8,10,9,15}
	//1 3 21 25 36 44 321
	fmt.Println(s)
	numsTmp := make([]int, len(s))
	copy(numsTmp, s)
	sort.Ints(numsTmp)
	fmt.Println(numsTmp)
	result := 0
	for i := 0; i < len(s); i++ {
		if numsTmp[i] != s[i] {
			result += 1
		}
	}
	fmt.Println(result)
}
