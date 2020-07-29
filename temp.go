package main

import "fmt"

func findType(i interface{})  {
	switch v:= i.(type) {
	case interface{}:
		fmt.Println(v)
	}
}

func main()  {
	var firstName = "vinay"
	findType(firstName)
}