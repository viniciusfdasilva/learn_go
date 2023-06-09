package main


import "fmt"

func main(){

	var i int64

	fmt.Scanf("%d", &i)

	fmt.Println(i)

	if i % 2 == 0 {
		fmt.Println("O valor é par!")
	}else{
		fmt.Println("O valor é ímpar!")
	}

	if i % 2 == 0{
		fmt.Println("O valor é divisível por 2")
	}else if i % 3 == 0{
		fmt.Println("O valor é divisível por 3")
	}else if i % 4 == 0{
		fmt.Println("O valor é divisível por 4")
	}
}