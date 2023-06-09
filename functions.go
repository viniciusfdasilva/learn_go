package main

import "fmt"

func main(){

	average(make([]float64, 5))
}

func average(xs []float64) float32{
	panic("Not Implemented")
	return 2.2
}

func panic(str string){
	fmt.Println(str)
}