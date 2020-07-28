//TODO zero value is nil
//TODO type assertion
//TODO type switch
package main

import (
	"fmt"
	"reflect"
)

//TODO interface
type TotalExpense interface {
	salary() int
}

//TODO contract employee
type permanent struct {
	employeeId int
	pay int
	pf int
}

//TODO consider permanent employee has more methods
func (p *permanent) benefits()  {
	fmt.Println("Congrats you got bonus for this year")
}

func (p permanent) salary() int {
	return p.pay + p.pf
}

//TODO contract employee
type contract struct {
	employeeId int
	pay int
}

func (c contract) salary() int {
	return c.pay
}

//TODO total expense for the company
func totalExpense(s *[]TotalExpense)  {
	amount := 0
	for _, v := range *s {
		amount += v.salary()

		/*//TODO type switch
		switch v.(type) {
		case permanent:
			var emp1 = v.(permanent)
			emp1.benefits()
		}*/

		//TODO type assertion if you are checking for only one type
		/*if value, ok := v.(permanent); ok {
			value.salary()
		}*/

	}
	fmt.Println("Total expense per month is ", amount)
}

func main()  {

	var pEmp1 TotalExpense = permanent{3802, 45000, 4500}
	fmt.Println(reflect.TypeOf(pEmp1))
	fmt.Printf("Concrete type is %T and value %v\n", pEmp1, pEmp1)
	pEmp2 := permanent{3803, 50000, 3500}
	cEmp1 := contract{3805, 30000}

	employees := []TotalExpense {pEmp1, pEmp2, cEmp1}
	totalExpense(&employees)
}