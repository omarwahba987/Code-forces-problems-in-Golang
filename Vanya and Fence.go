package main

import (
	"fmt"
)

func main() {
	var r int
	var i int
	fmt.Scan(&i)
	var x int
	fmt.Scan(&x)
	for  q:=0;q<i ;q++  {
		var z int
		fmt.Scan(&z)

		r += z/x
		if (z%x) >0{
			r++
		}
	}

	fmt.Println(r)



}
