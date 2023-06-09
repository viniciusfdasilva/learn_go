package main

import "fmt"

func main(){
	var x [5]int
	
	x[4] = 100

	var i int = 0

	for ; i < len(x); i++{
		x[i] = i
		fmt.Println(x)
	}
	
}