package main

import "fmt"

func main(){
	var mystr string
	var mystr2 string

	mystr  = "Três pratos de trigo para três tigres tristes"
	mystr2 = `Um ninho de mafagafos tinha sete mafagafinhos` 

	fmt.Println("Minha string 1 ", mystr)

	fmt.Println("Minha string 1 ", mystr2)

	fmt.Println("clá mundo"[0])

	mystr = "Hello " + "World"

	fmt.Println(mystr)



}