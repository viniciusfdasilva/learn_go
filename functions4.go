package main

import "fmt"


func main(){
	//xs := []int{1,2,3}

	add := func(x, y int) int{
		return x + y
	}

	fmt.Println(add(1,1))
	//add(xs...)
}

//func add(xs []int){
//	fmt.Println(xs)
//}

//func println(a ...interface{}) (n int, err error){ }