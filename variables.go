package main

import "fmt"

func main(){

	var x string = "Hello World"

	fmt.Println(x)
	
	x = "Hello"
	var y string = "Hello"

	fmt.Println(x == y)
}