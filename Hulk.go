package main

import "fmt"

func main() {
	h1:="I hate it "
	h2:="I hate that "
	l1:="I love it "
	l2:="I love that "

	var size int
	fmt.Scan(&size)

	flag:=true
	for size>0  {
		if size ==1 &&flag==true{
			fmt.Print(h1)
			size--
		}else if size ==1 &&flag==false{
			fmt.Print(l1)
			size--
		}else {
			if flag ==true{
				fmt.Print(h2)
				flag=false
				size--
			}else {
				fmt.Print(l2)
				flag=true
				size--
			}
		}

	}





}
