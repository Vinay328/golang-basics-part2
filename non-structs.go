package main

import (
	"fmt"
)

//TODO non struct type
type StringUtils string //TODO different type

//TODO Attach the method
func (su StringUtils) reverse() string  {
	var bytes = []rune(string(su)) // TODO you can change to uint8 as well but be careful
	i:=0
	j:= len(bytes)-1
	for i <= j {
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}
	return string(bytes)
}

func main()  {
	var name1 = StringUtils("vinay")
	fmt.Println(len(name1))
	fmt.Println(name1.reverse())
}