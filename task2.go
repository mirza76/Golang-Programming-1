//  Abdur Rahman
//  19I-0584
//  BlockChan : B
//  Date : 9/22/2022

package main

import "fmt"

type employee struct {
	name     string
	salary   int
	position string
}

type company struct {
	companyName string
	employees   []employee
}

func main() {
	emp1 := employee{"Amir", 80000, "Full-Stack Developer"}
	emp2 := employee{"John", 90000, "Full-Stack Developer"}
	emp3 := employee{"Mark", 100000, "Full-Stack Developer"}
	emplys := []employee{emp1, emp2, emp3}
	comp := company{"Tetra", emplys}
	fmt.Println()
	fmt.Println(comp)
}
