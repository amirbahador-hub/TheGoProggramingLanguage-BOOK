package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

func returnStruct() Employee {
	return Employee{
		ID:        0,
		Name:      "ali",
		Address:   "qom",
		DoB:       time.Time{},
		Position:  "junior",
		Salary:    100,
		ManagerID: 0,
	}
}

var ali Employee
var ali2 *Employee
var ali3 *Employee = &Employee{}

func returnPointer() *Employee {
	ali.ID = 1
	ali.Name = "ali"
	ali.Address = "qom"
	ali.DoB = time.Time{}
	ali.Position = "junior"
	ali.Salary = 100
	ali.ManagerID = 0
	return &ali
}

func returnPointer2() *Employee {
	//(*ali2).ID = 1
	//(*ali2).Name = "ali"
	//(*ali2).Address = "qom"
	//(*ali2).DoB = time.Time{}
	//(*ali2).Position = "junior"
	//(*ali2).Salary = 100
	//(*ali2).ManagerID = 0
	return ali2
}

func returnPointer3() *Employee {
	(*ali3).ID = 1
	(*ali3).Name = "ali"
	(*ali3).Address = "qom"
	(*ali3).DoB = time.Time{}
	(*ali3).Position = "junior"
	(*ali3).Salary = 100
	(*ali3).ManagerID = 0
	return ali3
}

func main() {
	var dilber Employee
	dilber.Salary -= 200

	fmt.Println(dilber) //	{0   0001-01-01 00:00:00 +0000 UTC  -200 0}

	position := dilber.Position
	dilber.Position = "senior" + position

	fmt.Println(dilber) //	{0   0001-01-01 00:00:00 +0000 UTC senior -200 0}

	position2 := &dilber.Position
	*position2 = "senior " + *position2

	fmt.Println(dilber) //	{0   0001-01-01 00:00:00 +0000 UTC senior senior -200 0}

	var employeeOfTheMonth *Employee = &dilber
	employeeOfTheMonth.Position += " team player "

	fmt.Println(dilber) //	{0   0001-01-01 00:00:00 +0000 UTC senior senior team player  -200 0}

	(*employeeOfTheMonth).Position += " t m "

	fmt.Println(dilber) //	{0   0001-01-01 00:00:00 +0000 UTC senior senior team player  t m  -200 0}

	fmt.Println(returnStruct())   //	{0 ali qom 0001-01-01 00:00:00 +0000 UTC junior 100 0}
	fmt.Println(returnPointer())  //	&{1 ali qom 0001-01-01 00:00:00 +0000 UTC junior 100 0}
	fmt.Println(returnPointer2()) //	panic: runtime error: invalid memory address or nil pointer dereference
	fmt.Println(returnPointer3()) //	&{1 ali qom 0001-01-01 00:00:00 +0000 UTC junior 100 0}

	seen := make(map[string]struct{})
	seen2 := make(map[string]Employee)
	seen3 := make(map[string]struct{ a int })

	fmt.Println(seen)  //	map[]
	fmt.Println(seen2) // map[]

	seen["ali"] = struct{}{}
	seen2["ali"] = Employee{
		ID:        0,
		Name:      "ali",
		Address:   "qom",
		DoB:       time.Time{},
		Position:  "junior",
		Salary:    100,
		ManagerID: 0,
	}
	seen3["ali"] = struct{ a int }{a: 2}

	fmt.Println(seen)  //	map[ali:{}]
	fmt.Println(seen2) //	map[ali:{0 ali qom 0001-01-01 00:00:00 +0000 UTC junior 100 0}]
	fmt.Println(seen3) //	map[ali:{2}]
}
