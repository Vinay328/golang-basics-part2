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

//TODO update Area field using person struct
func (p *Person )setStreeTName()  {
	p.streetName = "B T M"
}

func (p Area)getStreetName() string {
	fmt.Println(reflect.TypeOf(p))
	return p.streetName
}

func main() {

	//TODO initialization
	var personObject = Person{
		firstName: "vinay",
		lastName:  "K",
		Area : Area{
			streetName: "H S R layout",
			No:         75,
		},
	}

	//TODO pass by value
	/*personObject.setFirstName("dota")
	fmt.Println(personObject)*/

	//TODO to find the data type
   /* fmt.Println("Person type is", reflect.TypeOf(personObject))
	fmt.Println("Area type is", reflect.TypeOf(personObject.Area))*/

   //TODO additional methods
	fmt.Println(personObject.getFullname())

	personObject.getStreetName() // TODO IMP just like promoted fields you can access promoted methods as well

	//TODO can achieve constructors, getters and setters


}
