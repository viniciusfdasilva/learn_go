package main

import "fmt"

func main(){

	x := make(map[string]int)

	x["key"] = 10
	x["val"] = 10010010100

	fmt.Println(x["val"])

	delete(x, 1)

	fmt.Println(x)
}