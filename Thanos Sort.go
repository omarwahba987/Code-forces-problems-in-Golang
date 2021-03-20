package main

import (
	"fmt"
)

func isSorted(arr []int )bool  {
	for i:=0;i<len(arr)-1 ;i++  {
		if arr[i]>arr[i+1]{
			return false
		}else {
			continue
		}
	}
	return true
}
var lent =0
func thanos (arr []int){
	if isSorted(arr){
		if len(arr)>lent{
			lent=len(arr)
		}
		return
	}
	left:=arr[:(len(arr)/2)]
	right:=arr[(len(arr)/2):]
		thanos(left)
	thanos(right)
}
func main() {
	var size int
	var arr []int
	fmt.Scan(&size)
	var in int
	for i:=0;i<size ;i++  {
			fmt.Scan(&in)
			arr =append(arr,in)
	}
	thanos(arr)
	fmt.Println(lent)

}
