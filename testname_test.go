package main

import (
    "testing"
)

func f(p *int) int {
    *p = 123
    return *p
}

func foo() int {
    var x int
    y, _ := x, f(&x)
    return y
}

func bar() int {
    var x int
    var y, _ = x, f(&x)
    return y
}

func Test(t *testing.T) {
    println(foo(), bar())
}
