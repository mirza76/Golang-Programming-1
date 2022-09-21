//  Abdur Rahman
//  19I-0584
//  BlockChan : B
//  Date : 9/22/2022

package main

import "fmt"

type Person struct {
	f_name  string
	l_name  string
	age     int
	salary  int
	job     string
	address Address
}

type Address struct {
	city  string
	state string
}

func printPerson(p Person) {
	println("First Name:", p.f_name)
	println("Last Name:", p.l_name)
	println("Age:", p.age)
	println("Salary:", p.salary)
	println("Job:", p.job)
	println("City:", p.address.city)
	println("State:", p.address.state)
}

// in main pass a structure as a function argument and print its values.
func main() {
	p := Person{
		f_name: "John",
		l_name: "Doe",
		age:    30,
		salary: 10000,
		job:    "Software Engineer",
		address: Address{
			city:  "New York",
			state: "NY",
		},
	}
	printPerson(p)
	fmt.Println("\nPassed a structure by value to a function and printed its values")
	fmt.Println("\nTask 1 Done")
	fmt.Println("------------------------------------------------------------")
	fmt.Println("------------------------------------------------------------")

}
