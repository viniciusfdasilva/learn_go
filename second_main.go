package main

import "fmt"

var teste string

func main(){
	fmt.Println("1 + 1 =", 1 + 1)
	fmt.Println("1 + 1 =", 1.0 + 1.0)

	teste = `new string`
	teste = "new string"
	
	fmt.Println(teste)

}