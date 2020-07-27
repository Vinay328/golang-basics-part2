//TODO user defined data type -> contains a collection of fields
//TODO JAVA background - you can think struct more of like a class
//TODO STRUCT is pass by value
//TODO type printed is main.Temp and main.Employee and for anonymous structs is struct { firstName string }
//TODO ->  empty struct -> {} and for fields, all the struct fields are initialized with there zero value
package main

import (
	"fmt"
	"reflect"
)

//TODO unexported field and struct too
//TODO nested struct
type Employee struct {
	firstName , lastName string
	Age                  int     //exported field
	addr                 Address // TODO unexported field and this embedded field too initialized
	string                       //TODO the name of an anonymous field is its type
}

type Address struct {
	streetName string
	No int
}

type Temp struct {
}

//TODO -> IMP structs are pass by value
func updateEmployeeObject(employee Employee)  {
	employee.firstName = "Felix"

}

//TODO pointer to struct
func updateEmployeeObjectUsingPointer(employee *Employee)  {
	(*employee).addr.streetName = "Sector 3 , H S R layout"
	//employee.addr.streetName = "Sector 3 , H S R layout"
}

func main()  {

	var tempObject Temp = Temp{}
	fmt.Println(tempObject, "and", tempObject == Temp{}, "and", reflect.TypeOf(tempObject))

	var employeeObject = Employee{
		firstName: "vinay",
		lastName: "K",
		addr: Address{
			streetName: "H S R layout",
			No: 75,
		},
	}

	//TODO individual field access
	/*var firstName = &employeeObject.firstName
	*firstName = "felix"
	fmt.Println(employeeObject.firstName)*/

	// TODO pointer reference
	/*updateEmployeeObjectUsingPointer(&employeeObject)
	fmt.Println("employee data is ", employeeObject)*/


	//TODO anonymous structs
	/*a:= struct {
		firstName string
	}{
		firstName: "anonymous",
	}
	fmt.Println(reflect.TypeOf(a))*/


	/*//TODO promoted fields (anonymous fields)
	employeeObject.No = 45
	fmt.Println(employeeObject.Address.No)*/



	//TODO create structs using in-built new function
	/*pointerObject := new(Employee) // TODO initialize with zero value and return pointer
	updateEmployeeObjectUsingPointer(pointerObject)
	fmt.Println(pointerObject)*/

	fmt.Println(employeeObject)


}

//TODO two structs are equal if all there corresponding fields are equal and comparable too
