package main


import (
	"fmt"
	"strconv"
)

func main(){

	var number_input int
	var mod int

	fmt.Scanf("%d", &number_input)

	mod = number_input % 2

	switch mod {
		case 0 : fmt.Println(strconv.Itoa(number_input) + " mod 2 is zero")
		case 1 : fmt.Println(strconv.Itoa(number_input) + " mod 2 is one")
		default: fmt.Println("Unknown Number")
	}
}