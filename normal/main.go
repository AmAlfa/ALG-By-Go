package main

import (
	"fmt"
	"math/big"
)

func main() {
	var i int64
	for i=130;i > 0;i--{
		fmt.Printf("%v => %v \n", i, a(big.NewInt(i)))
	}
}

func a(s *big.Int) *big.Int {
	if s.Int64() == 1{
		return big.NewInt(1)
	}else{
		return s.Mul(s,a(big.NewInt(s.Int64()-1)))
	}
}
