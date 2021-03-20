package main

import (
	"fmt"
)

var x =10
func f()int{
	return x
}
func g()int{
	x:=20
	fmt.Println("g",x)
	return f()
}
func main() {
	fmt.Println(g())
}
