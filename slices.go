package main

import "fmt"

func main(){

	//var x []float64


	//x := make([]float64, 5)

	//arr := [5]float64{1,2,3,4,5}
	//x := arr[0:5]

	slice1 := []int{1,2,3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)


}