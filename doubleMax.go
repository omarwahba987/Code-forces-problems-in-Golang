package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)
func TreeConstructor(strArr []string) string {
	magicNumber := float64(len(strArr))/2.0
	flip:= int(math.Ceil(magicNumber))
	for i:=0;i<flip;i++{
		value:=strArr[i]
		if (value[1])>(value[3]){
			return "false"
		}
	}
	for i:=flip;i<len(strArr);i++{
		value:=strArr[i]
		if (value[1])<(value[3]){
			return "false"
		}
	}
	return "true"

}
func SumMultiplier(arr []int) string {
	doubleSum := 0
	for _,arrValue := range arr{
		doubleSum += arrValue*2
	}
	sort.Ints(arr)
	arrMaxValue := arr[len(arr)-1] *arr[len(arr)-2]
	if arrMaxValue > doubleSum{
		return "true"
	}else {
		return "false"
	}

}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(SumMultiplier(readline()))

}