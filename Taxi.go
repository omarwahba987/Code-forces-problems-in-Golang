package main

import (
	"fmt"
	"sort"
)

func main() {
	var groupsNumber int
	var groups []int
	fmt.Scan(&groupsNumber)
	var in int
	for i:=0;i<groupsNumber ;i++  {
		fmt.Scan(&in)
		groups =append(groups,in)
	}
	sort.Ints(groups)
	fmt.Println("dfgf",groups)
}
