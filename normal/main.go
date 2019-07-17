package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1,1,1,3,3,4,3,2,4,2}
	sort.Ints(a)
	fmt.Println(a)
}

//func a(s *big.Int) *big.Int {
//	if s.Int64() == 1{
//		return big.NewInt(1)
//	}else{
//		return s.Mul(s,a(big.NewInt(s.Int64()-1)))
//	}
//}
