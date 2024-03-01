package main

import (
	"fmt"
)

func main() {
	var array = [...]int{1,2,3,4,5}
	slice := array[1:3]
	fmt.Println(slice)
	var emp Employee
	emp.name = "name"
	emp.age = 15
	emp.city = "ALA"
	emp.jobTitle = "Student"
	emp.PrintEmployeeInfo()
}


type Person struct {
	name string
	age  int
	city string
}

type Employee struct{
	Person
	jobTitle string
}

func (emp Employee) PrintEmployeeInfo(){
	fmt.Print("Name : " + emp.name + "\nCity : " + emp.city + "\nAge : ")
	fmt.Println(emp.age)
	fmt.Println("Title : " + emp.jobTitle)
}