package main

import (
	"fmt"
	"regexp"
)

func main()  {
	a := "A man, a plan, a canal: Panama"
	reg := regexp.MustCompile(`[^0-9a-zA-Z]`)
	fmt.Println(reg.ReplaceAllString(a, ""))
}
