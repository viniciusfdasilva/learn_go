package main

import "fmt"

func main(){

	values := map[string]string{
		"H"  : "Hydrogen",
		"He" : "Helium",
		"Li" : "Lithium",
		"Be" : "Beryllium",
		"B"  : "Boron",
		"C"  : "Carbon",
		"N"  : "Nitrogen",
		"O"  : "Oxygen",
		"F"  : "Fluorine",
		"Ne" : "Neon",
	}

	for key, value := range values{
		fmt.Println(key,"-", value)
	}
}