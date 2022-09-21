//  Abdur Rahman
//  19I-0584
//  BlockChan : B
//  Date : 9/22/2022

package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

// adding an array for the subject a student
// is currently studying.
// Then , Calculating hash of the block data and display it

type Student struct {
	roll_number int
	name        string
	address     string
	subjects    []string
}

func new_student(roll_number int, name string, address string) *Student {
	s := new(Student)
	s.roll_number = roll_number
	s.name = name
	s.address = address
	return s
}

type StudentList struct {
	list []*Student
}

func (ls *StudentList) create_student(roll_number int, name string, address string) *Student {
	st := new_student(roll_number, name, address)
	ls.list = append(ls.list, st)
	return st
}

func print_student_list(ls *StudentList) {
	// Printing the student list.
	StudentList := ls.list
	for i := 0; i < len(StudentList); i++ {
		fmt.Printf("%s List %d %s \n", strings.Repeat("=", 10), i+1, strings.Repeat("=", 10))
		fmt.Println("Student Roll Number: ", StudentList[i].roll_number)
		fmt.Println("Student Name: ", StudentList[i].name)
		fmt.Println("Student Address: ", StudentList[i].address)
	}
}

func CalculateHash(s Student) string {
	h := sha256.New()
	h.Write([]byte(s.name))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func main() {

	ls := new(StudentList)
	ls.create_student(21, "Abdur Rahman", "Islamabad, Pakistan")
	ls.create_student(22, "Imran", "Washington, USA")
	ls.create_student(23, "Ali", "London, UK")
	ls.create_student(24, "Ijaz Ahmed", "Lahore, Pakistan")
	ls.create_student(25, "Naveed", "Karachi, Pakistan")
	ls.create_student(26, "Amir", "Lahore, Pakistan")

	print_student_list(ls)
	fmt.Println()

	// Calculatin hash of the block data and display it

	output_hash := CalculateHash(*ls.list[0])
	fmt.Println("Hash of the block data: ", output_hash)

}
