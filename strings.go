//TODO strings are a collection of bytes
//TODO %s or %x is the format specifier to print a string to print bytes
// TODO %c to print characters of the string
package main

import (
	"fmt"
)

func main()  {

	var a = "simple"

	var b  =  []rune(a) // TODO very important Type conversion
	fmt.Println(string(b[0]))


}