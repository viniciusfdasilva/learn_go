package main

import (
	"fmt"
	//"math"
	//"strconv"
)

func main(){

	maps := make(map[string]int)

	maps["key0"] = 12
	maps["key1"] = 12
	maps["key2"] = 12132
	maps["key6"] = 12
	maps["key3"] = 1223
	maps["key4"] = 12243
	maps["key5"] = 12423
	maps["key6"] = 22
	maps["key7"] = 11
	maps["key8"] = 3
	maps["key9"] = -2

	//var i int = 0

	//fmt.Println(math.Pi)

	//fmt.Println(math.E)


	name, ok := maps["key0"]

	fmt.Println(name, ok)


	if name, ok := maps["key0"]; ok{
		fmt.Println(name, ok)
	}

	//for _, _map := range maps{
	//	fmt.Println(_map)
	//	i++
	//}


}