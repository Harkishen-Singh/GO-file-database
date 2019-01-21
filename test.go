package main

import ("fmt")

func main() {
	var vals  = []int{1,2,3,1,0}
	fmt.Println(vals)
	var b = &vals
	fmt.Println(&vals[0])
	fmt.Println("address way ", *b)
}