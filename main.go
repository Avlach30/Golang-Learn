package main

import (
	"fmt"
	arraystudents "golang-basic-fundamental/array-students"
)

func main() {
	studentMales := arraystudents.StudentMales() //* Call a function which is returned an array
	studentMalesLen := len(studentMales) //* Get a length of an array

	fmt.Println(studentMales)
	fmt.Println(studentMalesLen)

	studentFemales := arraystudents.StudentFemales()
	studentFemaleLength := len(studentFemales)

	fmt.Println(studentFemales)
	fmt.Println(studentFemaleLength)

	//* Looping an array with loop limit of array length
	for _, student := range studentMales{
		println("List student: ", student)
	}

	for index, student := range studentFemales{
		orderedList := index + 1
		println("Absence:", orderedList, ". Name: ", student)
	}
}