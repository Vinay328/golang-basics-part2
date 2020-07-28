
package main

import (
	"fmt"
)

func customPrinter(a ...interface{})  {
	for _,value:= range a {
		fmt.Printf("Concrete type is %T and value is %v\n \n", value, value) //TODO if false zero value will be printed
	}
}

type Empty interface {

}

type Animal struct {
	petName string
}

type Human struct {
	firstName string
}

func main()  {

	customPrinter(23, "vinay", "false")

	/*
	var value interface{} = "vinay"
	fmt.Println(value)
	*/

	//TODO error example
	/*var animal1 Empty = Animal{"heda"}
	var animal2 = animal1.(Human)
	fmt.Println(animal2)*/

}