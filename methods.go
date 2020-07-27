//TODO value receiver -> same like arrays -> pass by value
//TODO pointer receiver -> same like pointers -> pass by reference
//TODO for both value receiver and pointer receiver -> you can achieve the same using functions
package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	firstName, lastName string
	Age                 int //exported field
	Area                    // TODO unexported field and this embedded field too initialized
	string                  //TODO the name of an anonymous field is its type
}

type Area struct {
	streetName string
	No         int
}

func (p *Person) getFullname() string {
	return p.firstName + " " + p.lastName
}

func (p Person) setFirstName(firstName string) {
	p.firstName = firstName
}

func updatePersonObject(person Person) {
	person.firstName = "Felix"
}

func updatePersonObjectUsingPointer(person *Person) {
	(*person).Area.streetName = "Sector 3 , H S R layout"
	//employee.addr.streetName = "Sector 3 , H S R layout"
}

func (p Area)getStreetName() string {
	fmt.Println(reflect.TypeOf(p))
	return p.streetName
}

func main() {

	var personObject = Person{
		firstName: "vinay",
		lastName:  "K",
		Area : Area{
			streetName: "H S R layout",
			No:         75,
		},
	}

	//personObject.setFirstName("dota")

/*	fmt.Println("Person type is", reflect.TypeOf(personObject))
	fmt.Println("Area type is", reflect.TypeOf(personObject.Area))
	fmt.Println(personObject.getFullname())*/

fmt.Println( reflect.TypeOf(personObject.getFullname))

}
