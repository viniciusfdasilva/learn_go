package main


import "fmt"

func main(){
	for i := 1; i < 10; i++{
		if i % 2 == 0{
			fmt.Println("O valor é par")
		}else{
			fmt.Println("O valor é ímpar")
		}
	}
}