package main

import (
	"fmt"
	"reflect"
)

func main()  {

	var singleChar = 'i'
	fmt.Println(singleChar, "and data type is", reflect.TypeOf(singleChar))

	var name = "vinay"
	fmt.Println(name, "and data type is", reflect.TypeOf(name))
	fmt.Println(name[0])
	// name[0] = 's' //TODO cannot assign to name[0]



}
