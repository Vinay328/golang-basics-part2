package main

import (
	"fmt"
)

type StringUtils string //TODO different type

func (su StringUtils) reverse() string  {
	var bytes = []rune(string(su))
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
	fmt.Println(name1.reverse())
}