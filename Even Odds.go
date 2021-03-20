package main

import (
	"fmt"
	"math"
)

func main() {
	var in int64
	fmt.Scan(&in)
	var position int64
	fmt.Scan(&position)
	var res int64

	if in%2==0{
		if position<=in/2{
			fmt.Println((position*2)-1)
		}else {
			x:=math.Ceil(float64(in)/2.0)
			fmt.Println((position-int64(x))*2)
		}
	}else {
		if position-1<=in/2{
			res=(position*2)-1
			fmt.Println(res)
		}else {
			x:=math.Ceil(float64(in)/2.0)
			res=(position-int64(x))*2
			fmt.Println(res)
		}
	}

}
