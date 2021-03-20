package main

import "fmt"

func main() {
	var in int64
	fmt.Scan(&in)
	if in%2==0{
		fmt.Print(in/2)
	}else {
		fmt.Println(((in-1)/2)-in)
	}
}
