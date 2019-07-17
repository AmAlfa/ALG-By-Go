package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := "ahsj"
	b := "ahsj"
	fmt.Println(a == b)
}

func a(s *big.Int) *big.Int {
	if s.Int64() == 1{
		return big.NewInt(1)
	}else{
		return s.Mul(s,a(big.NewInt(s.Int64()-1)))
	}
}
