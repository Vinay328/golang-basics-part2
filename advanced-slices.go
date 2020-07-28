package main

//TODO slice is pass by value
func main()  {

	//TODO example 1
	/*var slice1 = []int {1,2,3,4}
	var slice2 = slice1

	fmt.Println(len(slice1))
	fmt.Println(len(slice2))

	slice1 = append(slice1, 5)

	fmt.Println(len(slice1))
	fmt.Println(len(slice2))

	slice2[0] = 45
	fmt.Println(slice1[0], slice2[0])*/


	//TODO example2
	/*var slice1 = make([]int, 4 , 8)
	var slice2 = slice1

	slice1 = append(slice1, 23)
	slice2[0]  = 45
	fmt.Println(slice1[0], slice2[0])*/

	//TODO example3
	/*var slice1 = []int {1,2,3,4}
	var slice2 = slice1

	fmt.Println(len(slice1))
	fmt.Println(len(slice2))

	fmt.Println(append(slice1, 5))
	fmt.Println(append(slice1, 5))

	fmt.Println(slice1) // TODO o/p -> 12345
	fmt.Println(len(slice1))
	fmt.Println(len(slice2))

	slice2[0] = 45
	fmt.Println(slice1[0], slice2[0])*/

}