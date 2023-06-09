package main

import "fmt"

func main(){
	x := []float64{90,23,23,1,4,54,2}
	
	fmt.Println(average(x))

	fmt.Println(average([]float64{90,23,23,1,4,54,2}))
	//fmt.Println(average(make([]float64, 4)))
}

func average(x []float64) float64{
	total := 0.0

	for _, v := range x{
		total += v
	}

	return total / float64(len(x))

}