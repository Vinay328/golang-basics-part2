package main

//TODO slice is pass by value
func main()  {

	//TODO example 1
	/*
	var slice1 = []int {1,2,3,4}
	var slice2 = slice1

	fmt.Println("length of slice1 is ", len(slice1))
	fmt.Println("length of slice2 is", len(slice2))

	slice1 = append(slice1, 5) // TODO length plus 1

	fmt.Println("length of slice1 is", len(slice1))
	fmt.Println("capacity of slice1 is",cap(slice1))
	fmt.Println("length of slice2 is", len(slice2))
	fmt.Println("capacity of slice2 is", cap(slice2))

	slice2[0] = 45
	slice2[1] = 45
	slice2[2] = 45
	slice2[3] = 45
	fmt.Println(slice1) // TODO [1,2,3,4,5]
	fmt.Println(slice2) // TODO [45, 45, 45, 45]TODO capacity exceeded
	*/

	//TODO example2
	/*var slice1 = make([]int, 4 , 8)
	var slice2 = slice1

	slice1 = append(slice1, 23)

	//TODO comment code in this demo
	slice1 = append(slice1, 33)
	slice1 = append(slice1, 43)
	slice1 = append(slice1, 53)
	slice1 = append(slice1, 63)

	slice2[0]  = 45
	fmt.Println(slice1) //TODO [45,0,0,0,23] -> capacity not exceeded
	fmt.Println(cap(slice1))*/


	//TODO example3
	/*var slice1 = []int {1,2,3,4}
	var slice2 = slice1 //TODO length and capacity is 4

	//TODO all these are eventually cleaned up in memory
	fmt.Println(append(slice1, 5))
	fmt.Println(append(slice1, 5))
	fmt.Println(append(slice1, 5))
	fmt.Println(append(slice1, 5))

	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(len(slice2))

	slice2[0] = 45
	fmt.Println(slice1[0] == slice2[0])*/

	/* TODO eventhough pass by value memory location is different
	var array1 = []int{2, 3}
	fmt.Println(&array1[0] == &array1[1]) // TODO different address*/

}