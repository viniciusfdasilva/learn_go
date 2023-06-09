package main

import (
	"fmt"
	"math"
)

func main(){
	if f2() == math.Pi { fmt.Println("Sim") }else{ fmt.Println("Não") }
	fmt.Println(f3())
}

func f2() (r float64){
	r = math.Pi
	return
}

func f3() (int, int){
	return 4, 2
}