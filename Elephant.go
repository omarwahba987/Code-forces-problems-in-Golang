package main

import "fmt"

func main() {
	var in int64
	fmt.Scan(&in)
	var res int64
	res=in/5
	if in%5 !=0{
		in-=(in/5)*5
		res+=(in/4)
		if in%4 !=0{
			in-=(in/4)*4
			res+=(in/3)
			if in%3 !=0{
				in-=(in/3)*3
				res+=(in/2)
				if in%2 !=0 {
					in-=(in/2)*2
					res+=in

				}
			}
		}
	}
	fmt.Println(res)
}
