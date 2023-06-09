package main

import "fmt"

func main(){

	var total float64 = 0

	//var value float64
	values := [5]float64{23,55,4,34,32}

	for _, value := range values{
		total += value
	}

	fmt.Println(total  / float64(len(values)))
}