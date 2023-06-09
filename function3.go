package main

import "fmt"

func main(){
	fmt.Println(add(1,3,5,2,4,2,2,5,76,5))
}

func add(args ...int) int{

	total := 0

	for _, v := range args{
		total += v
	}

	return total
}